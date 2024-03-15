package main

import (
	"github.com/andyfilya/restapi"
	"github.com/andyfilya/restapi/config"
	"github.com/andyfilya/restapi/pkg/handler"
	"github.com/andyfilya/restapi/pkg/repo"
	"github.com/andyfilya/restapi/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.InitGlobalConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := repo.NewDataBase(&cfg.UserDatabaseCfg)
	if err != nil {
		logrus.Fatal(err)
	}

	repo := repo.InitNewRepository(db)
	services := service.InitNewService(repo)
	handlers := handler.InitNewHandler(services)

	serv := new(restapi.Server)
	if err := serv.InitServer(&cfg.ServCfg, handlers.StartRoute()); err != nil {
		logrus.Fatalf("error while run server : [%v]", err)
	}

	logrus.Info("server started...")
}
