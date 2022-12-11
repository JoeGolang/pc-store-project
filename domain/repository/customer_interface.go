package repository

import (
	"context"
	"pc-shop-final-project/domain/entity/customer"
)

type InterfaceCustomer interface {
	CreateCustomer(ctx context.Context, cust *customer.Customer) error
	ReadCustomer(ctx context.Context) ([]*customer.Customer, error)
	DeleteCustomer(ctx context.Context, uniqId string) error
}
