package main

import (
	"context"
	"fmt"

	"net/http"

	//dummy "pc-shop-final-project/docs"
	"pc-shop-final-project/domain/entity"
	handler "pc-shop-final-project/internal/repository/controller"
	handlerHTTP "pc-shop-final-project/internal/repository/controller/http_control"

	usecase "pc-shop-final-project/internal/repository/process"

	"github.com/gorilla/mux"
)

func main() {
	var (
		ctx         = context.Background()
		loginUserID int
		statusLogin string
		username    string
		output      int
		token       bool
	)

	r := mux.NewRouter()

	r.HandleFunc("/", handlerHTTP.ParamHandlerWithoutInput).Methods(http.MethodGet)
	r.HandleFunc("/user", handlerHTTP.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/customer", handlerHTTP.GetCustomers).Methods(http.MethodGet)
	r.HandleFunc("/IdUser:{ID}", handlerHTTP.ParamLoginIdUser).Methods(http.MethodGet)
	r.HandleFunc("/IdCustomer:{ID}", handlerHTTP.ParamIdCustomer).Methods(http.MethodGet)
	r.HandleFunc("/NewCustomer", handlerHTTP.ParamNewCustomer).Methods(http.MethodPost)
	r.HandleFunc("/Logout", handlerHTTP.Logout).Methods(http.MethodGet)
	r.HandleFunc("/NewSession", handlerHTTP.NewSession).Methods(http.MethodGet)
	r.HandleFunc("/ScannedItem:{IdItem}", handlerHTTP.ParamAddItemPurchase).Methods(http.MethodGet)
	r.HandleFunc("/BuildSettle", handlerHTTP.ParamSettleProcess).Methods(http.MethodGet)
	r.HandleFunc("/GetCoupon", handlerHTTP.ParamOwnerGetCoupon).Methods(http.MethodGet)
	r.HandleFunc("/GetTransaction", handlerHTTP.ParamOwnerGetTransaction).Methods(http.MethodGet)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)

	if loginUserID == 0 {
		fmt.Printf("Logging in...\nID User : ")
		fmt.Scanln(&loginUserID)
		users := handler.ReadUser(ctx)
		for _, user := range users {
			if user.GetValueIdUsr() == loginUserID {
				fmt.Println("Login SUCCESS...")
				statusLogin = user.GetValueStatusUsr()
				username = user.GetValueNameUsr()
				token = true
			}
		}
		if statusLogin == "" {
			fmt.Println("Login FAILED...")
		}
	}
	if statusLogin == "Employee" {
		fmt.Println("Login as Employee...\nWelcome ", username)
		for token {
			fmt.Println("1. Register Customer\n2. Transaction\n9. Logout")
			fmt.Printf("Output : ")
			fmt.Scanln(&output)
			switch output {
			case 1:
				{
					var nameCust string
					fmt.Printf("Enter name : ")
					fmt.Scanln(&nameCust)
					usecase.NewCustomer(loginUserID, nameCust)
				}
			case 2:
				{
					var (
						item          int
						shoppingTotal int
						priceTotal    int
						price         int
						cust          string = "SON-BGR1-121222093643"
						LUCID         int
						curID         int
						codeUniqs     []string
						sts           string
					)

					purchasedItem := make([]*entity.SettlePurchase, 0)

					dataCpns, _ := handler.ReadCoupon(ctx)
					for _, dataCpn := range dataCpns {
						if dataCpn.GetValueIdCustomerCpn() == cust && dataCpn.GetValueActiveCpn() {
							for _, dataUCpns := range dataCpn.GetValueUniqCpn() {
								LUCID = dataUCpns.GetValueIdCoupon()
								if dataUCpns.GetValueStatusCoupon() {
									codeUniq := dataUCpns.GetValueUniqIdCoupon()
									codeUniqs = append(codeUniqs, codeUniq)
								}
							}
						}
					}

					fmt.Printf("Creating new transaction...\nCustomer ID : ")
					//fmt.Scanln(&cust)
					for item != 9 {
						fmt.Println("Scanning item(s)...")
						fmt.Printf("Item ID (enter 9 to proceed) : ")
						fmt.Scanln(&item)
						purchasedItem, price = usecase.TransactionAddItem(codeUniqs, item, purchasedItem)
						priceTotal += price
						fmt.Println("Total : ", priceTotal)
					}

					shoppingTotal = priceTotal

					for _, dataCpn := range dataCpns {
						if dataCpn.GetValueIdCustomerCpn() == cust && dataCpn.GetValueActiveCpn() {
							if dataCpn.GetValueIdCpn() > LUCID {
								shoppingTotal += dataCpn.GetValueRevenueCpn()
							}
						}
					}

					if len(dataCpns) != 0 {
						curID = handler.GetLastCouponId(ctx) + 1
					} else {
						curID = 1
					}
					cpn := entity.AutoGenerateCoupon(curID, shoppingTotal, cust)

					settle := usecase.NewTransaction(priceTotal, loginUserID, cust, cpn.GetValueIdCpn(), purchasedItem)

					fmt.Println("Status payment? (Y/N)")
					fmt.Scanln(&sts)
					if sts == "Y" || sts == "y" {
						handler.UpdateSettle(ctx, settle.GetValueCodeSett())
						handler.CreateSettleItem(ctx, purchasedItem)
						handler.CreateCoupon(ctx, cpn, cpn.GetValueUniqCpn())
					}
				}
			case 9:
				{
					token = false
				}
			}
		}

	}
	if statusLogin == "Owner" {
		fmt.Println("Login as Owner...\nWelcome ", username)
		for token {
			fmt.Println("1. Show Revenue Database\n2. Show Coupon Database\n3. Clear Expired Data\n9. Logout")
			fmt.Printf("Output : ")
			fmt.Scanln(&output)
			switch output {
			case 1:
				{
					var (
						total int
					)
					trns := handler.ReadSettle(ctx)
					// cust := handler.ReadCustomer(ctx)
					// for _, cst := range cust {
					// 	fmt.Println(cst.GetValueNameCust(), cst.GetValueUniqIdCust())
					// }

					fmt.Println("Data revenue customer this month :")
					revs := entity.NewCustomerRevenue(trns)
					for _, rev := range revs {
						fmt.Println(rev.GetValueIdCustCR(), rev.GetValueRevenueCR(), rev.GetValueDateTransactionCR())
						total += rev.GetValueRevenueCR()
					}
					fmt.Println("Total : ", total)
				}
			case 2:
				{
					trns := handler.ReadSettle(ctx)
					for _, trn := range trns {
						fmt.Println("Code : ", trn.GetValueCodeSett(), "\nTotal : ", trn.GetValueTotalPriceSett())
						fmt.Println("Username : ", trn.GetValueUserSett().GetValueNameUsr())
						fmt.Println("Customer : ", trn.GetValueCustomerSett().GetValueNameCust())
						for _, UC := range trn.GetValueCouponSett().GetValueUniqCpn() {
							fmt.Println("Coupon : ", UC.GetValueUniqIdCoupon(), UC.GetValueStatusCoupon(), trn.GetValueCouponSett().GetValueGenDateCpn())
						}
						fmt.Println("______________________________")
					}
				}
			case 3:
				{
					cpns, _ := handler.ReadCoupon(ctx)
					for _, cpn := range cpns {
						if !cpn.GetValueActiveCpn() {
							handler.DeleteCoupon(ctx, cpn.GetValueIdCpn())
							fmt.Println("Expired data deleted...")
						}
					}
				}
			case 4:
				{
					sets := handler.ReadSettle(ctx)
					for _, set := range sets {
						fmt.Println(set.GetValueCustomerSett().GetValueNameCust(), set.GetValueCouponSett().GetValueIdCustomerCpn(), set.GetValueCouponSett().GetValueRevenueCpn(), set.GetValueTotalPriceSett())
					}
				}
			case 9:
				{
					token = false
				}
			}
		}
	}
}

// dummyInventory := dummy.MakeDummyInventory()
// for _, dumI := range dummyInventory {
// 	handler.CreateInventory(ctx, dumI)
// }
// Inventories := handler.ReadInventory(ctx)
// for _, inv := range Inventories {
// 	fmt.Println(inv.GetValueIdInv())
// 	fmt.Println(inv.GetValueProductNameInv())
// 	fmt.Println(inv.GetValueBrandInv())
// 	fmt.Println(inv.GetvaluePriceInv())
// 	fmt.Println(inv.GetValueCategoryInv())
// 	fmt.Println("____________________")
// }

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
