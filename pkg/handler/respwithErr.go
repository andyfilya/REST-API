package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type errorMessage struct {
	Message string `json:"message"`
}

func newErrWrite(w http.ResponseWriter, httpErr int, message string) {
	newErr := &errorMessage{
		Message: message,
	}
	bytesToSend, err := json.Marshal(newErr)
	if err != nil {
		logrus.Errorf("error while marshaling error in struct : [%v]", err)
		return
	}

	w.WriteHeader(httpErr)
	w.Write(bytesToSend)
}
