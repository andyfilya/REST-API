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
	CreateFilmWithoutActor(film restapi.Film) (int, error)
	AddActorToFilm(actorId int, filmId int) error
	GetAllFilms() ([]restapi.Film, error)
	CreateFilm(actorId int, film restapi.Film) (int, error)
	CreateFilmActors(actorIds []int, film restapi.Film) (int, error)
	DeleteFilm(film restapi.Film) error
	ChangeFilm(newFilm restapi.Film, oldFilm restapi.Film) error
	ActorFilms(actorId int) ([]restapi.Film, error)
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
