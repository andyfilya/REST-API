package handler

import (
	"encoding/json"
	"github.com/andyfilya/restapi"
	"github.com/sirupsen/logrus"
	"net/http"
)

type NewActor struct {
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	DateBirth string `json:"date_birth"`
}

func (hr *Handler) createActor(w http.ResponseWriter, r *http.Request) {
	newActor := NewActor{}

	err := json.NewDecoder(r.Body).Decode(&newActor)
	if err != nil {
		logrus.Errorf("error while decode request body : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "bad body of request.")
		return
	}

	toSendActor := restapi.Actor{
		FirstName: newActor.FirstName,
		LastName:  newActor.LastName,
		DateBirth: newActor.DateBirth,
	}

	actorId, err := hr.services.CreateActor(toSendActor)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}
	toSend := map[string]interface{}{
		"actorId": actorId,
	}

	sendBytes, err := json.Marshal(toSend)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Write(sendBytes)
}
