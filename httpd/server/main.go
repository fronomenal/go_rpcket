package main

import (
	"log"

	"github.com/fronomenal/go_rpcket/modules/db"
	"github.com/fronomenal/go_rpcket/modules/rocket"
)

func Start() error {
	dbpool, err := db.Conn()
	if err != nil {
		return err
	}

	_ = rocket.GetService(dbpool)

	return nil
}

func main() {
	if err := Start(); err != nil {
		log.Fatal(err)
	}
}
