//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/fronomenal/go_rpcket/modules/rocket Repository

package rocket

import "context"

type Rocket struct {
	ID      int32
	Name    string
	Type    string
	Flights int
}

type Service struct {
	Repo Repository
}

type Repository interface {
	GetByID(id int32) (Rocket, error)
	Insert(roc Rocket) (Rocket, error)
	Remove(id int32) error
}

type RocketService interface {
	GetRocketByID(ctx context.Context, id int32) (Rocket, error)
	InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error)
	RemoveRocket(ctx context.Context, id int32) error
}
