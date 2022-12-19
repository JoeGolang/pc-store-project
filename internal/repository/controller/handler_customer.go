package handler

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity"
	_interface "pc-shop-final-project/domain/repository"
	"pc-shop-final-project/internal/repository/mysql"
)

var (
	repoCustomerMysql = mysql.NewCustomerMysql(mysqlConnection)
	HandlerCustomer   = NewCustomerHandler(repoCustomerMysql)
)

type CustomerInteractor struct {
	repository _interface.InterfaceCustomer
}

func NewCustomerHandler(Repo _interface.InterfaceCustomer) *CustomerInteractor {
	return &CustomerInteractor{
		repository: Repo,
	}
}

func CreateCustomer(ctx context.Context, cust *entity.Customer) {
	err := HandlerCustomer.repository.CreateCustomer(ctx, cust)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadCustomer(ctx context.Context) []*entity.Customer {
	customers, err := HandlerCustomer.repository.ReadCustomer(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return customers
}

func DeleteCustomer(ctx context.Context, uniqId string) {
	err := HandlerCustomer.repository.DeleteCustomer(ctx, uniqId)
	if err != nil {
		fmt.Println(err)
	}
}
