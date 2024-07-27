package api

import (
	"encoding/json"
	"fmt"
	"io"
	kh "messaggio_test/internal/kafkahandle"
	pgstorage "messaggio_test/internal/pgstorage"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if err := pgstorage.UpdateProcessed([]int{4, 3, 2}); err != nil { // ARRAY!!!
		Log.Errorf("can not update DB:%v", err)
		return
	}
	SetHeaders(&w, CORSAllow, http.StatusOK, "application/json")
	fmt.Fprint(w, string(CreateJsonMessage("Server is running", false)))
}

func messageHandler(w http.ResponseWriter, r *http.Request) {

	// Read request body for message
	messageBuff, err := io.ReadAll(r.Body)
	if err != nil {
		Log.Errorf("can not read request body")
		return
	}
	defer r.Body.Close()

	// convert []byte to string
	message := string(messageBuff)
	Log.Infof("GET MESSAGE '%s'", message)

	go func() {
		id, err := pgstorage.MessageCreate(message)
		if err != nil {
			Log.Errorf("can not write message to BD: %v", err)
			return
		}

		// Send message to kafka topic
		kh.Kafka.SendMessage(message, id)
	}()

	SetHeaders(&w, CORSAllow, http.StatusOK, "application/json")
	fmt.Fprint(w, string(CreateJsonMessage("message was getting", false)))
}

func statHandler(w http.ResponseWriter, r *http.Request) {

	stat, err := pgstorage.GetStatistic()
	if err != nil {
		Log.Errorf("can not get statistic from DB: %v", err)
		return
	}

	if data, err := json.Marshal(stat); err != nil {
		Log.Errorf("can not Marshal statistic to JSON: %v", err)
		SetHeaders(&w, CORSAllow, http.StatusInternalServerError, "")
	} else {
		SetHeaders(&w, CORSAllow, 200, "application/json")
		fmt.Fprint(w, string(data))
		return
	}
	// SetHeaders(&w, CORSAllow, http.StatusOK, "application/json")
	// fmt.Fprint(w, string(CreateJsonMessage("Statistic info:", false)))
}
