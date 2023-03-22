package main

import (
	"log"

	"github.com/fronomenal/go_rpcket/modules/db"
	"github.com/fronomenal/go_rpcket/modules/rocket"
	"github.com/fronomenal/go_rpcket/modules/transport/grpc"
)

func Start(port int) error {
	dbpool, err := db.Conn()
	if err != nil {
		return err
	}

	if err := dbpool.Migrate(); err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	grpc.Port = port
	rocService := rocket.GetService(dbpool)
	rocHandler := grpc.GetHandler(rocService)

	if err := rocHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Start(51515); err != nil {
		log.Fatal(err)
	}
}
