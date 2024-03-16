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

type Repository struct {
	Authorization
	Actor
	Film
}

func InitNewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: InitAuthDataBase(db),
		Actor:         InitActorDataBase(db),
		Film:          InitFilmDataBase(db),
	}
}
