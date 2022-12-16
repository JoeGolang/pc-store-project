package couponhandler

import (
	"context"
	_interface "pc-shop-final-project/domain/repository"
)

type CouponHandler struct {
	ctx        context.Context
	repoCoupon _interface.InterfaceCoupon
}

func NewCouponHandler(ctx context.Context, repoCoupon _interface.InterfaceCustomer) *CouponHandler {
	return &CouponHandler{
		ctx:        ctx,
		repoCoupon: repoCoupon,
	}
}

//var (
//	repoCouponMysql = mysql.NewCouponMysql(mysqlCon.InitMysqlDB())
//	HandlerCoupon   = NewCouponHandler(repoCouponMysql)
//)
//
//type CouponInteractor struct {
//	repository _interface.InterfaceCoupon
//}
//
//func NewCouponHandler(Repo _interface.InterfaceCoupon) *CouponInteractor {
//	return &CouponInteractor{
//		repository: Repo,
//	}
//}
//
//func CreateCoupon(ctx context.Context, coupon *coupon.Coupon, UC []*coupon.UniqCoupon) {
//	err := HandlerCoupon.repository.CreateCoupon(ctx, coupon, UC)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func ReadCoupon(ctx context.Context) ([]*coupon.Coupon, []*coupon.UniqCoupon) {
//	Coupons, UCs, err := HandlerCoupon.repository.ReadCoupon(ctx)
//	if err != nil {
//		fmt.Println(err)
//		return nil, nil
//	}
//	return Coupons, UCs
//}
//
//func UpdateCoupon(ctx context.Context, uniqId string) {
//	err := HandlerCoupon.repository.UpdateCoupon(ctx, uniqId)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func DeleteCoupon(ctx context.Context, code int) {
//	err := HandlerCoupon.repository.DeleteCoupon(ctx, code)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
