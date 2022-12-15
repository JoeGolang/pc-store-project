package repository

import (
	"context"
	user2 "pc-shop-final-project/domain/entity/user"
)

type InterfaceUser interface {
	InsertDataUser(ctx context.Context, dataUser *user2.User) error
	GetListUser(ctx context.Context) ([]*user2.User, error)
	GetUserById(ctx context.Context, id string) (*user2.User, error)
	UpdateUserById(ctx context.Context, dataUser *user2.User, idUser string) error
	DeleteUserById(ctx context.Context, idUser string) error
}
