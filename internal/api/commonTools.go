package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	gmux "github.com/gorilla/mux"
)

const (
	CORSAllow     bool = true
	CORSForbidden bool = false
)

// SetHeaders set next headers to responce:
// - accessCORS - bool;
// - status - int;
// - Content-type - string.
func SetHeaders(w *http.ResponseWriter, cors bool, status int, contentType string) {
	if cors {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, DELETE")
		(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}

	(*w).WriteHeader(status)

	if contentType != "" {
		(*w).Header().Set("Content-Type", contentType)
	}
}

func CreateJsonMessage(msg string, error bool) []byte {
	data, err := json.Marshal(JSONMessage{
		Message:   msg,
		ErrorBool: error,
	})
	if err != nil {
		data = []byte("error during create message")
	}
	return data
}

func getIDFromURL(r *http.Request) (id int, err error) {

	if idString := r.URL.Query().Get("id"); idString != "" {
		id, err = strconv.Atoi(idString)
		if err != nil {
			return 0, err
		}
	} else {
		return -1, nil
	}
	return id, nil
}

func getUIntPathParams(paramName string, r *http.Request) (paramValue int, err error) {
	pathParams := gmux.Vars(r)

	if value, ok := pathParams[paramName]; ok {
		paramValue, err = strconv.Atoi(value)
		if err != nil || paramValue < 0 {
			return -1, fmt.Errorf("error, %s incorrect", paramName)
		}
	}
	return
}

func ResponceCORSAllowed(w http.ResponseWriter, r *http.Request) {
	// SetHeaders(&w, CORSAllow, 200, "application/json")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(CreateJsonMessage("CORS allowed", false)))
}

func CreateJSONMessageResponse(w *http.ResponseWriter, status int, message string, errbool bool) {
	SetHeaders(w, CORSAllow, status, "application/json") // failed
	fmt.Fprint(*w, string(CreateJsonMessage(message, true)))
}

func UnmarshalJSONBody(w *http.ResponseWriter, r *http.Request, structJSON interface{}) (err error) {
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(structJSON)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			Log.Warning("Bad Request. Wrong Type provided for field " + unmarshalErr.Field)
			// CreateJSONMessageResponse(w, 400, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, true)
		} else {
			Log.Warning("Bad Request " + err.Error())
			// CreateJSONMessageResponse(w, 400, "Bad Request "+err.Error(), true)
		}
	}
	return
}
