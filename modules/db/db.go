package db

import (
	"github.com/fronomenal/go_rpcket/modules/rocket"
	"github.com/jmoiron/sqlx"
)

type Pool struct {
	db *sqlx.DB
}

func Conn() (Pool, error) {
	db, err := sqlx.Connect("sqlite3", "../rockets.db")
	if err != nil {
		return Pool{db: db}, err
	}
	return Pool{db: db}, nil
}

func (p Pool) GetByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (p Pool) Insert(roc rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (p Pool) Remove(id string) error {
	return nil
}
