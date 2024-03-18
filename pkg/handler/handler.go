package handler

import (
	_ "github.com/andyfilya/restapi/docs"
	"github.com/andyfilya/restapi/pkg/service"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
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

	// SWAGGER //
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// REGISTER (BEFORE AUTH) //

	mux.HandleFunc("/auth/register", hr.registerNewUser) // register new user
	mux.HandleFunc("/auth/signin", hr.signinUser)        // sign in user

	// ENDPOINTS WITH AUTH (AFTER AUTH ... WITH JWT TOKEN) ACTORS //

	mux.HandleFunc("/auth/check", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.checkMiddlewareHealth))))               // check health (middleware)
	mux.HandleFunc("/api/create/actor", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.createActor))))                   // create actor
	mux.HandleFunc("/api/delete/actor", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.deleteActor))))                   // delete actor
	mux.HandleFunc("/api/update/actor", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.updateActor))))                   // update actor
	mux.HandleFunc("/api/find/actorfragments", hr.setIdMiddleware(hr.loggerMiddleware(hr.middlewareAuth(hr.findFilmByActorFragment)))) // find by fragments

	// ENDPOINTS WITH AUTH (AFTER AUTH ... WITH JWT TOKEN) FILMS //

	mux.HandleFunc("/api/add/actor/film", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.addActorFilmToFilm))))          // insert actor into film
	mux.HandleFunc("/api/get/film", hr.setIdMiddleware(hr.loggerMiddleware(hr.middlewareAuth(hr.getAllFilmsWithActors))))              // get all films
	mux.HandleFunc("/api/create/film/without", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.createFilmWithoutActor)))) // create film without actor
	mux.HandleFunc("/api/create/film/one", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.createFilm))))                 // create film with one actor_id
	mux.HandleFunc("/api/create/film/many", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.createFilmActors))))          // create film with many actor_ids
	mux.HandleFunc("/api/delete/film", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.deleteFilm))))                     // delete film
	mux.HandleFunc("/api/update/film", hr.setIdMiddleware(hr.loggerMiddleware(hr.adminMiddleware(hr.changeFilm))))                     // change film

	return mux
}
