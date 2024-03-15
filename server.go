package restapi

import (
	"context"
	"github.com/andyfilya/restapi/config"
	"net/http"
	"time"
)

type Server struct {
	serv *http.Server
}

func (srv *Server) InitServer(cfg *config.ServerConfig, defaultHandler http.Handler) error {
	srv.serv = &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      defaultHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.serv.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.serv.Shutdown(ctx)
}
