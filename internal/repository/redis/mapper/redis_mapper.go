package mapper

import (
	"encoding/json"
	"pc-shop-final-project/domain/entity"
	"strconv"
)

type User struct {
	IdUser     int    `json:"UserID"`
	NameUser   string `json:"UserName"`
	OutletCode string `json:"Outlet"`
	StatusUser string `json:"Status"`
}

type Customer struct {
	IdCustomer   string `json:"CustomerID"`
	NameCustomer string `json:"CustomerName"`
	JoinDate     string `json:"JoinDate"`
}

type SettleItem struct {
	IdItem int `json:"ItemID"`
	Qty    int `json:"Qty"`
}

// item
func MapSetSettleItemToStringSlice(data []*entity.SettlePurchase) []string {
	var (
		listItem []string
		i        int
	)

	for _, item := range data {
		SettleItem := &SettleItem{
			IdItem: item.GetValueIdProduct(),
			Qty:    item.GetValueQtyProduct(),
		}

		Json, _ := json.Marshal(SettleItem)

		listItem = append(listItem, strconv.Itoa(i))
		listItem = append(listItem, string(Json))
		i++
	}

	return listItem
}

func MapToSliceSettleItem(data []SettleItem) []*entity.SettlePurchase {
	SettleList := make([]*entity.SettlePurchase, 0)
	for _, set := range data {
		redisSettle := entity.NewSettlePurchase(set.IdItem, set.Qty)
		SettleList = append(SettleList, redisSettle)
	}
	return SettleList
}

// user
func MapSetUserToString(data *entity.User) string {
	User := &User{
		IdUser:     data.GetValueIdUsr(),
		NameUser:   data.GetValueNameUsr(),
		OutletCode: data.GetValueOutletCodeUsr(),
		StatusUser: data.GetValueStatusUsr(),
	}

	Json, _ := json.Marshal(User)

	return string(Json)
}

func MapJsonStringToUser(data string) (*entity.User, error) {
	Users := new(User)
	err := json.Unmarshal([]byte(data), Users)
	if err != nil {
		return nil, err
	}

	redisUser, errFetch := entity.NewUser(&entity.DTOUser{
		Id:         Users.IdUser,
		Name:       Users.NameUser,
		OutletCode: Users.OutletCode,
		Status:     Users.StatusUser,
	})
	if errFetch != nil {
		return nil, errFetch
	}

	return redisUser, nil
}

// customer
func MapSetCustToString(data *entity.Customer) string {
	Customer := &Customer{
		IdCustomer:   data.GetValueUniqIdCust(),
		NameCustomer: data.GetValueNameCust(),
		JoinDate:     data.GetValueJoinDateCust(),
	}

	Json, _ := json.Marshal(Customer)

	return string(Json)
}

func MapJsonStringToCust(data string) (*entity.Customer, error) {
	Custs := new(Customer)
	err := json.Unmarshal([]byte(data), Custs)
	if err != nil {
		return nil, err
	}

	redisCust, errFetch := entity.NewCustomer(&entity.DTOCustomer{
		UniqId:   Custs.IdCustomer,
		Name:     Custs.NameCustomer,
		JoinDate: Custs.JoinDate,
	}, "")
	if errFetch != nil {
		return nil, errFetch
	}

	return redisCust, nil
}
