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

func (fs *FilmService) CreateFilm(film restapi.Film) (int, error) {
	if !titleFilmCheck(film.Title) {
		return -1, errors.New("title of film is very big (max 150 chars) or empty")
	}
	if !descriptionFilmCheck(film.Description) {
		return -1, errors.New("description of film is very big (max 1000 chars)")
	}

	return fs.repo.CreateFilm(film)
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

func (fs *FilmService) ActorsFilm(filmId int) ([]restapi.Actor, error) {
	return nil, nil
}
