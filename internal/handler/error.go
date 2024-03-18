package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Error string `json:"error"`
}

func returnErr(w http.ResponseWriter, statusCode int, errStr string) {
	js, err := json.Marshal(errorResponse{errStr})
	if err != nil {
		logrus.Error("Cannot convert error to json")
	}
	http.Error(w, string(js), statusCode)
}
