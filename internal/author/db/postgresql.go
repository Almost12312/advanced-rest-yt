package db

import (
	"advanced-rest-yt/internal/author"
	"advanced-rest-yt/pkg/client/postgresql"
	"advanced-rest-yt/pkg/logging"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, author author.Author) (string, error) {
	q := `
		insert into author (name) 
		values ($name) 
		returning id
	`

	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			errMsg := fmt.Sprintf("sql error: %s. Detail: %s. Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			r.logger.Error(errMsg)
		}
	}
}

func (r *repository) FindOne(ctx context.Context, id string) (a author.Author, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, author author.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindAll(ctx context.Context) (a []author.Author, err error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return repository{
		client: client,
		logger: logger,
	}
}
