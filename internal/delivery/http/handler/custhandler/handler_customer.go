package handler

import (
	"context"
	_interface "pc-shop-final-project/domain/repository"
)

type CustomerHandler struct {
	ctx      context.Context
	repoCust _interface.InterfaceUser
}

func NewCustomerHandler(ctx context.Context, repoCust _interface.InterfaceCustomer) *CustomerHandler {
	return &CustomerHandler{
		ctx:      ctx,
		repoCust: repoCust,
	}
}

//func CreateCustomer(ctx context.Context, cust *customer.Customer) {
//	err := HandlerCustomer.repository.CreateCustomer(ctx, cust)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func ReadCustomer(ctx context.Context) []*customer.Customer {
//	customers, err := HandlerCustomer.repository.ReadCustomer(ctx)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	return customers
//}
//
//func DeleteCustomer(ctx context.Context, uniqId string) {
//	err := HandlerCustomer.repository.DeleteCustomer(ctx, uniqId)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
