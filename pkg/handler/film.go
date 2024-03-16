package handler

import (
	"encoding/json"
	"github.com/andyfilya/restapi"
	"net/http"
)

func (hr *Handler) createFilm(w http.ResponseWriter, r *http.Request) {
	newFilm := restapi.Film{}

	err := json.NewDecoder(r.Body).Decode(&newFilm)
	if err != nil {
		hr.logger.Errorf("error while decode body request : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "need to create right request.")
		return
	}

	filmId, err := hr.services.CreateFilm(newFilm)
	if err != nil {
		hr.logger.Errorf("error while creating new film")
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	var toSend = map[string]interface{}{
		"filmId": filmId,
	}

	toSendBytes, err := json.Marshal(toSend)
	if err != nil {
		hr.logger.Errorf("error while marshal to send map")
		newErrWrite(w, http.StatusInternalServerError, "unknown error.")
		return
	}

	hr.logger.Infof("successs create film with id [id : %d]", filmId)
	w.Write(toSendBytes)
}

type toDelete struct {
	Title string `json:"title"'`
}

func (hr *Handler) deleteFilm(w http.ResponseWriter, r *http.Request) {
	toDeleteFilm := restapi.Film{}
	newDeleteFilm := toDelete{}
	err := json.NewDecoder(r.Body).Decode(&newDeleteFilm)

	toDeleteFilm.Title = newDeleteFilm.Title

	if err != nil {
		hr.logger.Errorf("error while decode to delete film : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "bad request.")
		return
	}

	err = hr.services.DeleteFilm(toDeleteFilm)
	if err != nil {
		hr.logger.Errorf("erorr delete to delete film : [%v]", err)
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	var toSend = map[string]interface{}{
		"deleted": toDeleteFilm.Title,
		"success": true,
	}

	toSendBytes, err := json.Marshal(toSend)
	if err != nil {
		hr.logger.Errorf("error while decode the message : [%v]", err)
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write(toSendBytes)
}

func (hr *Handler) changeFilm(w http.ResponseWriter, r *http.Request) {
	updateFilm := restapi.ToChangeFilm{}
	oldFilm := restapi.Film{}
	newFilm := restapi.Film{}

	err := json.NewDecoder(r.Body).Decode(&updateFilm)
	if err != nil {
		hr.logger.Errorf("error while decode request body : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "bad body of request.")
		return
	}

	oldFilm.Title = updateFilm.Title
	oldFilm.Date = updateFilm.Date
	oldFilm.Rate = updateFilm.Rate
	oldFilm.Description = updateFilm.Description

	newFilm.Title = updateFilm.ToChangeTitle
	newFilm.Date = updateFilm.ToChangeDate
	newFilm.Rate = updateFilm.ToChangeRate
	newFilm.Description = updateFilm.ToChangeDescription

	err = hr.services.ChangeFilm(newFilm, oldFilm)
	if err != nil {
		hr.logger.Errorf("error while change film in database : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "something went wrong...")
		return
	}

	toSend := map[string]interface{}{
		"old":     oldFilm,
		"new":     newFilm,
		"success": true,
	}

	sendBytes, err := json.Marshal(toSend)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Write(sendBytes)
}
