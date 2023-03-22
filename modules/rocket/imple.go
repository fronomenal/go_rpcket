package rocket

import "context"

func GetService(repo Repository) Service {
	return Service{
		Repo: repo,
	}
}

func (s Service) GetRocketByID(ctx context.Context, id int32) (Rocket, error) {
	roc, err := s.Repo.GetByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return roc, nil
}

func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	roc, err := s.Repo.Insert(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return roc, nil
}

func (s Service) RemoveRocket(ctx context.Context, id int32) error {
	err := s.Repo.Remove(id)
	return err
}
