package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/fronomenal/go_rpcket/modules/rocket"
	roc "github.com/fronomenal/go_rpcket/protos/v1"
	"google.golang.org/grpc"
)

var Port int

type Handler struct {
	RocketService rocket.RocketService
	roc.UnimplementedRocketServiceServer
}

func GetHandler(rocService rocket.RocketService) Handler {
	return Handler{RocketService: rocService}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		log.Printf("Failed to listen on port %d", Port)
		return err
	}

	rpcServer := grpc.NewServer()
	roc.RegisterRocketServiceServer(rpcServer, &h)
	log.Printf("Listening on Port: %d\n", Port)

	if err := rpcServer.Serve(lis); err != nil {
		log.Printf("Failed to server: %s\n", err)
		return err
	}

	return nil
}
