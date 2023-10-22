package author

import (
	"advanced-rest-yt/pkg/logging"
	"context"
)

type Service struct {
	Repository Repository
	Logger     *logging.Logger
}

func (s *Service) GetAll(ctx context.Context) ([]Author, error) {
	all, err := s.Repository.FindAll(ctx)
	if err != nil {
		s.Logger.Error("cant find all, err: %s", err)
		return nil, err
	}

	return all, nil
}
