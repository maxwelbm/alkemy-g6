package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	// default status code
	defaultStatusCode := http.StatusInternalServerError
	// check if status code is valid
	if statusCode > 299 && statusCode < 600 {
		defaultStatusCode = statusCode
	}

	// response
	body := ErrorResponse{
		Status:  http.StatusText(defaultStatusCode),
		Message: message,
	}
	bytes, err := json.Marshal(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write response
	// - set header: before code due to it sets by default "text/plain"
	w.Header().Set("Content-Type", "application/json")
	// - set status code
	w.WriteHeader(defaultStatusCode)
	// - write body
	_, err = w.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Errorf(w http.ResponseWriter, statusCode int, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	Error(w, statusCode, message)
}
