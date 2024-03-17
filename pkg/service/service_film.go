package service

import (
	"errors"
	"github.com/andyfilya/restapi"
	"github.com/andyfilya/restapi/pkg/repo"
)

type FilmService struct {
	repo repo.Film
}

func InitFilmService(repo repo.Film) Film {
	return &FilmService{
		repo: repo,
	}
}

func (fs *FilmService) CreateFilmWithoutActor(film restapi.Film) (int, error) {
	return fs.repo.CreateFilmWithoutActor(film)
}

func (fs *FilmService) AddActorToFilm(actorId int, filmId int) error {
	return fs.repo.AddActorToFilm(actorId, filmId)
}

func (fs *FilmService) GetAllFilms() ([]restapi.Film, error) {
	return fs.repo.GetAllFilms()
}

func (fs *FilmService) CreateFilm(actorId int, film restapi.Film) (int, error) {
	if !titleFilmCheck(film.Title) {
		return -1, errors.New("title of film is very big (max 150 chars) or empty")
	}
	if !descriptionFilmCheck(film.Description) {
		return -1, errors.New("description of film is very big (max 1000 chars)")
	}

	return fs.repo.CreateFilm(actorId, film)
}

func (fs *FilmService) DeleteFilm(film restapi.Film) error {
	if !titleFilmCheck(film.Title) {
		return errors.New("title of film is very big (max 150 chars) or empty")
	}
	return fs.repo.DeleteFilm(film)
}

func (fs *FilmService) ChangeFilm(newFilm restapi.Film, oldFilm restapi.Film) error {
	oldFilm, newFilm = checkEmptyFilm(newFilm, oldFilm)
	return fs.repo.ChangeFilm(newFilm, oldFilm)
}

func (fs *FilmService) CreateFilmActors(actorIds []int, film restapi.Film) (int, error) {
	return fs.repo.CreateFilmActors(actorIds, film)
}
func (fs *FilmService) ActorsFilm(filmId int) ([]restapi.Actor, error) {
	return nil, nil
}
