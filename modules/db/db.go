package db

import (
	"github.com/fronomenal/go_rpcket/modules/rocket"
	"github.com/jmoiron/sqlx"
)

type Pool struct {
	db *sqlx.DB
}

func Conn() (Pool, error) {
	db, err := sqlx.Connect("sqlite3", "./rockets.db")
	if err != nil {
		return Pool{db: db}, err
	}
	return Pool{db: db}, nil
}

func (p Pool) GetByID(id int32) (rocket.Rocket, error) {
	var roc rocket.Rocket

	row := p.db.QueryRow(`SELECT * FROM rockets WHERE id=$1;`, id)
	if err := row.Scan(&roc); err != nil {
		return rocket.Rocket{}, err
	}

	return roc, nil
}

func (p Pool) Insert(roc rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (p Pool) Remove(id int32) error {
	return nil
}
