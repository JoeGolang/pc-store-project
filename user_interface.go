package _interface

import (
	"context"
	"pc-shop-final-project/domain/entity"
)

type InterfaceUser interface {
	CreateUser(ctx context.Context, user *entity.User) error
	ReadUser(ctx context.Context) ([]*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}
