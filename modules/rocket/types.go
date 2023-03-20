//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/fronomenal/go_rpcket/modules/rocket Repository

package rocket

import "context"

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

type Service struct {
	Repo Repository
}

type Repository interface {
	GetByID(id string) (Rocket, error)
	Insert(roc Rocket) (Rocket, error)
	Remove(id string) error
}

type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (Rocket, error)
	InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error)
	RemoveRocket(ctx context.Context, id string) error
}
