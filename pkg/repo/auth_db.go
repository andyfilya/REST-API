package repo

import (
	"errors"
	"fmt"
	"github.com/andyfilya/restapi"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	usrTbl = "users"
)

type AuthDataBase struct {
	db *sqlx.DB
}

func InitAuthDataBase(db *sqlx.DB) *AuthDataBase {
	return &AuthDataBase{
		db: db,
	}
}

func (auth *AuthDataBase) NewUser(user restapi.User) (int, error) {
	var lastInsertId int
	query := fmt.Sprintf("INSERT INTO %s (create_time, username, password) VALUES (NOW(), $1, $2) RETURNING user_id", usrTbl)
	logrus.Infof("%s", query, user.Username, user.Password)
	row := auth.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&lastInsertId); err != nil {
		logrus.Errorf("error while query row to postgres : [%v]", err)
		return -1, err
	}
	return lastInsertId, nil
}

func (auth *AuthDataBase) FindUser(username, password string) (restapi.User, error) {
	findUsr := restapi.User{}
	var timestamp interface{}

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", usrTbl)
	row := auth.db.QueryRow(query, username)
	err := row.Scan(&findUsr.Id, &findUsr.Username, &findUsr.Password, &timestamp)
	if err != nil {
		logrus.Errorf("error while scan row from postgreswql : [%v]", err)
		return findUsr, errors.New("can't find user with username: " + username)
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUsr.Password), []byte(password))
	return findUsr, err
}
