package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andyfilya/restapi/config"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Host   string
	Port   string
	Router *http.ServeMux
	Logger *logrus.Logger
}

func InitServer(cfg *config.ServerConfig) *Server {
	server := &Server{
		Host:   cfg.Host,
		Port:   cfg.Port,
		Router: http.NewServeMux(),
	}

	return server
}

func (s *Server) Init() error {
	s.serverRoute()
	fmt.Printf("server listening on port %s", s.Port)
	fmt.Printf("server addr: %s", s.Host+":"+s.Port)
	if err := http.ListenAndServe(s.Host+":"+s.Port, s.Router); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Server) serverRoute() {
	s.Router.HandleFunc("/", s.helloHandler)
}

func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello !"))
}
