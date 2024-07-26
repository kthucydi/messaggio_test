package api

import (
	"fmt"
	"net/http"
)

func serverStatus(w http.ResponseWriter, r *http.Request) {
	SetHeaders(&w, CORSAllow, http.StatusOK, "application/json")
	fmt.Fprint(w, string(CreateJsonMessage("Server is running", false)))
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders(&w, CORSAllow, http.StatusOK, "application/json")
	fmt.Fprint(w, string(CreateJsonMessage("message was getting", false)))
}

func statHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders(&w, CORSAllow, http.StatusOK, "application/json")
	fmt.Fprint(w, string(CreateJsonMessage("Statistic info:", false)))
}
