package repository

import (
	"context"
	"pc-shop-final-project/domain/entity/user"
)

type InterfaceUser interface {
	CreateUser(ctx context.Context, user *user.User) error
	ReadUser(ctx context.Context) ([]*user.User, error)
	DeleteUser(ctx context.Context, id int) error
}
