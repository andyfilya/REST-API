package handler

import (
	"github.com/andyfilya/restapi/pkg/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Handler struct {
	services *service.Service
	logger   *logrus.Logger
}

func InitNewHandler(services *service.Service) *Handler {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     true,
		},
	}

	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (hr *Handler) StartRoute() http.Handler {
	hr.logger.Infof("starting route")
	mux := http.NewServeMux()

	// REGISTER (BEFORE AUTH) //

	mux.HandleFunc("/auth/register", hr.registerNewUser)
	mux.HandleFunc("/auth/signin", hr.signinUser)

	// ENDPOINTS WITH AUTH (AFTER AUTH ... WITH JWT TOKEN) ACTORS //

	mux.HandleFunc("/auth/check", hr.middlewareAuth(hr.checkMiddlewareHealth))                // check health (middleware)
	mux.HandleFunc("/api/create/actor", hr.middlewareAuth(hr.createActor))                    // create actor
	mux.HandleFunc("/api/delete/actor", hr.middlewareAuth(hr.deleteActor))                    // delete actor
	mux.HandleFunc("/api/update/actor", hr.middlewareAuth(hr.updateActor))                    // update actor
	mux.HandleFunc("/api/find/actorfragments", hr.middlewareAuth(hr.findFilmByActorFragment)) // find by fragments

	// ENDPOINTS WITH AUTH (AFTER AUTH ... WITH JWT TOKEN) FILMS //

	mux.HandleFunc("/api/add/actor/film", hr.middlewareAuth(hr.addActorFilmToFilm))          // insert actor into film
	mux.HandleFunc("/api/get/film", hr.middlewareAuth(hr.getAllFilmsWithActors))             // get all films
	mux.HandleFunc("/api/create/film/without", hr.middlewareAuth(hr.createFilmWithoutActor)) // create film without actor
	mux.HandleFunc("/api/create/film/one", hr.middlewareAuth(hr.createFilm))                 // create film with one actor_id
	mux.HandleFunc("/api/create/film/many", hr.middlewareAuth(hr.createFilmActors))          // create film with many actor_ids
	mux.HandleFunc("/api/delete/film", hr.middlewareAuth(hr.deleteFilm))                     // delete film
	mux.HandleFunc("/api/update/film", hr.middlewareAuth(hr.changeFilm))                     // change film

	return mux
}
