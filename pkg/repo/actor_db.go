package repo

import (
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
	query := fmt.Sprintf("DELETE FROM %s WHERE actor_name=$1 AND actor_surname=$2", actorTbl)
	_, err := adb.db.Exec(query, actor.FirstName, actor.LastName)
	if err != nil {
		logrus.Errorf("error while exec actor table : [%v]", err)
		return err
	}
	return nil
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

func (adb *ActorDataBase) FindActorFilm(actor string) ([]restapi.Film, error) {
	return nil, nil
}
