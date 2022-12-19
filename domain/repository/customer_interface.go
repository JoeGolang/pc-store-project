package _interface

import (
	"context"
	"pc-shop-final-project/domain/entity"
)

type InterfaceCustomer interface {
	CreateCustomer(ctx context.Context, cust *entity.Customer) error
	ReadCustomer(ctx context.Context) ([]*entity.Customer, error)
	DeleteCustomer(ctx context.Context, uniqId string) error
}
