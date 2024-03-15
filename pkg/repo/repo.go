package repo

import (
	"github.com/andyfilya/restapi"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	NewUser(user restapi.User) (int, error)
	FindUser(username, password string) (restapi.User, error)
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

type Repository struct {
	Authorization
	Actor
	Film
}

func InitNewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: InitAuthDataBase(db),
	}
}
