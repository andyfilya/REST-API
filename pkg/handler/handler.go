package handler

import (
	"github.com/andyfilya/restapi/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func InitNewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (hr *Handler) StartRoute() http.Handler {
	mux := http.NewServeMux()
	// REGISTER (BEFORE AUTH) //

	mux.HandleFunc("/auth/register", hr.registerNewUser)
	mux.HandleFunc("/auth/signin", hr.signinUser)
	mux.HandleFunc("/auth/check", hr.middlewareAuth(hr.checkMiddleware))

	return mux
}
