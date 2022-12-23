package handlerHTTP

import (
	"fmt"
	"net/http"
	"pc-shop-final-project/domain/entity"
	"pc-shop-final-project/internal/delivery/http/http_response"
	handler "pc-shop-final-project/internal/repository/controller"
	handler_redis "pc-shop-final-project/internal/repository/controller/redis_control"
	usecase "pc-shop-final-project/internal/repository/process"
)

func ParamSettleProcess(w http.ResponseWriter, r *http.Request) {
	var (
		shoppingTotal int
		gross         int
		discount      int
		priceTotal    int
		price         int
		LUCID         int
		curID         int
		codeUniqs     []string
	)
	userData, custData := handler_redis.GetUserCustomerRedis(ctx)

	if userData.GetValueIdUsr() == 0 {
		fmt.Println("LOGIN TO PROCEED...")
		fmt.Fprintf(w, "LOGIN TO PROCEED...")
		w.WriteHeader(http.StatusBadRequest)
	}
	if userData.GetValueIdUsr() != 0 {
		if userData.GetValueStatusUsr() == "Employee" {
			if custData.GetValueUniqIdCust() != "" {

				dataCpns, _ := handler.ReadCoupon(ctx)
				for _, dataCpn := range dataCpns {
					if dataCpn.GetValueIdCustomerCpn() == custData.GetValueUniqIdCust() && dataCpn.GetValueActiveCpn() {
						for _, dataUCpns := range dataCpn.GetValueUniqCpn() {
							LUCID = dataUCpns.GetValueIdCoupon()
							if dataUCpns.GetValueStatusCoupon() {
								codeUniq := dataUCpns.GetValueUniqIdCoupon()
								codeUniqs = append(codeUniqs, codeUniq)
							}
						}
					}
				}

				purchasedItem := make([]*entity.SettlePurchase, 0)
				_, _, i := handler_redis.GetAllRedis(ctx)

				for _, item := range i {
					_, priceNoDiscount := usecase.TransactionNewItem(item.GetValueIdProduct(), purchasedItem)
					purchasedItem, price = usecase.TransactionAddItem(codeUniqs, item.GetValueIdProduct(), purchasedItem)
					gross += priceNoDiscount
					priceTotal += price
					discount += (priceNoDiscount - price)
				}

				fmt.Println("Price before discount : ", gross)
				fmt.Println("Discount : ", discount)
				fmt.Println("Total : ", priceTotal)

				fmt.Fprintf(w, "Transaction Receipt For : %s \n", custData.GetValueNameCust())
				fmt.Fprintf(w, "Cashier: %s \n", userData.GetValueNameUsr())

				fmt.Fprintln(w, "item detail : ")
				for n := 0; n < len(purchasedItem); n++ {
					invID := purchasedItem[n].GetValueIdProduct()
					invDetail, invCat := usecase.GetInvDetail(invID)
					fmt.Fprintln(w, invDetail, "(", invCat, ")", "____________ x ", purchasedItem[n].GetValueQtyProduct())
				}
				fmt.Fprintln(w, "Price before discount : ", gross)
				fmt.Fprintln(w, "Discount : ", discount)
				fmt.Fprintln(w, "Total : ", priceTotal)

				shoppingTotal = priceTotal

				for _, dataCpn := range dataCpns {
					if dataCpn.GetValueIdCustomerCpn() == custData.GetValueUniqIdCust() && dataCpn.GetValueActiveCpn() {
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
				cpn := entity.AutoGenerateCoupon(curID, shoppingTotal, custData.GetValueUniqIdCust())

				fmt.Fprintln(w, "Generated Coupon : ")

				for _, UC := range cpn.GetValueUniqCpn() {
					fmt.Println("Generated Coupon : ", UC.GetValueUniqIdCoupon())
					fmt.Fprintln(w, UC.GetValueUniqIdCoupon())

				}

				settle := usecase.NewTransaction(priceTotal, userData.GetValueIdUsr(), custData.GetValueUniqIdCust(), cpn.GetValueIdCpn(), purchasedItem)

				handler.UpdateSettle(ctx, settle.GetValueCodeSett())
				handler.CreateSettleItem(ctx, purchasedItem)
				handler.CreateCoupon(ctx, cpn, cpn.GetValueUniqCpn())
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Transaction Success..")

				response, errMap := http_response.MapResponseSettlement(settle, 200, "Success")
				if errMap != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Error mapping data"))
				}

				w.WriteHeader(200)
				w.Write(response)

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
