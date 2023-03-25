package grpc

import (
	"context"
	"log"

	"github.com/fronomenal/go_rpcket/modules/rocket"
	roc "github.com/fronomenal/go_rpcket/protos/v1"
)

func (h Handler) GetRocket(ctx context.Context, req *roc.GetReq) (*roc.GetRes, error) {
	log.Print("In Get Rocket Endpoint")

	rocket, err := h.RocketService.GetRocketByID(ctx, req.Id)
	if err != nil {
		return &roc.GetRes{}, err
	}

	return &roc.GetRes{Rocket: &roc.Rocket{Id: rocket.ID, Name: rocket.Name, Type: rocket.Type, Flights: int32(rocket.Flights)}}, nil
}

func (h Handler) SetRocket(ctx context.Context, req *roc.SetReq) (*roc.SetRes, error) {
	log.Print("In Set Rocket Endpoint")

	rocket, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{ID: req.Rocket.Id, Type: req.Rocket.Type, Name: req.Rocket.Name, Flights: int(req.Rocket.Flights)})
	if err != nil {
		log.Println("Failed inserting rocket into database")
		return &roc.SetRes{}, err
	}

	return &roc.SetRes{Rocket: &roc.Rocket{Id: rocket.ID, Name: rocket.Name, Type: rocket.Type, Flights: int32(rocket.Flights)}}, nil
}

func (h Handler) RemRocket(ctx context.Context, req *roc.RemReq) (*roc.RemRes, error) {
	return &roc.RemRes{}, nil
}
