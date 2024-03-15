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

func (adb *ActorDataBase) DeleteActor(actorId int) error {
	return nil
}

func (adb *ActorDataBase) ChangeActor(actorId int, toChange string) error {
	return nil
}

func (adb *ActorDataBase) FindActorFilm(actor string) ([]restapi.Film, error) {
	return nil, nil
}
