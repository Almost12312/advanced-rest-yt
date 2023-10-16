package author

import "context"

type Repository interface {
	Create(ctx context.Context, author Author) (string, error)
	FindOne(ctx context.Context, id string) (a Author, err error)
	Update(ctx context.Context, author Author) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) (a []Author, err error)
}
