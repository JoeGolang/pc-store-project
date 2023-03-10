package entity

import (
	"errors"
	"strconv"
	"time"
)

type Settlement struct {
	id         int
	code       string
	user       User
	customer   Customer
	product    []*SettlePurchase
	coupon     Coupon
	totalPrice int
	statusTrns bool
}

type DTOSettlement struct {
	Id         int
	Code       string
	User       User
	Customer   Customer
	Product    []*SettlePurchase
	Coupon     Coupon
	TotalPrice int
	StatusTrns bool
}

type SettlePurchase struct {
	codeTransaction string
	idProduct       int
	qty             int
}

func NewSettlement(set DTOSettlement) (*Settlement, error) {
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

	for _, product := range set.Product {
		product.codeTransaction = set.Code
	}

	return &Settlement{
		id:         set.Id,
		code:       set.Code,
		user:       set.User,
		customer:   set.Customer,
		product:    set.Product,
		coupon:     set.Coupon,
		totalPrice: set.TotalPrice,
		statusTrns: set.StatusTrns,
	}, nil
}

func FetchSettlement(usr int, cst string, cpn int, set DTOSettlement) (int, string, int, *Settlement) {
	return usr, cst, cpn, &Settlement{
		id:         set.Id,
		code:       set.Code,
		totalPrice: set.TotalPrice,
		statusTrns: set.StatusTrns,
	}
}

func (sett *Settlement) GetValueIdSett() int {
	return sett.id
}

func (sett *Settlement) GetValueUserSett() *User {
	return &sett.user
}

func (sett *Settlement) GetValueCustomerSett() *Customer {
	return &sett.customer
}

func (sett *Settlement) GetValueCouponSett() *Coupon {
	return &sett.coupon
}

func (sett *Settlement) GetValueProductSett() []*SettlePurchase {
	return sett.product
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

func NewSettlePurchase(idP int, qty int) *SettlePurchase {
	return &SettlePurchase{
		idProduct: idP,
		qty:       qty,
	}
}

func FetchSettlePurchase(SC string, idP int, qty int) *SettlePurchase {
	return &SettlePurchase{
		codeTransaction: SC,
		idProduct:       idP,
		qty:             qty,
	}
}

func (set *SettlePurchase) GetValueCodeSettlement() string {
	return set.codeTransaction
}

func (set *SettlePurchase) GetValueIdProduct() int {
	return set.idProduct
}

func (set *SettlePurchase) GetValueQtyProduct() int {
	return set.qty
}

func (set *SettlePurchase) SetValueQtyProduct() *SettlePurchase {
	set.qty++
	return set
}
