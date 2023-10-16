package user

import (
	"advanced-rest-yt/pkg/logging"
	"context"
)

type Service struct {
	Storage Repository
	Logger  *logging.Logger
}

func Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	return
}
