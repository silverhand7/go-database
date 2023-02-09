package repository

import (
	"context"
	"go-database/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	FindById(ctx context.Context, id int32) (entity.User, error)
	All(ctx context.Context) ([]entity.User, error)
}
