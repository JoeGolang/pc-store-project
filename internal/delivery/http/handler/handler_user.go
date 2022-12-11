package handler

import (
	"context"
	"fmt"
	user2 "pc-shop-final-project/domain/entity/user"
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

func CreateUser(ctx context.Context, user *user2.User) {
	err := HandlerUser.repository.CreateUser(ctx, user)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadUser(ctx context.Context) []*user2.User {
	Users, err := HandlerUser.repository.ReadUser(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return Users
}

func DeleteUser(ctx context.Context, code int) {
	err := HandlerUser.repository.DeleteUser(ctx, code)
	if err != nil {
		fmt.Println(err)
	}
}
