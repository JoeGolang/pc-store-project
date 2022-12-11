//package customer
//
//import "pc-shop-final-project/domain/entity/user"
//
//func MakeDummyUser() []*user.User {
//	DataOwner1 := &user.DTOUser{
//		Id:         7001,
//		Name:       "Joe",
//		OutletCode: "BGR1",
//		Status:     "Owner",
//	}
//	UserOwner1, err := user.NewUser(DataOwner1)
//	if err != nil {
//		panic("wrong data user")
//	}
//
//	DataEmployee1 := &user.DTOUser{
//		Id:         1001,
//		Name:       "Yuna",
//		OutletCode: "BGR1",
//		Status:     "Employee",
//	}
//	UserEmployee1, err := user.NewUser(DataEmployee1)
//	if err != nil {
//		panic("wrong data user")
//	}
//
//	DataEmployee2 := &user.DTOUser{
//		Id:         1002,
//		Name:       "Yeji",
//		OutletCode: "BGR1",
//		Status:     "Employee",
//	}
//	UserEmployee2, err := user.NewUser(DataEmployee2)
//	if err != nil {
//		panic("wrong data user")
//	}
//
//	DataEmployee3 := &user.DTOUser{
//		Id:         1003,
//		Name:       "Ryujin",
//		OutletCode: "BGR1",
//		Status:     "Employee",
//	}
//	UserEmployee3, err := user.NewUser(DataEmployee3)
//	if err != nil {
//		panic("wrong data user")
//	}
//
//	DataOwner2 := &user.DTOUser{
//		Id:         7002,
//		Name:       "Juria",
//		OutletCode: "BGR2",
//		Status:     "Owner",
//	}
//	UserOwner2, err := user.NewUser(DataOwner2)
//	if err != nil {
//		panic("wrong data user")
//	}
//
//	DataEmployee4 := &user.DTOUser{
//		Id:         1004,
//		Name:       "Lia",
//		OutletCode: "BGR2",
//		Status:     "Employee",
//	}
//	UserEmployee4, err := user.NewUser(DataEmployee4)
//	if err != nil {
//		panic("wrong data user")
//	}
//
//	DataEmployee5 := &user.DTOUser{
//		Id:         1005,
//		Name:       "Chae",
//		OutletCode: "BGR2",
//		Status:     "Employee",
//	}
//	UserEmployee5, err := user.NewUser(DataEmployee5)
//	if err != nil {
//		panic("wrong data user")
//	}
//
//	listUser := make([]*user.User, 0)
//	listUser = append(listUser, UserOwner1, UserEmployee1, UserEmployee2, UserEmployee3, UserOwner2, UserEmployee4, UserEmployee5)
//	return listUser
//}
