package main

import (
	"context"
	"fmt"
	dummy "pc-shop-final-project/docs"
	handler "pc-shop-final-project/internal/repository/controller"
)

func main() {
	var (
		ctx = context.Background()
	)

	dummyInventory := dummy.MakeDummyInventory()
	for _, dumI := range dummyInventory {
		handler.CreateInventory(ctx, dumI)
	}
	Inventories := handler.ReadInventory(ctx)
	for _, inv := range Inventories {
		fmt.Println(inv.GetValueIdInv())
		fmt.Println(inv.GetValueProductNameInv())
		fmt.Println(inv.GetValueBrandInv())
		fmt.Println(inv.GetvaluePriceInv())
		fmt.Println(inv.GetValueCategoryInv())
		fmt.Println("____________________")
	}

	// dummyUser := dummy.MakeDummyUser()
	// for _, dumU := range dummyUser {
	// 	handler.CreateUser(ctx, dumU)
	// }
	// Users := handler.ReadUser(ctx)
	// for _, user := range Users {
	// 	fmt.Println(user.GetValueIdUsr())
	// 	fmt.Println(user.GetValueNameUsr())
	// 	fmt.Println(user.GetValueOutletCodeUsr())
	// 	fmt.Println(user.GetValueStatusUsr())
	// 	fmt.Println("____________________")
	// }

	// dummyCustomer := dummy.MakeDummyCustomer()
	// for _, dumC := range dummyCustomer {
	// 	handler.CreateCustomer(ctx, dumC)
	// }
	// Customers := handler.ReadCustomer(ctx)
	// for _, customer := range Customers {
	// 	fmt.Println(customer.GetValueUniqIdCust())
	// 	fmt.Println(customer.GetValueNameCust())
	// 	fmt.Println(customer.GetValueJoinDateCust())
	// 	fmt.Println("____________________")
	// }
}
