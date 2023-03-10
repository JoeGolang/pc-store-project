package _interface

import (
	"context"
	"pc-shop-final-project/domain/entity"
)

type InterfaceSettlement interface {
	CreateSettle(ctx context.Context, idUser int, idCustomer string, idCoupon int, settle *entity.Settlement) error
	ReadSettle(ctx context.Context) ([]int, []string, []int, []*entity.Settlement, error)
	UpdateSettle(ctx context.Context, code string) error
	DeleteSettle(ctx context.Context, code string) error
}

type InterfaceSettlementItem interface {
	CreatePurchase(ctx context.Context, purchase []*entity.SettlePurchase) error
	ReadPurchase(ctx context.Context) ([]*entity.SettlePurchase, error)
	DeletePurchase(ctx context.Context, codeItem string) error
}
