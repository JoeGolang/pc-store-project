package handlerHTTP

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pc-shop-final-project/domain/entity"
	"pc-shop-final-project/internal/delivery/http/http_request"
	handler_redis "pc-shop-final-project/internal/repository/controller/redis_control"
	usecase "pc-shop-final-project/internal/repository/process"
)

func ParamNewCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		req          http_request.RequestCustomer
		decoder      = json.NewDecoder(r.Body)
		customerData *entity.Customer
		err          error
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	userData := handler_redis.GetUserRedis(ctx)
	if userData.GetValueIdUsr() == 0 {
		fmt.Fprintf(w, "LOGIN TO PROCEED...")
		fmt.Println("LOGIN TO PROCEED...")
		w.WriteHeader(http.StatusBadRequest)
	}
	if userData.GetValueIdUsr() != 0 {
		if userData.GetValueStatusUsr() == "Employee" {
			fmt.Println("Login as ", userData.GetValueStatusUsr())
			fmt.Println("New Customer Name : ", req.Name)
			customerData, err = usecase.NewCustomer(userData.GetValueIdUsr(), req.Name)
			if err != nil {
				fmt.Println(err)
			}
			//response data new customer
			handler_redis.SetCustomerRedis(ctx, customerData)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "New Customer ID Created...")
			fmt.Fprintf(w, "\n ID : %s", customerData.GetValueUniqIdCust())
			fmt.Fprintf(w, "\n NAME : %s", customerData.GetValueNameCust())
			fmt.Println("New ID : ", customerData.GetValueUniqIdCust())
		}
		if userData.GetValueStatusUsr() == "Owner" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Only Employee can set new Customer")
			fmt.Println("Login as ", userData.GetValueStatusUsr())
			fmt.Println("cannot access settlement")
		}
	}
}

func NewSession(w http.ResponseWriter, r *http.Request) {
	user := handler_redis.GetUserRedis(ctx)
	if user.GetValueIdUsr() != 0 {
		if user.GetValueStatusUsr() == "Employee" {
			handler_redis.HandlerShopRedis.Conn.FlushAll(ctx)
			handler_redis.SetUserRedis(ctx, user)

			fmt.Println("New SESSION Created...")

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "New Session OK...")
		}
		if user.GetValueStatusUsr() == "Owner" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "no access...")
		}
	}
	if user.GetValueIdUsr() == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "PLEASE LOGIN FIRST.")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	handler_redis.HandlerShopRedis.Conn.FlushAll(ctx)
	fmt.Println("LOGOUT SUCCESS...")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "LOGOUT SUCCESS...")
}
