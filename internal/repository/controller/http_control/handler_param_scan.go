package handlerHTTP

import (
	"fmt"
	"net/http"
	"pc-shop-final-project/domain/entity"
	handler "pc-shop-final-project/internal/repository/controller"
	handler_redis "pc-shop-final-project/internal/repository/controller/redis_control"
	usecase "pc-shop-final-project/internal/repository/process"
	"strconv"

	"github.com/gorilla/mux"
)

func ParamAddItemPurchase(w http.ResponseWriter, r *http.Request) {
	var (
		item          int
		priceTotal    int
		price         int
		user          *entity.User
		cust          *entity.Customer
		purchasedItem []*entity.SettlePurchase
	)
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Input ID User : %v\n", vars["IdItem"])

	id := vars["IdItem"]
	item, _ = strconv.Atoi(id)

	userData, custData := handler_redis.GetUserCustomerRedis(ctx)

	if userData.GetValueIdUsr() == 0 {
		fmt.Println("LOGIN TO PROCEED...")
		fmt.Fprintf(w, "LOGIN TO PROCEED...")
		w.WriteHeader(http.StatusBadRequest)
	}
	if custData.GetValueUniqIdCust() == "" {
		fmt.Println("Customer not found...")
		fmt.Fprintf(w, "PLEASE ENTER CUSTOMER DATA...")
		w.WriteHeader(http.StatusBadRequest)
	}
	if userData.GetValueIdUsr() != 0 {
		if userData.GetValueStatusUsr() == "Employee" {
			if custData.GetValueUniqIdCust() != "" {

				fmt.Println("Scanning item(s)...")
				fmt.Println("Item ID : ", item)

				user, cust, purchasedItem = handler_redis.GetAllRedis(ctx)
				purchasedItem, price = usecase.TransactionNewItem(item, purchasedItem)
				handler_redis.SetSetItRedis(ctx, purchasedItem)
				items := handler.ReadInventory(ctx)
				for _, p := range purchasedItem {
					for _, i := range items {
						if p.GetValueIdProduct() == i.GetValueIdInv() {
							priceTotal += i.GetvaluePriceInv()
							fmt.Fprintln(w, i.GetValueProductNameInv(), "(", i.GetValueCategoryInv(), ")", i.GetvaluePriceInv())

						}
					}
				}

				fmt.Println("Transaction : ", user.GetValueNameUsr(), " - ", cust.GetValueNameCust())
				fmt.Println("Price : ", price)
				fmt.Println("Total : ", priceTotal)
				fmt.Println("______________________________")

				fmt.Fprintln(w, "Total : ", priceTotal)
				fmt.Fprintln(w, "______________________________")
				fmt.Fprintln(w, "Transaction : ", user.GetValueNameUsr(), " - ", cust.GetValueNameCust())

				//response show data item purchased, show data item new added, price, total
				w.WriteHeader(http.StatusProcessing)
			}
		}
		if userData.GetValueStatusUsr() == "Owner" {
			if custData.GetValueUniqIdCust() != "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "no access...")
			}
		}
	}
}
