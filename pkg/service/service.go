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
	DeleteActor(actorId int) error
	ChangeActor(actorId int, toChange string) error
	FindActorFilm(actor string) ([]restapi.Film, error)
}

type Film interface {
	CreateFilm(film restapi.Film) (int, error)
	DeleteFilm(filmId int) error
	ChanageFilm(filmId int, toChange string) error
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
	}
}
