package book

import (
	"advanced-rest-yt/internal/author"
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) (a []author.Author, err error)
}
