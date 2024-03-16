package repo

import (
	"errors"
	"fmt"
	"github.com/andyfilya/restapi"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	filmTbl = "films"
)

type FilmDataBase struct {
	db *sqlx.DB
}

func InitFilmDataBase(db *sqlx.DB) *FilmDataBase {
	return &FilmDataBase{
		db: db,
	}
}

func (fdb *FilmDataBase) CreateFilm(film restapi.Film) (int, error) {
	var filmId int
	query := fmt.Sprintf("INSERT INTO %s (film_title, film_date, film_description, film_rate) VALUES ($1, $2, $3, $4) RETURNING film_id", filmTbl)
	row := fdb.db.QueryRow(query, film.Title, film.Date, film.Description, film.Rate)
	err := row.Scan(&filmId)
	if err != nil {
		logrus.Errorf("error while scan film id in var : [%v]", err)
		return -1, err
	}

	return filmId, nil
}

func (fdb *FilmDataBase) DeleteFilm(film restapi.Film) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE film_title=$1", filmTbl)
	_, err := fdb.db.Exec(query, film.Title)
	if err != nil {
		logrus.Errorf("error while exec film table : [%v]", err)
		return errors.New("error while delete film from database")
	}
	return nil
}

func (fdb *FilmDataBase) ChangeFilm(newFilm restapi.Film, oldFilm restapi.Film) error {
	query := fmt.Sprintf("UPDATE %s SET film_title=$1, film_description=$2, film_date=$3, film_rate=$4 WHERE film_title=$5", filmTbl)
	_, err := fdb.db.Exec(query, newFilm.Title, newFilm.Description, newFilm.Date, newFilm.Rate, oldFilm.Title)
	if err != nil {
		logrus.Errorf("error while exec update film table : [%v]", err)
		return err
	}
	return nil
}

func (fdb *FilmDataBase) ActorsFilm(filmId int) ([]restapi.Actor, error) {
	return nil, nil
}
