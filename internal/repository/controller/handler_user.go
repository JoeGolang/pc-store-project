package handler

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity"
	_interface "pc-shop-final-project/domain/repository"
	"pc-shop-final-project/internal/repository/mysql"
)

var (
	repoUserMysql = mysql.NewUserMysql(mysqlConnection)
	HandlerUser   = NewUserHandler(repoUserMysql)
)

type UserInteractor struct {
	repository _interface.InterfaceUser
}

func NewUserHandler(Repo _interface.InterfaceUser) *UserInteractor {
	return &UserInteractor{
		repository: Repo,
	}
}

func CreateUser(ctx context.Context, user *entity.User) {
	err := HandlerUser.repository.CreateUser(ctx, user)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadUser(ctx context.Context) []*entity.User {
	Users, err := HandlerUser.repository.ReadUser(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return Users
}

func DeleteUser(ctx context.Context, code int) error {
	err := HandlerUser.repository.DeleteUser(ctx, code)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
