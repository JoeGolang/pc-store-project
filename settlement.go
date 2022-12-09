package entity

import (
	"errors"
	"pc-shop-final-project/domain/entity/coupon"
	customer2 "pc-shop-final-project/domain/entity/customer"
	user2 "pc-shop-final-project/domain/entity/user"
	"strconv"
	"time"
)

type Settlement struct {
	id         int
	code       string
	user       user2.User
	customer   customer2.Customer
	product    []Inventory
	coupon     coupon.Coupon
	totalPrice int
	statusTrns bool
}

type DTOSettlement struct {
	Id         int
	Code       string
	User       user2.User
	Customer   customer2.Customer
	Product    []Inventory
	Coupon     coupon.Coupon
	TotalPrice int
	StatusTrns bool
}

type SettlePurchase struct {
	idSettle  int
	idProduct int
	qty       int
}

func NewSettlement(set DTOSettlement) (*Settlement, error) {
	inven := Inventory{}
	inventories := make([]Inventory, 0)
	for _, inv := range set.Product {
		inven.id = inv.id
		inven.productName = inv.productName
		inven.brand = inv.brand
		inven.category = inv.category
		inventories = append(inventories, inven)
	}

	if set.Code == "" {
		if len([]rune(set.User.GetValueOutletCodeUsr())) != 4 {
			err := errors.New("outlet code must be 4 digit")
			return nil, err
		}
		time := time.Now().Format("020106150405")
		userSettleCode := set.User.GetValueIdUsr()
		stringUserSettleCode := strconv.Itoa(userSettleCode)
		if len([]rune(stringUserSettleCode)) != 4 {
			err := errors.New("user code must be 4 digit")
			return nil, err
		}
		set.Code = set.User.GetValueOutletCodeUsr() + stringUserSettleCode + time
	}

	return &Settlement{
		id:         set.Id,
		code:       set.Code,
		user:       set.User,
		customer:   set.Customer,
		product:    inventories,
		coupon:     set.Coupon,
		totalPrice: set.TotalPrice,
		statusTrns: set.StatusTrns,
	}, nil
}

func (sett *Settlement) GetValueIdSett() int {
	return sett.id
}

func (sett *Settlement) GetValueUserSett() *user2.User {
	return &sett.user
}

func (sett *Settlement) GetValueCustomerSett() *customer2.Customer {
	return &sett.customer
}

func (sett *Settlement) GetValueCouponSett() *coupon.Coupon {
	return &sett.coupon
}

func (sett *Settlement) GetValueProductSett() *[]Inventory {
	return &sett.product
}

func (sett *Settlement) GetValueCodeSett() string {
	return sett.code
}

func (sett *Settlement) GetValueTotalPriceSett() int {
	return sett.totalPrice
}

func (sett *Settlement) SetStatusTrns() *Settlement {
	sett.statusTrns = true
	return &Settlement{}
}

func (sett *Settlement) GetValueStatusTrns() bool {
	return sett.statusTrns
}

func NewSettlePurchase(idS int, idP int, qty int) *SettlePurchase {
	return &SettlePurchase{
		idSettle:  idS,
		idProduct: idP,
		qty:       qty,
	}
}

func (set *SettlePurchase) GetValueIdSettlement() int {
	return set.idSettle
}

func (set *SettlePurchase) GetValueIdProduct() int {
	return set.idProduct
}

func (set *SettlePurchase) GetValueQtyProduct() int {
	return set.qty
}
