package repo

import (
	"errors"
	"fmt"
	"github.com/andyfilya/restapi"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	filmTbl      = "films"
	actorFilmTbl = "actors_films"
)

type FilmDataBase struct {
	db *sqlx.DB
}

func InitFilmDataBase(db *sqlx.DB) *FilmDataBase {
	return &FilmDataBase{
		db: db,
	}
}

func (fdb *FilmDataBase) CreateFilmWithoutActor(film restapi.Film) (int, error) {
	tx, err := fdb.db.Begin()
	if err != nil {
		return -1, err
	}
	var filmId int
	createFilmQuery := fmt.Sprintf("INSERT INTO %s (film_title, film_date, film_description, film_rate) VALUES ($1, $2, $3, $4) RETURNING film_id", filmTbl)
	row := tx.QueryRow(createFilmQuery, film.Title, film.Date, film.Description, film.Rate)
	err = row.Scan(&filmId)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("error while scan film id in var : [%v]", err)
		return -1, err
	}

	return filmId, tx.Commit()
}
func (fdb *FilmDataBase) CreateFilm(actorId int, film restapi.Film) (int, error) {
	tx, err := fdb.db.Begin()
	if err != nil {
		return -1, err
	}
	var filmId int
	createFilmQuery := fmt.Sprintf("INSERT INTO %s (film_title, film_date, film_description, film_rate) VALUES ($1, $2, $3, $4) RETURNING film_id", filmTbl)
	row := tx.QueryRow(createFilmQuery, film.Title, film.Date, film.Description, film.Rate)
	err = row.Scan(&filmId)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("error while scan film id in var : [%v]", err)
		return -1, err
	}
	logrus.Debugf("actor_id: %d, film_id: %d", actorId, filmId)
	createActorFilm := fmt.Sprintf("INSERT INTO %s (a_id, f_id) VALUES ($1, $2) RETURNING id", actorFilmTbl)
	_, err = tx.Exec(createActorFilm, actorId, filmId)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("error while exec in actors films database : [%v]", err)
		return -1, err
	}
	return filmId, tx.Commit()
}

func (fdb *FilmDataBase) CreateFilmActors(actorIds []int, film restapi.Film) (int, error) {
	tx, err := fdb.db.Begin()
	if err != nil {
		return -1, err
	}
	var filmId int
	createFilmQuery := fmt.Sprintf("INSERT INTO %s (film_title, film_date, film_description, film_rate) VALUES ($1, $2, $3, $4) RETURNING film_id", filmTbl)
	row := tx.QueryRow(createFilmQuery, film.Title, film.Date, film.Description, film.Rate)
	err = row.Scan(&filmId)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("error while scan film id in var : [%v]", err)
		return -1, err
	}

	for _, actorId := range actorIds {
		createActorFilm := fmt.Sprintf("INSERT INTO %s (actor_id, film_id) VALUES ($1, $2) RETURNING id", actorFilmTbl)
		_, err = tx.Exec(createActorFilm, actorId, filmId)
		if err != nil {
			tx.Rollback()
			logrus.Errorf("error while exec in actors films database : [%v]", err)
			return -1, err
		}
	}
	return filmId, tx.Commit()
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

func (fdb *FilmDataBase) GetAllFilms() ([]restapi.Film, error) {
	var toReturn []restapi.Film
	tx, err := fdb.db.Begin()
	if err != nil {
		return nil, err
	}
	queryAllFilms := fmt.Sprintf("SELECT * FROM %s", filmTbl)

	rows, err := tx.Query(queryAllFilms)
	defer rows.Close()
	for rows.Next() {
		var tmp restapi.Film
		var (
			filmId          int
			filmTitle       string
			filmDescription string
			filmDate        string
			filmRate        string
		)
		err := rows.Scan(&filmId, &filmTitle, &filmDescription, &filmDate, &filmRate)

		tmp.FilmId = filmId
		tmp.Title = filmTitle
		tmp.Description = filmDescription
		tmp.Date = filmDate
		tmp.Rate = filmRate

		if err != nil {
			tx.Rollback()
			logrus.Errorf("error while scan film data to film struct : [%v]", err)
			return nil, err
		}
		toReturn = append(toReturn, tmp)
	}
	logrus.Infof("films : %v", toReturn)
	for idx, film := range toReturn {
		var actors []restapi.Actor

		query := fmt.Sprintf("SELECT actor_id, actor_name, actor_surname, actor_birth_date FROM %s act_fil INNER JOIN %s act ON act_fil.a_id = act.actor_id WHERE act_fil.f_id=$1", actorFilmTbl, actorTbl)
		err := fdb.db.Select(&actors, query, film.FilmId)
		if err != nil {
			logrus.Errorf("error while select from db")
			return nil, err
		}
		for idx, actor := range actors {
			films, err := fdb.ActorFilms(actor.ActorId)
			if err != nil {
				logrus.Errorf("error get films for actor with id %d : [%v]", actor.ActorId, err)
				return nil, errors.New("unknown error")
			}

			actors[idx].Films = films
		}
		toReturn[idx].Actors = actors
	}

	return toReturn, tx.Commit()
}

func (fdb *FilmDataBase) AddActorToFilm(actorId int, filmId int) error {
	var actor restapi.Actor
	tx, err := fdb.db.Begin()
	if err != nil {
		logrus.Errorf("error start tx : [%v]", err)
		return err
	}
	logrus.Infof("actorId : %d", actorId)
	query := fmt.Sprintf("SELECT * FROM %s WHERE actor_id=$1", actorTbl)
	row := tx.QueryRow(query, actorId)
	err = row.Scan(&actor.ActorId, &actor.FirstName, &actor.LastName, &actor.DateBirth)
	if err != nil {
		logrus.Errorf("error scan actor to actor struct : [%v]", err)
		return errors.New("something went wrong")
	}
	query = fmt.Sprintf("INSERT INTO %s (a_id, f_id) VALUES ($1, $2)", actorFilmTbl)
	_, err = tx.Exec(query, actorId, filmId)
	if err != nil {
		logrus.Errorf("error exec actor to databse")
		return err
	}

	return tx.Commit()
}

func (fdb *FilmDataBase) ActorFilms(actorId int) ([]restapi.Film, error) {
	var toReturn []restapi.Film
	tx, err := fdb.db.Begin()
	if err != nil {
		tx.Rollback()
		logrus.Errorf("error start transaction : [%v]", err)
		return nil, errors.New("error start transaction")
	}
	query := fmt.Sprintf("SELECT film_id, film_title, film_date, film_rate FROM %s act_fil INNER JOIN %s f ON act_fil.f_id = f.film_id WHERE act_fil.a_id=$1", actorFilmTbl, filmTbl)
	err = fdb.db.Select(&toReturn, query, actorId)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("error select rows from database : [%v]", err)
		return nil, errors.New("bad crud operation")
	}

	return toReturn, tx.Commit()
}
