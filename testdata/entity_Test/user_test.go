package entity_Test_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"pc-shop-final-project/domain/entity"
	"testing"
)

type TestUsersValidate struct {
	id         int
	name       string
	outletCode string
	status     string
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

func TestNewUser(t *testing.T) {
	listUsers := []TestUsersValidate{
		{
			id:         7001,
			name:       "Joe",
			outletCode: "BGR",
			status:     "Owner",
			want:       errors.New("user code must 4 characters"),
		},
	}
	for _, testCase := range listUsers {
		_, checkErr := entity.NewUser(&entity.DTOUser{
			Id:         testCase.id,
			Name:       testCase.name,
			OutletCode: testCase.outletCode,
			Status:     testCase.status,
		})

		assert.Equal(t, testCase.want, checkErr)
	}
}
