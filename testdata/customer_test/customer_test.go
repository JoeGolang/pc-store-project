package customer_positive_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"pc-shop-final-project/domain/entity/customer"
	"testing"
)

type TestCustomerValidate struct {
	UniqId   string
	Name     string
	JoinDate string
	want     error
}

// positif FLOW
func TestPositiveNewCustomer(t *testing.T) {
	dataCustomer, err := customer.NewCustomer(customer.DTOCustomer{
		UniqId:   "12111",
		Name:     "Sakura",
		JoinDate: "23-12-2022",
	}, "manggadua")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Sakura", dataCustomer.GetValueNameCust())
}

func TestNewCustomer(t *testing.T) {
	listCustomers := []TestCustomerValidate{
		{
			UniqId:   "11201",
			Name:     "",
			JoinDate: "10-12-2022",
			want:     errors.New("name cannot be empty"),
		},
		{
			UniqId:   "11201",
			Name:     "SA",
			JoinDate: "10-12-2022",
			want:     errors.New("name cannot be less than 3 character"),
		},
	}
	for _, testCase := range listCustomers {
		_, checkErr := customer.NewCustomer(customer.DTOCustomer{
			UniqId:   testCase.UniqId,
			Name:     testCase.Name,
			JoinDate: testCase.JoinDate,
		}, "manggadua")

		assert.Equal(t, testCase.want, checkErr)
	}
}
