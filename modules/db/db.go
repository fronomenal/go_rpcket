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
	if err := row.Scan(&roc.ID, &roc.Type, &roc.Name, &roc.Flights); err != nil {
		return rocket.Rocket{}, err
	}

	return roc, nil
}

func (p Pool) Insert(roc rocket.Rocket) (rocket.Rocket, error) {

	rows, err := p.db.NamedQuery(`INSERT INTO rockets(rkt_type, rkt_name, flights) VALUES (:type, :name, :flights) RETURNING *;`, roc)
	if err != nil {
		return rocket.Rocket{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return rocket.Rocket{}, err
	}

	if err := rows.Scan(&roc.ID, &roc.Type, &roc.Name, &roc.Flights); err != nil {
		return roc, err
	}

	return roc, nil
}

func (p Pool) Remove(id int32) error {

	_, err := p.db.Exec(`DELETE FROM rockets WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	return nil
}
