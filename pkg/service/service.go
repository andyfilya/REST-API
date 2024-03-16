package service

import (
	"github.com/andyfilya/restapi"
	"github.com/andyfilya/restapi/pkg/repo"
)

type Authorization interface {
	NewUser(user restapi.User) (int, error)
	NewUserToken(username, password string) (string, error)
	ParseUserToken(authToken string) (int, error)
}

type Actor interface {
	CreateActor(actor restapi.Actor) (int, error)
	DeleteActor(actor restapi.Actor) error
	ChangeActor(oldActor restapi.Actor, newActor restapi.Actor) error
	FindActorFilm(actor string) ([]restapi.Film, error)
}

type Film interface {
	CreateFilm(film restapi.Film) (int, error)
	DeleteFilm(film restapi.Film) error
	ChangeFilm(newFilm restapi.Film, oldFilm restapi.Film) error
	ActorsFilm(filmid int) ([]restapi.Actor, error)
}

type Service struct {
	Authorization
	Actor
	Film
}

func InitNewService(repo *repo.Repository) *Service {
	return &Service{
		Authorization: InitAuthService(repo.Authorization),
		Actor:         InitActorService(repo.Actor),
		Film:          InitFilmService(repo.Film),
	}
}
