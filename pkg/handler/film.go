package handler

import (
	"encoding/json"
	"github.com/andyfilya/restapi"
	"github.com/sirupsen/logrus"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type Create struct {
	ActorId int `json:"actorId"`
	restapi.Film
}

type ToCreateWithActors struct {
	ActorIds []int `json:"actorIds"`
	restapi.Film
}

// @Summary CreateFilmWithActors
// @Security ApiKeyAuth
// @Tags create
// @Description Create film with actors
// @ID create_with
// @Accept  json
// @Produce  json
// @Param input body ToCreateWithActors true "film information"
// @Success 200 {object} ToSend
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/create/film/many [post]
func (hr *Handler) createFilmActors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
	newCreate := ToCreateWithActors{}

	err := json.NewDecoder(r.Body).Decode(&newCreate)
	if err != nil {
		hr.logger.Errorf("error while decode body request : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "need to create right request.")
		return
	}
	hr.logger.Debugf("Creating film : %v", newCreate)
	newFilm := newCreate.Film
	actorIds := newCreate.ActorIds
	filmId, err := hr.services.CreateFilmActors(actorIds, newFilm)
	if err != nil {
		hr.logger.Errorf("error while creating new film")
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	var toSend = ToSend{map[string]interface{}{
		"filmId": filmId,
	}}

	toSendBytes, err := json.Marshal(toSend)
	if err != nil {
		hr.logger.Errorf("error while marshal to send map")
		newErrWrite(w, http.StatusInternalServerError, "unknown error.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(toSendBytes)
}

// @Summary CreateFilm
// @Security ApiKeyAuth
// @Tags create
// @Description Create film without actors
// @ID create_only_one
// @Accept  json
// @Produce  json
// @Param input body restapi.Film true "film information"
// @Success 200 {object} ToSend
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/create/film/without [post]
func (hr *Handler) createFilmWithoutActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
	newFilm := restapi.Film{}
	err := json.NewDecoder(r.Body).Decode(&newFilm)
	if err != nil {
		hr.logger.Errorf("error while decode body request : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "need to create right request.")
		return
	}

	filmId, err := hr.services.CreateFilmWithoutActor(newFilm)
	if err != nil {
		hr.logger.Errorf("error create film without actor : [%v]", err)
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	toSend := map[string]interface{}{
		"filmId": filmId,
	}

	toSendBytes, err := json.Marshal(toSend)
	if err != nil {
		hr.logger.Errorf("error create film without actor : [%v]", err)
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(toSendBytes)
}

// @Summary CreateFilmWithOne
// @Security ApiKeyAuth
// @Tags create
// @Description Create film with one actor
// @ID create_one
// @Accept  json
// @Produce  json
// @Param input body Create true "film information"
// @Success 200 {object} ToSend
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/create/film/one  [post]
func (hr *Handler) createFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}

	newCreate := Create{}

	err := json.NewDecoder(r.Body).Decode(&newCreate)
	if err != nil {
		hr.logger.Errorf("error while decode body request : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "need to create right request.")
		return
	}
	hr.logger.Debugf("Creating film : %v", newCreate)
	newFilm := newCreate.Film
	actorId := newCreate.ActorId
	filmId, err := hr.services.CreateFilm(actorId, newFilm)
	if err != nil {
		hr.logger.Errorf("error while creating new film")
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	var toSend = ToSend{map[string]interface{}{
		"filmId": filmId,
	}}

	toSendBytes, err := json.Marshal(toSend)
	if err != nil {
		hr.logger.Errorf("error while marshal to send map")
		newErrWrite(w, http.StatusInternalServerError, "unknown error.")
		return
	}

	hr.logger.Infof("successs create film with id [id : %d]", filmId)
	w.Header().Set("Content-Type", "application/json")
	w.Write(toSendBytes)
}

type toDelete struct {
	Title string `json:"title"'`
}

// @Summary DeleteFilm
// @Security ApiKeyAuth
// @Tags delete
// @Description Delete film
// @ID delete_film
// @Accept  json
// @Produce  json
// @Param input body toDelete true "delete film information"
// @Success 200 {object} ToSend
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/delete/film [delete]
func (hr *Handler) deleteFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
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

	var toSend = ToSend{map[string]interface{}{
		"deleted": toDeleteFilm.Title,
		"success": true,
	}}

	toSendBytes, err := json.Marshal(toSend)

	if err != nil {
		hr.logger.Errorf("error while decode the message : [%v]", err)
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(toSendBytes)
}

// @Summary ChangeFilm
// @Security ApiKeyAuth
// @Tags update
// @Description Update film
// @ID update
// @Accept  json
// @Produce  json
// @Param input body ToCreateWithActors true "film information"
// @Success 200 {object} ToSend
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/create/film/many [post]
func (hr *Handler) changeFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
	updateFilm := restapi.ToChangeFilm{}
	oldFilm := restapi.Film{}
	newFilm := restapi.Film{}

	err := json.NewDecoder(r.Body).Decode(&updateFilm)
	if err != nil {
		hr.logger.Errorf("error while decode request body : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "bad body of request")
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

	toSend := ToSend{map[string]interface{}{
		"old":     oldFilm,
		"new":     newFilm,
		"success": true,
	}}

	sendBytes, err := json.Marshal(toSend)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(sendBytes)
}

type ToSort struct {
	SortBy string `json:"sort_by"`
}

// @Summary GetAllFilms
// @Security ApiKeyAuth
// @Tags get_all
// @Description Get all films
// @ID get
// @Accept  json
// @Produce  json
// @Param input body ToSort true "to sort film information"
// @Success 200 {object} ToSend
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/get/film [post]
func (hr *Handler) getAllFilmsWithActors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
	var toSort ToSort
	err := json.NewDecoder(r.Body).Decode(&toSort)
	if err != nil {
		hr.logger.Errorf("error while decode request body")
		newErrWrite(w, http.StatusBadRequest, "bad request, repeat please")
		return
	}

	resp, err := hr.services.GetAllFilms()
	if err != nil {
		hr.logger.Errorf("error while get all films : [%v]", err)
		newErrWrite(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	switch toSort.SortBy {
	case "title":
		sort.Slice(resp, func(i, j int) bool {
			return resp[i].Title > resp[j].Title
		})
	case "rate":
		sort.Slice(resp, func(i, j int) bool {
			first, _ := strconv.ParseFloat(resp[i].Rate, 64)
			second, _ := strconv.ParseFloat(resp[j].Rate, 64)
			return first > second
		})
	case "date":
		sort.Slice(resp, func(i, j int) bool {
			t1, _ := time.Parse(resp[i].Date, "1978-01-01T00:00:00Z")
			t2, _ := time.Parse(resp[j].Date, "1978-01-01T00:00:00Z")

			return t1.Unix() > t2.Unix()
		})
	default:
		sort.Slice(resp, func(i, j int) bool {
			t1, _ := time.Parse(resp[i].Date, "1978-01-01T00:00:00Z")
			t2, _ := time.Parse(resp[j].Date, "1978-01-01T00:00:00Z")

			return t1.Unix() < t2.Unix()
		})
	}

	mp := map[string]interface{}{}
	for i, v := range resp {
		name := "film#" + strconv.Itoa(i)
		mp[name] = v
	}

	toReturn := ToSend{mp}
	hr.logger.Infof("sort_by : %s, before : %v", toSort.SortBy, resp)
	toSendBytes, err := json.Marshal(toReturn)
	if err != nil {
		hr.logger.Errorf("erorr while marshal to RETURN map : [%v]", err)
		newErrWrite(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(toSendBytes)
}

type AddActorToFilm struct {
	ActorId int `json:"actor_id"`
	FilmId  int `json:"film_id"`
}

// @Summary AddActorToFilm
// @Security ApiKeyAuth
// @Tags add_actor
// @Description add actor to film
// @ID update
// @Accept  json
// @Produce  json
// @Param input body AddActorToFilm true "film information"
// @Success 200 {object} ToSend
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/create/film/many [post]
func (hr *Handler) addActorFilmToFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
	add := AddActorToFilm{}

	err := json.NewDecoder(r.Body).Decode(&add)
	if err != nil {
		logrus.Errorf("error while decode request body : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "bad body, repeat please")
		return
	}

	hr.logger.Infof("add actor struct : [%v]", add)
	err = hr.services.AddActorToFilm(add.ActorId, add.FilmId)
	if err != nil {
		logrus.Errorf("error  : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "bad body, repeat please")
		return
	}

	var toSend = ToSend{map[string]interface{}{
		"success":  "true",
		"insetred": "add",
		"into":     add.FilmId,
	}}

	toSendBytes, err := json.Marshal(toSend)
	if err != nil {
		logrus.Errorf("error while marshal : [%v]", err)
		newErrWrite(w, http.StatusBadRequest, "bad body, repeat please")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(toSendBytes)
}
