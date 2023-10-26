package postgresql

import (
	"advanced-rest-yt/internal/author/model"
	"advanced-rest-yt/internal/author/storage"
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"

	"advanced-rest-yt/pkg/client/postgresql"
	"advanced-rest-yt/pkg/logging"

	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) storage.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, author *model.Author) (string, error) {
	q := `
		insert into public.author (name) 
		values ($1) 
		returning id;
	`

	qFormat := formatQuery(q)

	r.logger.Tracef("Sql query: %s", qFormat)

	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
			errMsg := fmt.Sprintf("sql error: %s. Code: %s, Detail: %s. Where: %s", pgErr.Message, pgErr.Code, pgErr.Detail, pgErr.Where)
			r.logger.Error(errMsg)
		}
		return "", err
	}

	return author.ID, nil
}

func (r *repository) FindAll(ctx context.Context, options storage.SortOptions) (a []model.Author, err error) {
	q := sq.Select("id, name, age, is_alive, created_at").From("public.author")

	if options != nil {
		q = q.OrderBy(options.GetOrderBy())
	}

	sql, i, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	qFormat := formatQuery(sql)

	r.logger.Tracef("Sql query: %s", qFormat)

	rows, err := r.client.Query(ctx, sql, i...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ath Author
		err = rows.Scan(&ath.ID, &ath.Name, &ath.Age, &ath.IsAlive, &ath.CreatedAt)
		if err != nil {
			return nil, err
		}

		a = append(a, ath.ToDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return a, nil
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) FindOne(ctx context.Context, id string) (a model.Author, err error) {
	q := `select id,name from public.author where id = $1;`

	qFormat := formatQuery(q)

	r.logger.Tracef("Sql query: %s", qFormat)
	err = r.client.QueryRow(ctx, q, id).Scan(&a.ID, &a.Name)

	return a, nil
}

func (r *repository) Update(ctx context.Context, author model.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//q := `DELETE FROM author WHERE id = $id;`
	//
	//qFormat := formatQuery(q)
	//r.logger.Tracef("Sql query %s", qFormat)
	//
	//row, _ := r.client.Query(ctx,q,id)
	//row.

	panic("implement me!")
}
