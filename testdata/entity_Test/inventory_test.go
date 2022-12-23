package entity_Test_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"pc-shop-final-project/domain/entity"
	"testing"
)

type TestInventoryValidate struct {
	id          int
	productName string
	brand       string
	price       int
	category    string
	want        error
}

// positif FLOW
func TestPositiveNewInventory(t *testing.T) {
	dataInventory, err := entity.NewInventory(&entity.DTOInventory{
		Id:          1234,
		ProductName: "Switch",
		Brand:       "Nintendo",
		Price:       2500000,
		Category:    "New Console",
	})

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Switch", dataInventory.GetValueProductNameInv())
}

func TestNewInventory(t *testing.T) {
	listInventories := []TestInventoryValidate{
		{
			id:          1234,
			productName: "",
			brand:       "Nintendo",
			price:       2500000,
			category:    "New Console",
			want:        errors.New("product name cannot be empty"),
		},
		{
			id:          1234,
			productName: "Switch",
			brand:       "",
			price:       2500000,
			category:    "New Console",
			want:        errors.New("brand cannot be empty"),
		},
	}
	for _, testCase := range listInventories {
		_, checkErr := entity.NewInventory(&entity.DTOInventory{
			Id:          testCase.id,
			ProductName: testCase.productName,
			Brand:       testCase.brand,
			Price:       testCase.price,
			Category:    testCase.category,
		})

		assert.Equal(t, testCase.want, checkErr)
	}
}
