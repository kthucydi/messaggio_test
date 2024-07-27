package api

import (
	"fmt"
	"io"
	kh "messaggio_test/internal/kafkahandle"
	pgstorage "messaggio_test/internal/pgstorage"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
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
	id, err := pgstorage.MessageCreate(message)
	if err != nil {
		Log.Errorf("can not write message to BD: %v", err)
		return
	}

	// Send message to kafka topic
	kh.Kafka.SendMessage(message, id)
	SetHeaders(&w, CORSAllow, http.StatusOK, "application/json")
	fmt.Fprint(w, string(CreateJsonMessage("message was getting", false)))
}

func statHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders(&w, CORSAllow, http.StatusOK, "application/json")
	fmt.Fprint(w, string(CreateJsonMessage("Statistic info:", false)))
}
