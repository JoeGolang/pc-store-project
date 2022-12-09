package repository

import (
	"context"
	"pc-shop-final-project/domain/entity/coupon"
)

type InterfaceCoupon interface {
	CreateCoupon(ctx context.Context, coupon *coupon.Coupon, uniqcoupon []*coupon.UniqCoupon) error
	ReadCoupon(ctx context.Context) ([]*coupon.Coupon, []*coupon.UniqCoupon, error)
	UpdateCoupon(ctx context.Context, uniqId string) error
	DeleteCoupon(ctx context.Context, id int) error
}
