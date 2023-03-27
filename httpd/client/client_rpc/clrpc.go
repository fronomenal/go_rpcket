package clrpc

import (
	"context"
	"flag"
	"time"

	roc "github.com/fronomenal/go_rpcket/protos/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Rarg struct {
	Id       int32  `json:"-"`
	Name     string `json:"name"`
	Rkt_type string `json:"type"`
	Flights  int32  `json:"flights"`
	Valid    bool
}

var (
	addr   = flag.String("addr", "localhost:51515", "the listening address of the client")
	client roc.RocketServiceClient
)

func Connect() (*grpc.ClientConn, error) {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client = roc.NewRocketServiceClient(conn)
	return conn, nil
}

func Get(args *Rarg) (*roc.Rocket, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := client.GetRocket(ctx, &roc.GetReq{Id: args.Id})
	if err != nil {
		return &roc.Rocket{}, err
	}

	return r.Rocket, nil
}

func Set(args *Rarg) (*roc.Rocket, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := client.SetRocket(ctx, &roc.SetReq{Rocket: &roc.Rocket{Name: args.Name, Type: args.Rkt_type, Flights: args.Flights}})
	if err != nil {
		return &roc.Rocket{}, err
	}

	return r.Rocket, nil
}

func Rem(args *Rarg) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := client.RemRocket(ctx, &roc.RemReq{Id: args.Id})
	if err != nil {
		return "", err
	}

	return r.Status, nil
}
