package storage

import (
	"advanced-rest-yt/internal/author/model"
	"context"
)

type Repository interface {
	Create(ctx context.Context, author *model.Author) (string, error)
	FindOne(ctx context.Context, id string) (a model.Author, err error)
	Update(ctx context.Context, author model.Author) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, options SortOptions) (a []model.Author, err error)
}

type SortOptions interface {
	GetOrderBy() string
}
