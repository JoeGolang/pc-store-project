package custhandler

import (
	"context"
	_interface "pc-shop-final-project/domain/repository"
)

type CustomerHandler struct {
	ctx          context.Context
	repoCustomer _interface.InterfaceCustomer
}

func NewCustomerHandler(ctx context.Context, repoCustomer _interface.InterfaceCustomer) *CustomerHandler {
	return &CustomerHandler{
		ctx:          ctx,
		repoCustomer: repoCustomer,
	}
}
