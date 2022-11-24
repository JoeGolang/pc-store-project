package storeUser_test

import (
	storeUser2 "gamestore/domain/entity/storeUser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStoreUser(t *testing.T) {
	storeUser, err := storeUser2.NewStoreUser(1001, "Rudi", "Butarbutar", true)
	assert.Nil(t, err)
	assert.Equal(t, 1001, storeUser.GetUserId())
	assert.Equal(t, "Rudi", storeUser.GetFirstName())
	assert.Equal(t, "Butarbutar", storeUser.GetLastName())
}
