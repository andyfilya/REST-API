package repo

import (
	"fmt"
	"github.com/andyfilya/restapi"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	var lastInsertId int
	films := pq.Array(actor.Films)
	query := fmt.Sprintf("INSERT INTO %s (actor_name, actor_surname, actor_date_birth, actor_films) VALUES ($1, $2, $3, $4)", actorTbl)

	row := adb.db.QueryRow(query, actor.FirstName, actor.LastName, actor.DateBirth, films)
	if err := row.Scan(&lastInsertId); err != nil {
		logrus.Errorf("error while row scan : [%v]", err)
		return -1, err
	}

	return lastInsertId, nil
}
