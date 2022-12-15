package repository

import (
	"context"
	"pc-shop-final-project/domain/entity/user"
)

type InterfaceUser interface {
	InsertDataUser(ctx context.Context, dataUser *user.User) error
	GetListUser(ctx context.Context) ([]*user.User, error)
	GetUserById(ctx context.Context, id string) (*user.User, error)
	UpdateUserById(ctx context.Context, dataUser *user.User, idUser string) error
	DeleteUserById(ctx context.Context, id string) error
}
