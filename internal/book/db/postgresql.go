package book

import (
	"advanced-rest-yt/internal/author/model"
	"advanced-rest-yt/internal/book"
	"advanced-rest-yt/pkg/client/postgresql"
	"advanced-rest-yt/pkg/logging"
	"context"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) FindAll(ctx context.Context) (books []book.Book, err error) {
	q := `select id, name, age from book;`

	qFormat := formatQuery(q)

	r.logger.Tracef("Sql query: %s", qFormat)

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bk Book
		err = rows.Scan(&bk.ID, &bk.Name, &bk.Age)
		if err != nil {
			return nil, err
		}

		subq := `
		select a.name,
       		a.id
		from public.authors_books ab
        join public.author a on a.id = ab.author_id
		where ab.book_id = $1;`

		//s := `-- select a.id, a.name from public.authors_books ab join public.author a on a.id = ab.author_id where book_id = $1`

		authorsRows, err := r.client.Query(ctx, subq, bk.ID)
		if err != nil {
			return nil, err
		}

		authors := make([]model.Author, 0)

		for authorsRows.Next() {
			var auth model.Author

			err := authorsRows.Scan(&auth.ID, &auth.Name)
			if err != nil {
				return nil, err
			}
			authors = append(authors, auth)
		}

		bk.Authors = authors

		books = append(books, bk.ToDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
