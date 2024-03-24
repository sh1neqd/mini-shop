package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newResponse(w http.ResponseWriter, statusCode int, message string) {
	var response = errorResponse{Message: message}
	w.Header().Set("Content-Type", "application/json")
	logrus.Error(message)
	w.WriteHeader(statusCode)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
