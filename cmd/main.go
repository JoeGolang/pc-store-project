package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	handlerHTTP "pc-shop-final-project/internal/repository/controller/http_control"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", handlerHTTP.ParamHandlerWithoutInput).Methods(http.MethodGet)
	r.HandleFunc("/user", handlerHTTP.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/NewUser", handlerHTTP.ParamNewUser).Methods(http.MethodPost)
	r.HandleFunc("/customer", handlerHTTP.GetCustomers).Methods(http.MethodGet)
	r.HandleFunc("/IdUser:{ID}", handlerHTTP.ParamLoginIdUser).Methods(http.MethodGet)
	r.HandleFunc("/DeleteIdUser:{ID}", handlerHTTP.ParamDeleteIdUser).Methods(http.MethodDelete)
	r.HandleFunc("/IdCustomer:{ID}", handlerHTTP.ParamIdCustomer).Methods(http.MethodGet)
	r.HandleFunc("/DeleteIdCustomer:{ID}", handlerHTTP.ParamDeleteIdCustomer).Methods(http.MethodDelete)
	r.HandleFunc("/NewCustomer", handlerHTTP.ParamNewCustomer).Methods(http.MethodPost)
	r.HandleFunc("/Logout", handlerHTTP.Logout).Methods(http.MethodGet)
	r.HandleFunc("/NewSession", handlerHTTP.NewSession).Methods(http.MethodGet)
	r.HandleFunc("/ScannedItem:{IdItem}", handlerHTTP.ParamAddItemPurchase).Methods(http.MethodGet)
	r.HandleFunc("/BuildSettle", handlerHTTP.ParamSettleProcess).Methods(http.MethodGet)
	r.HandleFunc("/GetCoupon", handlerHTTP.ParamOwnerGetCoupon).Methods(http.MethodGet)
	r.HandleFunc("/GetTransaction", handlerHTTP.ParamOwnerGetTransaction).Methods(http.MethodGet)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
