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

type toSendCreate struct {
	ActorId int
}

// @Summary CreateActor
// @Security ApiKeyAuth
// @Tags create
// @Description Create actor
// @ID create_actor
// @Accept  json
// @Produce  json
// @Param input body restapi.Actor true "actor info"
// @Success 200 {object} toSendCreate
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/create/actor [post]
func (hr *Handler) createActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
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
	toSend := toSendCreate{
		ActorId: actorId,
	}

	sendBytes, err := json.Marshal(toSend)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Write(sendBytes)
}

type toSendDeleted struct {
	Deleted string
}

// @Summary DeleteActor
// @Security ApiKeyAuth
// @Tags delete
// @Description Delete actor
// @ID delete_actor
// @Accept  json
// @Produce  json
// @Param input body NewActor true "actor info"
// @Success 200 {object} toSendDeleted
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/create/actor [delete]
func (hr *Handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
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
	toSend := toSendDeleted{
		Deleted: toSendActor.FirstName + " " + toSendActor.LastName,
	}

	sendBytes, err := json.Marshal(toSend)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Write(sendBytes)
}

type toSendUpdate struct {
	Old string
	New string
}

// @Summary UpdateActor
// @Security ApiKeyAuth
// @Tags update
// @Description Update actor
// @ID update_actor
// @Accept  json
// @Produce  json
// @Param input body restapi.ToChange true "to_change_actor info"
// @Success 200 {object} toSendUpdate
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/create/actor [put]
func (hr *Handler) updateActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
	updateActor := restapi.ToChange{}
	oldActor := restapi.Actor{}
	newActor := restapi.Actor{}

	err := json.NewDecoder(r.Body).Decode(&updateActor)
	if err != nil {
		hr.logger.Errorf("error while decode request body : [%v]", err)
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
		hr.logger.Errorf("error while change actor in database : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "something went wrong...")
		return
	}

	toSend := toSendUpdate{
		Old: oldActor.FirstName + " " + oldActor.LastName,
		New: newActor.FirstName + " " + newActor.LastName,
	}

	sendBytes, err := json.Marshal(toSend)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Write(sendBytes)
}

// @Summary FindFilmByActorFragment
// @Security ApiKeyAuth
// @Tags find
// @Description Find films by actor fragments
// @ID find_films
// @Accept  json
// @Produce  json
// @Param input body restapi.ActorFragment true "actor fragment info"
// @Success 200 {object} ToSend
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/find/actorfragments [post]
func (hr *Handler) findFilmByActorFragment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}

	actorFragments := restapi.ActorFragment{}
	err := json.NewDecoder(r.Body).Decode(&actorFragments)

	if err != nil {
		logrus.Errorf("error while decode request body")
		newErrWrite(w, http.StatusBadRequest, "error while decode body | bad request")
		return
	}

	films, err := hr.services.FindActorFilm(actorFragments)
	if err != nil {
		logrus.Errorf("error while find actor by fragments")
		newErrWrite(w, http.StatusInternalServerError, "unknown error")
		return
	}
	toSendMap := map[string]interface{}{
		"actor_films": films,
	}
	toSend := ToSend{toSendMap}
	toSendBytes, err := json.Marshal(toSend)
	if err != nil {
		logrus.Errorf("error while marshal to send map")
		newErrWrite(w, http.StatusInternalServerError, "unknown error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(toSendBytes)
}
