package handlerHTTP

import (
	"context"
	"fmt"
	"net/http"
	"pc-shop-final-project/domain/entity"
	"pc-shop-final-project/internal/delivery/http/http_response"
	handler "pc-shop-final-project/internal/repository/controller"
	handler_redis "pc-shop-final-project/internal/repository/controller/redis_control"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	ctx = context.Background()
)

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "/user -> show users\n/IdUser:(enter id user) -> login")
}

func ParamLoginIdUser(w http.ResponseWriter, r *http.Request) {
	var (
		user *entity.User
	)
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Input ID User : %v\n", vars["ID"])
	id := vars["ID"]
	intId, _ := strconv.Atoi(id)

	fmt.Println("ID User : ", intId)
	dataUsers := handler.ReadUser(ctx)
	for _, dataUser := range dataUsers {
		if dataUser.GetValueIdUsr() == intId {
			fmt.Println("LOGIN SUCCESS...")
			fmt.Println("Welcome ", dataUser.GetValueNameUsr(), " , ", dataUser.GetValueStatusUsr())
			user = dataUser
			handler_redis.SetUserRedis(ctx, dataUser)
			fmt.Fprintf(w, "LOGIN SUCCESS...")
			fmt.Fprintln(w, "Welcome ", dataUser.GetValueNameUsr(), ",\n you logged in as ", dataUser.GetValueStatusUsr())
			w.WriteHeader(http.StatusFound)

			response, errMap := http_response.MapResponseUser(user, 200, "Success")
			if errMap != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Error mapping data"))
			}

			w.WriteHeader(200)
			w.Write(response)

		}
	}
	if user == nil {
		handler_redis.HandlerShopRedis.Conn.FlushAll(ctx)
		fmt.Println("LOGIN FAILED...")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "LOGIN FAILED...")
	}
}

func ParamIdCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		cust *entity.Customer
	)
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Input ID Customer : %v\n", vars["ID"])
	id := vars["ID"]

	userData := handler_redis.GetUserRedis(ctx)
	if userData.GetValueIdUsr() != 0 {
		if userData.GetValueStatusUsr() == "Employee" {
			fmt.Println("Login as ", userData.GetValueStatusUsr())
			fmt.Println("ID Customer : ", id)
			dataCusts := handler.ReadCustomer(ctx)
			for _, dataCust := range dataCusts {
				if dataCust.GetValueUniqIdCust() == id {
					fmt.Println("CUSTOMER DATA FOUND...")
					cust = dataCust
					handler_redis.SetCustomerRedis(ctx, dataCust)
					w.WriteHeader(http.StatusFound)
					fmt.Fprintf(w, "Customer Data FOUND...\n\n Name : %s \n ID : %s \n Join Date : %s \n", dataCust.GetValueNameCust(), dataCust.GetValueUniqIdCust(), dataCust.GetValueJoinDateCust())

					response, errMap := http_response.MapResponseCustomer(cust, 200, "Success")
					if errMap != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("Error mapping data"))
					}

					w.WriteHeader(200)
					w.Write(response)

				}
			}
			if cust.GetValueUniqIdCust() == "" {
				fmt.Println("CUSTOMER DATA NOT EXIST...")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Customer Data NOT EXIST...")
			}
		}
		if userData.GetValueStatusUsr() == "Owner" {
			fmt.Println("Login as ", userData.GetValueStatusUsr())
			fmt.Println("cannot access settlement")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "no access...")
		}
	}
	if userData.GetValueIdUsr() == 0 {
		fmt.Fprintf(w, "LOGIN TO PROCEED...")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("LOGIN TO PROCEED...")
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	listUser := handler.ReadUser(ctx)

	response, errMap := http_response.MapResponseListUser(listUser, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	listCustomer := handler.ReadCustomer(ctx)

	response, errMap := http_response.MapResponseListCustomer(listCustomer, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
