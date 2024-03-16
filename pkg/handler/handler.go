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

	// ENDPOINTS WITH AUTH (AFTER AUTH ... WITH JWT TOKEN) //
	mux.HandleFunc("/auth/check", hr.middlewareAuth(hr.checkMiddlewareHealth)) // check health (middleware)
	mux.HandleFunc("/api/create/actor", hr.middlewareAuth(hr.createActor))     // create actor
	mux.HandleFunc("/api/delete/actor", hr.middlewareAuth(hr.deleteActor))     // delete actor
	mux.HandleFunc("/api/update/actor", hr.middlewareAuth(hr.updateActor))     // update actor

	return mux
}
