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

func (hr *Handler) deleteActor(w http.ResponseWriter, r *http.Request) {
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
	err = hr.services.DeleteActor(toSendActor)
	if err != nil {
		logrus.Errorf("error delete user : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "no such actors.")
		return
	}
	toSend := map[string]interface{}{
		"delete": toSendActor.FirstName + " " + toSendActor.LastName,
	}

	sendBytes, err := json.Marshal(toSend)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Write(sendBytes)
}

func (hr *Handler) updateActor(w http.ResponseWriter, r *http.Request) {
	updateActor := restapi.ToChange{}
	oldActor := restapi.Actor{}
	newActor := restapi.Actor{}

	err := json.NewDecoder(r.Body).Decode(&updateActor)
	if err != nil {
		logrus.Errorf("error while decode request body : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "bad body of request.")
		return
	}

	oldActor.FirstName = updateActor.FirstName
	oldActor.LastName = updateActor.LastName
	oldActor.DateBirth = updateActor.DateBirth
	newActor.FirstName = updateActor.ToChangeUsername
	newActor.LastName = updateActor.ToChangeSurname
	newActor.DateBirth = updateActor.ToChangeBirth

	err = hr.services.ChangeActor(oldActor, newActor)
	if err != nil {
		logrus.Errorf("error while change actor in database : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "something went wrong...")
		return
	}

	toSend := map[string]interface{}{
		"old": oldActor.FirstName + " " + oldActor.LastName,
		"new": newActor.FirstName + " " + newActor.LastName,
	}

	sendBytes, err := json.Marshal(toSend)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Write(sendBytes)
}
