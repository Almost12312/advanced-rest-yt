package book

import (
	"advanced-rest-yt/internal/author"
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

func (r *repository) FindAll(ctx context.Context) (a []author.Author, err error) {
	q := `select id, name from book;`

	qFormat := formatQuery(q)

	r.logger.Tracef("Sql query: %s", qFormat)

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bk book.Book
		err = rows.Scan(&bk.ID, &bk.Name)
		if err != nil {
			return nil, err
		}

		subq := `
		select a.name,
       		author_id,
       		book_id
		from public.authors_books
        join public.author a on a.id = authors_books.author_id
		where book_id = $1;`

		authorsRows, err := r.client.Query(ctx, subq, "16a7fb15-6289-46c9-8738-863ea6292d6f")
		if err != nil {
			return nil, err
		}

		for authorsRows.Next() {
			var auth author.Author

			err := rows.Scan(&auth.ID, &auth.Name)
			if err != nil {
				return nil, err
			}
			a = append(a, auth)
		}

		bk.Authors = a
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return a, nil
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
