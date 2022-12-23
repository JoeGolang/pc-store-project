package usecase

import (
	"context"
	"errors"
	"fmt"
	"pc-shop-final-project/domain/entity"
	handler "pc-shop-final-project/internal/repository/controller"
)

var ctx = context.Background()

func NewCustomer(userID int, nameCustomer string) (*entity.Customer, error) {
	var currentUser = &entity.User{}

	DTONewCust := &entity.DTOCustomer{
		Name: nameCustomer,
	}

	users := handler.ReadUser(ctx)
	for _, user := range users {
		if userID == user.GetValueIdUsr() {
			currentUser = user
		}
	}

	if currentUser == nil {
		return nil, errors.New("cannot find user")
	}

	NewCust, err := entity.NewCustomer(DTONewCust, currentUser.GetValueOutletCodeUsr())
	if err != nil {
		return nil, err
	}

	handler.CreateCustomer(ctx, NewCust)

	return NewCust, nil
}

func NewUser(userID int, id int, name string, outlet string, status string) (*entity.User, error) {
	var currentUser = &entity.User{}

	DTONewUser := &entity.DTOUser{
		Id:         id,
		Name:       name,
		OutletCode: outlet,
		Status:     status,
	}

	users := handler.ReadUser(ctx)
	for _, user := range users {
		if userID == user.GetValueIdUsr() {
			currentUser = user
		}
	}

	if currentUser == nil {
		return nil, errors.New("cannot find user")
	}

	NewUser, err := entity.NewUser(DTONewUser)
	if err != nil {
		return nil, err
	}

	handler.CreateUser(ctx, NewUser)

	return NewUser, nil
}

func TransactionAddItem(UCs []string, itemID int, itemList []*entity.SettlePurchase) ([]*entity.SettlePurchase, int) {
	var (
		price int
		cut   int
		sts   bool
	)
	Inventories := handler.ReadInventory(ctx)

	for _, Item := range Inventories {
		if Item.GetValueIdInv() == itemID {
			for _, Purchased := range itemList {
				if Purchased.GetValueIdProduct() == itemID {
					Purchased.SetValueQtyProduct()
					sts = true
				}
			}
			if !sts {
				ItemAdd := entity.NewSettlePurchase(itemID, 1)
				itemList = append(itemList, ItemAdd)
			}

			for _, codeUniq := range UCs {
				strCodeU := []rune(codeUniq)
				if strCodeU[0] == 'B' {
					if Item.GetValueCategoryInv() == "Second Game" {
						cut = Item.GetvaluePriceInv() * 5 / 100
						handler.UpdateCoupon(ctx, codeUniq)
						fmt.Println("Coupon ", codeUniq, " get discount 5%")
					}
				}
				if strCodeU[0] == 'P' {
					if Item.GetValueCategoryInv() == "Service Console" {
						cut = Item.GetvaluePriceInv() * 15 / 100
						handler.UpdateCoupon(ctx, codeUniq)
						fmt.Println("Coupon ", codeUniq, " get discount 15%")
					}
				}
				if strCodeU[0] == 'U' {
					if Item.GetValueCategoryInv() == "Accessories Console" || Item.GetValueCategoryInv() == "New Game" {
						cut = Item.GetvaluePriceInv() * 30 / 100
						handler.UpdateCoupon(ctx, codeUniq)
						fmt.Println("Coupon ", codeUniq, " get discount 30%")
					}
				}
			}

			price = Item.GetvaluePriceInv() - cut
		}
	}
	return itemList, price
}

func TransactionNewItem(itemID int, itemList []*entity.SettlePurchase) ([]*entity.SettlePurchase, int) {
	var price int
	Inventories := handler.ReadInventory(ctx)

	for _, Item := range Inventories {
		if Item.GetValueIdInv() == itemID {
			fmt.Println("Item ", Item.GetValueProductNameInv(), "-", Item.GetValueCategoryInv(), " Added...")

			ItemAdd := entity.NewSettlePurchase(itemID, 1)
			itemList = append(itemList, ItemAdd)
			price = Item.GetvaluePriceInv()
		}
	}
	return itemList, price
}

func TransactionNewItems(productDetail string, itemID int, itemList []*entity.SettlePurchase) ([]*entity.SettlePurchase, int, string) {
	var price int
	Inventories := handler.ReadInventory(ctx)

	for _, Item := range Inventories {
		if Item.GetValueIdInv() == itemID {
			fmt.Println("Item ", Item.GetValueProductNameInv(), "-", Item.GetValueCategoryInv(), " Added...")

			ItemAdd := entity.NewSettlePurchase(itemID, 1)
			itemList = append(itemList, ItemAdd)
			price = Item.GetvaluePriceInv()
			productDetail = Item.GetValueProductNameInv()
		}
	}
	return itemList, price, productDetail
}

func GetInvDetail(invID int) (invDetail string, invCat string) {
	Inventories := handler.ReadInventory(ctx)

	for _, Inv := range Inventories {
		if Inv.GetValueIdInv() == invID {
			invDetail = Inv.GetValueProductNameInv()
			invCat = Inv.GetValueCategoryInv()
		}
	}
	return invDetail, invCat
}

func NewTransaction(totP int, userID int, customerID string, couponID int, SI []*entity.SettlePurchase) *entity.Settlement {
	var (
		a entity.User
		b entity.Customer
		c entity.Coupon
	)

	usrs := handler.ReadUser(ctx)
	csts := handler.ReadCustomer(ctx)
	cpns, _ := handler.ReadCoupon(ctx)

	for _, usr := range usrs {
		if usr.GetValueIdUsr() == userID {
			a = *usr
		}
	}
	for _, cst := range csts {
		if cst.GetValueUniqIdCust() == customerID {
			b = *cst
		}
	}
	for _, cpn := range cpns {
		if cpn.GetValueIdCpn() == couponID {
			c = *cpn
		}
	}

	settle := entity.DTOSettlement{
		User:       a,
		Customer:   b,
		Product:    SI,
		Coupon:     c,
		TotalPrice: totP,
	}
	Transaction, err := entity.NewSettlement(settle)
	if err != nil {
		fmt.Println(err)
	}

	handler.CreateSettle(ctx, userID, customerID, couponID, Transaction)

	return Transaction
}
