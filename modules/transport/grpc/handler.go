package grpc

import (
	"log"
	"net"

	"github.com/fronomenal/go_rpcket/modules/rocket"
	roc "github.com/fronomenal/go_rpcket/protos/v1"
	"google.golang.org/grpc"
)

type Handler struct {
	RocketService rocket.RocketService
	roc.UnimplementedRocketServiceServer
}

func GetHandler(rocService rocket.RocketService) Handler {
	return Handler{RocketService: rocService}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":5151")
	if err != nil {
		log.Print("Failed to listen on port 5151")
		return err
	}

	rpcServer := grpc.NewServer()
	roc.RegisterRocketServiceServer(rpcServer, &h)

	if err := rpcServer.Serve(lis); err != nil {
		log.Printf("Failed to server: %s\n", err)
		return err
	}

	return nil
}
