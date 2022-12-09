package user

import (
	"errors"
	"pc-shop-final-project/domain/value_object"
)

type User struct {
	id         int
	name       string
	outletCode string
	status     *value_object.StatusUser
}

type DTOUser struct {
	Id         int
	Name       string
	OutletCode string
	Status     string
}

func NewUser(usr *DTOUser) (*User, error) {
	if usr.Name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len([]rune(usr.Name)) < 3 {
		return nil, errors.New("user code must be more than 3 characters")
	}

	if usr.OutletCode == "" {
		return nil, errors.New("code outlet cannot be empty")
	}

	if len([]rune(usr.OutletCode)) != 4 {
		return nil, errors.New("user code must 4 characters")
	}

	if usr.Status == "" {
		return nil, errors.New("status cannot be empty")
	}

	status, errUser := value_object.NewStatusUser(value_object.StatusUserInt(usr.Status))
	if errUser != nil {
		return nil, errUser
	}

	return &User{
		id:         usr.Id,
		name:       usr.Name,
		outletCode: usr.OutletCode,
		status:     status,
	}, nil
}

func (usr *User) GetValueIdUsr() int {
	return usr.id
}

func (usr *User) GetValueNameUsr() string {
	return usr.name
}

func (usr *User) GetValueOutletCodeUsr() string {
	return usr.outletCode
}

func (usr *User) GetValueStatusUsr() string {
	return usr.status.StatusUserToString()
}
