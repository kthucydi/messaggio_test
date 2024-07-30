package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
		(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Methods, Authorization")
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
