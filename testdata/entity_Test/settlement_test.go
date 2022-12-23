package entity_Test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"pc-shop-final-project/domain/entity"
	"testing"
)

type TestSettlementValidate struct {
	id         int
	code       string
	user       entity.User
	customer   entity.Customer
	product    []*entity.SettlePurchase
	coupon     entity.Coupon
	totalPrice int
	statusTrns bool
	want       error
}

////// positif FLOW
////func TestPositiveNewInventory(t *testing.T) {
////	dataInventory, err := inventory.NewInventory(inventory.DTOInventory{
////		Id:          1234,
////		ProductName: "Switch",
////		Brand:       "Nintendo",
////		Price:       2500000,
////		Category:    "New Console",
////	})
////
////	if err != nil {
////		t.Error(err)
////	}
////
////	assert.Equal(t, "Switch", dataInventory.GetValueProductNameInv())
//}

func TestNewSettlement(t *testing.T) {
	listSettlements := []TestSettlementValidate{
		{
			id:   22324,
			code: "",
			//user:     user2.User,
			//customer: customer2.Customer,
			//product:    []inventory.Inventory,
			//coupon:     coupon.Coupon,
			totalPrice: 2700000,
			statusTrns: true,
			want:       errors.New("outlet code must be 4 digit"),
		},
	}
	for _, testCase := range listSettlements {
		_, checkErr := entity.NewSettlement(entity.DTOSettlement{
			Id:         testCase.id,
			Code:       testCase.code,
			User:       testCase.user,
			Customer:   testCase.customer,
			Product:    testCase.product,
			Coupon:     testCase.coupon,
			TotalPrice: testCase.totalPrice,
			StatusTrns: testCase.statusTrns,
		})

		assert.Equal(t, testCase.want, checkErr)
	}
}
