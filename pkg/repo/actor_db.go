package repo

import (
	"errors"
	"fmt"
	"github.com/andyfilya/restapi"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	actorTbl = "actors"
)

type ActorDataBase struct {
	db *sqlx.DB
}

func InitActorDataBase(db *sqlx.DB) *ActorDataBase {
	return &ActorDataBase{
		db: db,
	}
}

func (adb *ActorDataBase) CreateActor(actor restapi.Actor) (int, error) {
	var actorId int
	query := fmt.Sprintf("INSERT INTO %s (actor_name, actor_surname, actor_birth_date) VALUES ($1, $2, $3) RETURNING actor_id", actorTbl)
	row := adb.db.QueryRow(query, actor.FirstName, actor.LastName, actor.DateBirth)
	err := row.Scan(&actorId)
	if err != nil {
		logrus.Errorf("error while scan actor id in var : [%v]", err)
		return -1, err
	}
	return actorId, nil
}

func (adb *ActorDataBase) DeleteActor(actor restapi.Actor) error {
	var actorId int
	tx, err := adb.db.Begin()
	if err != nil {
		logrus.Errorf("error begin transaction : [%v]", err)
		return errors.New("unknown error")
	}
	query := fmt.Sprintf("SELECT actor_id FROM %s WHERE actor_name=$1 AND actor_surname=$2", actorTbl)
	row := tx.QueryRow(query, actor.FirstName, actor.LastName)
	err = row.Scan(&actorId)

	if err != nil {
		logrus.Errorf("error in tx : [%v]", err)
		return errors.New("unknown error")
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE actor_id=$1", actorTbl)
	_, err = tx.Exec(query, actorId)
	if err != nil {
		logrus.Errorf("error while exec actor table : [%v]", err)
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE a_id=$1", actorFilmTbl)
	_, err = tx.Exec(query, actorId)
	if err != nil {
		logrus.Errorf("error while exec actor table : [%v]", err)
		return err
	}

	return tx.Commit()
}

func (adb *ActorDataBase) ChangeActor(oldActor restapi.Actor, newActor restapi.Actor) error {
	query := fmt.Sprintf("UPDATE %s SET actor_name=$1, actor_surname=$2, actor_birth_date=$3 WHERE actor_name=$4 AND actor_surname=$5", actorTbl)
	_, err := adb.db.Exec(query, newActor.FirstName, newActor.LastName, newActor.DateBirth, oldActor.FirstName, oldActor.LastName)
	if err != nil {
		logrus.Errorf("error while exec update table : [%v]", err)
		return err
	}
	return nil
}

func (adb *ActorDataBase) FindActorFilm(actorFragments restapi.ActorFragment) ([]restapi.Film, error) {
	var toReturn []restapi.Film
	tx, err := adb.db.Begin()
	if err != nil {
		tx.Rollback()
		logrus.Errorf("error start transaction : [%v]", err)
		return nil, errors.New("error start transaction")
	}

	query := fmt.Sprintf("SELECT film_id, film_title, film_date, film_description FROM films f INNER JOIN actors_films act_fil ON act_fil.f_id = f.film_id INNER JOIN actors  a ON act_fil.a_id = a.actor_id WHERE actor_name LIKE CONCAT($1::text,'%%') AND actor_surname LIKE CONCAT($2::text,'%%')")
	err = adb.db.Select(&toReturn, query, actorFragments.ActorNameFragment, actorFragments.ActorSurnameFragment)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("error select rows from database : [%v]", err)
		return nil, errors.New("bad crud operation")
	}

	return toReturn, tx.Commit()
}
