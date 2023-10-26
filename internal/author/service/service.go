package service

import (
	"advanced-rest-yt/internal/author/model"
	"advanced-rest-yt/internal/author/storage"
	model2 "advanced-rest-yt/internal/author/storage/model"
	"advanced-rest-yt/pkg/api/filter"
	"advanced-rest-yt/pkg/api/sort"
	"advanced-rest-yt/pkg/logging"
	"context"
)

type Service struct {
	Repository storage.Repository
	Logger     *logging.Logger
}

func NewService(repository storage.Repository, logger *logging.Logger) *Service {
	return &Service{
		Repository: repository,
		Logger:     logger,
	}
}

func (s *Service) GetAll(ctx context.Context, sortOptions sort.Options, filterOptions filter.Options) ([]model.Author, error) {
	opt := model2.NewSortOptions(sortOptions.Field, sortOptions.Order)

	all, err := s.Repository.FindAll(ctx, opt)
	if err != nil {
		s.Logger.Errorf("cant find all, err: %s", err)
		return nil, err
	}

	return all, nil
}
