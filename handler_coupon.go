package handler

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity"
	_interface "pc-shop-final-project/domain/repository"
	"pc-shop-final-project/internal/repository/mysql"
)

var (
	repoCouponMysql = mysql.NewCouponMysql(mysqlConnection)
	HandlerCoupon   = NewCouponHandler(repoCouponMysql)
)

type CouponInteractor struct {
	repository _interface.InterfaceCoupon
}

func NewCouponHandler(Repo _interface.InterfaceCoupon) *CouponInteractor {
	return &CouponInteractor{
		repository: Repo,
	}
}

func CreateCoupon(ctx context.Context, coupon *entity.Coupon, UC []*entity.UniqCoupon) {
	err := HandlerCoupon.repository.CreateCoupon(ctx, coupon, UC)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadCoupon(ctx context.Context) ([]*entity.Coupon, []*entity.UniqCoupon) {
	Coupons, UCs, err := HandlerCoupon.repository.ReadCoupon(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return Coupons, UCs
}

func UpdateCoupon(ctx context.Context, uniqId string) {
	err := HandlerCoupon.repository.UpdateCoupon(ctx, uniqId)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteCoupon(ctx context.Context, code int) {
	err := HandlerCoupon.repository.DeleteCoupon(ctx, code)
	if err != nil {
		fmt.Println(err)
	}
}
