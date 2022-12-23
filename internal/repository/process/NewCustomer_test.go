package usecase_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	usecase "pc-shop-final-project/internal/repository/process"
	"testing"
)

// positive test
func TestNewCustomer(t *testing.T) {

	userId := 1002
	CustName := "Akira"
	data, err := usecase.NewCustomer(userId, CustName)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Equal(t, "Akira", data.GetValueNameCust())
}

// negative test
func TestNewCustomerNeg(t *testing.T) {

	userId := 1002
	CustName := "Ak"
	data, err := usecase.NewCustomer(userId, CustName)
	if err != nil {
		fmt.Println(err.Error())
	}
	if data != nil {
		fmt.Println(data)
	}

	assert.Equal(t, errors.New("name cannot be less than 3 character"), err)
}
