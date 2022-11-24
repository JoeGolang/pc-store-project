package storeUser

import (
	"errors"
	"time"
)

type StoreUser struct {
	user_id    int
	first_name string
	last_name  string
	is_admin   bool
	createDate time.Time
}

func NewStoreUser(user_id_set int, first_name_set string, last_name_set string, is_admin_set bool) (*StoreUser, error) {
	if user_id_set == 0 {
		return nil, errors.New("UserId must be filled")
	}
	if first_name_set == "" {
		return nil, errors.New("First Name must be filled")
	}
	if last_name_set == "" {
		return nil, errors.New("UserId must be filled")
	}
	return &StoreUser{
		user_id:    user_id_set,
		first_name: first_name_set,
		last_name:  last_name_set,
		is_admin:   is_admin_set,
		createDate: time.Now(),
	}, nil
}

func (st *StoreUser) GetUserId() int {
	return st.user_id
}
func (st *StoreUser) GetFirstName() string {
	return st.first_name
}
func (st *StoreUser) GetLastName() string {
	return st.last_name
}
