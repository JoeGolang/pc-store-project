package user

import (
	"errors"
	"pc-shop-final-project/domain/value_object"
	"pc-shop-final-project/internal/delivery/http/http_request"
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

func NewUser(dto DTOUser) (*User, error) {
	if dto.Name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len([]rune(dto.Name)) < 3 {
		return nil, errors.New("user code must be more than 3 characters")
	}

	if dto.OutletCode == "" {
		return nil, errors.New("code outlet cannot be empty")
	}

	if len([]rune(dto.OutletCode)) != 4 {
		return nil, errors.New("outlet code code must 4 characters")
	}

	if dto.Status == "" {
		return nil, errors.New("status cannot be empty")
	}

	status, errUser := value_object.NewStatusUser(value_object.StatusUserInt(dto.Status))
	if errUser != nil {
		return nil, errUser
	}

	return &User{
		id:         dto.Id,
		name:       dto.Name,
		outletCode: dto.OutletCode,
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

func (us *User) SetUpdateData(req http_request.RequestUser) {
	if req.ID_USER != 0 {
		us.id() = req.ID_USER
	}
	if req.NAME != "" {
		us.name() = req.NAME
	}
	if req.OUTLET_CODE != "" {
		us.outletCode = req.OUTLET_CODE
	}
	if req.STATUS != "" {
		us.status = req.STATUS
	}

}
