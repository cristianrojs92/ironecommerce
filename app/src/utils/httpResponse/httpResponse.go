package httpResponse

import (
	"encoding/json"
	"net/http"
	"../../models"
)

func Error(w http.ResponseWriter, message string, code int) {
	errObj := models.HttpError{
		Message:    message,
		HttpStatus: code,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errObj); err == nil {
		w.Write(j)
	}
}
