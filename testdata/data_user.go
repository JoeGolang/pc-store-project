package customer

import "pc-shop-final-project/domain/entity"

func MakeDummyUser() []*entity.User {
	DataOwner1 := &entity.DTOUser{
		Id:         7001,
		Name:       "Joe",
		OutletCode: "BGR1",
		Status:     "Owner",
	}
	UserOwner1, err := entity.NewUser(DataOwner1)
	if err != nil {
		panic("wrong data user")
	}

	DataEmployee1 := &entity.DTOUser{
		Id:         1001,
		Name:       "Yuna",
		OutletCode: "BGR1",
		Status:     "Employee",
	}
	UserEmployee1, err := entity.NewUser(DataEmployee1)
	if err != nil {
		panic("wrong data user")
	}

	DataEmployee2 := &entity.DTOUser{
		Id:         1002,
		Name:       "Yeji",
		OutletCode: "BGR1",
		Status:     "Employee",
	}
	UserEmployee2, err := entity.NewUser(DataEmployee2)
	if err != nil {
		panic("wrong data user")
	}

	DataEmployee3 := &entity.DTOUser{
		Id:         1003,
		Name:       "Ryujin",
		OutletCode: "BGR1",
		Status:     "Employee",
	}
	UserEmployee3, err := entity.NewUser(DataEmployee3)
	if err != nil {
		panic("wrong data user")
	}

	DataOwner2 := &entity.DTOUser{
		Id:         7002,
		Name:       "Juria",
		OutletCode: "BGR2",
		Status:     "Owner",
	}
	UserOwner2, err := entity.NewUser(DataOwner2)
	if err != nil {
		panic("wrong data user")
	}

	DataEmployee4 := &entity.DTOUser{
		Id:         1004,
		Name:       "Lia",
		OutletCode: "BGR2",
		Status:     "Employee",
	}
	UserEmployee4, err := entity.NewUser(DataEmployee4)
	if err != nil {
		panic("wrong data user")
	}

	DataEmployee5 := &entity.DTOUser{
		Id:         1005,
		Name:       "Chae",
		OutletCode: "BGR2",
		Status:     "Employee",
	}
	UserEmployee5, err := entity.NewUser(DataEmployee5)
	if err != nil {
		panic("wrong data user")
	}

	listUser := make([]*entity.User, 0)
	listUser = append(listUser, UserOwner1, UserEmployee1, UserEmployee2, UserEmployee3, UserOwner2, UserEmployee4, UserEmployee5)
	return listUser
}
