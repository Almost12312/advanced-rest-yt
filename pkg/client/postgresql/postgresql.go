package postgresql

import (
	"advanced-rest-yt/pkg/repeatable"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

//type StorageConfig struct {
//	username, password, host, port, db string
//	maxAttempts                        int
//}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

// TODO: logger
func NewClient(ctx context.Context, maxAttempts uint, username, password, host, port, db string) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, db)

	err = repeatable.DoWithAttempts(func() error {
		ctxTimeout, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
		defer cancelFunc()

		pxpool, err := pgxpool.Connect(ctxTimeout, dsn)
		if err != nil {
			fmt.Println("Cant connect to postgres")
			return nil
		}
		pool = pxpool

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatalf("postgress cant connect, %s", err)
	}

	return pool, nil
}
