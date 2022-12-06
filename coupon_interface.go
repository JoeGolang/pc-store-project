package _interface

import (
	"context"
	"pc-shop-final-project/domain/entity"
)

type InterfaceCoupon interface {
	CreateCoupon(ctx context.Context, coupon *entity.Coupon, uniqcoupon []*entity.UniqCoupon) error
	ReadCoupon(ctx context.Context) ([]*entity.Coupon, []*entity.UniqCoupon, error)
	UpdateCoupon(ctx context.Context, uniqId string) error
	DeleteCoupon(ctx context.Context, id int) error
}
