package repo

import (
	"fmt"
	"github.com/andyfilya/restapi/config"
	"github.com/jmoiron/sqlx"
)

func NewDataBase(dc *config.UserDatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		dc.Host, dc.Port, dc.Username, dc.DatabaseName, dc.SSLmode, dc.Password)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
