package grpc

import (
	"context"

	roc "github.com/fronomenal/go_rpcket/protos/v1"
)

func (h Handler) GetRocket(ctx context.Context, req *roc.GetReq) (*roc.GetRes, error) {
	return &roc.GetRes{}, nil
}

func (h Handler) SetRocket(ctx context.Context, req *roc.SetReq) (*roc.SetRes, error) {
	return &roc.SetRes{}, nil
}

func (h Handler) RemRocket(ctx context.Context, req *roc.RemReq) (*roc.RemRes, error) {
	return &roc.RemRes{}, nil
}
