package testdata_test

import (
	"fmt"
	"pc-shop-final-project/domain/entity/coupon"
	"testing"
)

var (
	randy = 12000000
	donny = 24000000
	rika  = 35000000
	monna = 2000000
)

func TestNewCoupon1(t *testing.T) {
	fmt.Println(coupon.AutoGenerateCoupon(randy, "Randy"))
}

func TestNewCoupon2(t *testing.T) {
	fmt.Println(coupon.AutoGenerateCoupon(donny, "Donny"))
}

func TestNewCoupon3(t *testing.T) {
	fmt.Println(coupon.AutoGenerateCoupon(rika, "Rika"))
}

func TestNewCoupon4(t *testing.T) {
	fmt.Println(coupon.AutoGenerateCoupon(monna, "Monna"))
}
