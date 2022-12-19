package entity_test

import (
	"fmt"
	"pc-shop-final-project/domain/entity"
	"testing"
)

var (
	randy = 12000000
	donny = 24000000
	rika  = 35000000
	monna = 2000000
)

func TestNewCoupon1(t *testing.T) {
	fmt.Println(entity.AutoGenerateCoupon(1, randy, "Randy"))
}

func TestNewCoupon2(t *testing.T) {
	fmt.Println(entity.AutoGenerateCoupon(1, donny, "Donny"))
}

func TestNewCoupon3(t *testing.T) {
	fmt.Println(entity.AutoGenerateCoupon(1, rika, "Rika"))
}

func TestNewCoupon4(t *testing.T) {
	fmt.Println(entity.AutoGenerateCoupon(1, monna, "Monna"))
}
