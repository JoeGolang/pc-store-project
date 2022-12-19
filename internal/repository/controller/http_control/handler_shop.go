package handlerHTTP

import (
	"context"
	_interface "pc-shop-final-project/domain/repository"
)

type UserHandler struct {
	ctx      context.Context
	repoUser _interface.InterfaceUser
}
type CustomerHandler struct {
	ctx      context.Context
	repoCust _interface.InterfaceCustomer
}
type InventoryHandler struct {
	ctx       context.Context
	repoInven _interface.InterfaceInventory
}
type CouponHandler struct {
	ctx        context.Context
	repoCoupon _interface.InterfaceCoupon
}
type SettlementHandler struct {
	ctx      context.Context
	repoSett _interface.InterfaceSettlement
}
type SettlementPurchaseHandler struct {
	ctx      context.Context
	repoSetP _interface.InterfaceSettlementItem
}

func NewUserHandler(ctx context.Context, repoUser _interface.InterfaceUser) *UserHandler {
	return &UserHandler{
		ctx:      ctx,
		repoUser: repoUser,
	}
}

func NewCustomerHandler(ctx context.Context, repoCust _interface.InterfaceCustomer) *CustomerHandler {
	return &CustomerHandler{
		ctx:      ctx,
		repoCust: repoCust,
	}
}

func NewInventoryHandler(ctx context.Context, repoInven _interface.InterfaceInventory) *InventoryHandler {
	return &InventoryHandler{
		ctx:       ctx,
		repoInven: repoInven,
	}
}

func NewCouponHandler(ctx context.Context, repoCoupon _interface.InterfaceCoupon) *CouponHandler {
	return &CouponHandler{
		ctx:        ctx,
		repoCoupon: repoCoupon,
	}
}

func NewSettlementHandler(ctx context.Context, repoSett _interface.InterfaceSettlement) *SettlementHandler {
	return &SettlementHandler{
		ctx:      ctx,
		repoSett: repoSett,
	}
}

func NewSettlePurchaseHandler(ctx context.Context, repoSetP _interface.InterfaceSettlementItem) *SettlementPurchaseHandler {
	return &SettlementPurchaseHandler{
		ctx:      ctx,
		repoSetP: repoSetP,
	}
}
