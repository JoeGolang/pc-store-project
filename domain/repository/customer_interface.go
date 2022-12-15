package repository

import (
	"context"
	"pc-shop-final-project/domain/entity/customer"
)

type InterfaceCustomer interface {
	InsertDataCustomer(ctx context.Context, cust *customer.Customer) error
	GetListCustomer(ctx context.Context) ([]*customer.Customer, error)
	DeleteCustomer(ctx context.Context, uniqId string) error
	GetCustomerById(ctx context.Context, id string) (*customer.Customer, error)
}
