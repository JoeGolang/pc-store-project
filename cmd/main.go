package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/handler/couponhandler"
	"pc-shop-final-project/internal/delivery/http/handler/custhandler"
	"pc-shop-final-project/internal/delivery/http/handler/userhandler"
	"pc-shop-final-project/internal/repository/mysql"
	mysqlCon "pc-shop-final-project/pkg/mysql_connection"
)

var (
	ctx        = context.Background()
	mysqlConn  = mysqlCon.InitMysqlDB()
	repoUser   = mysql.NewUserMysql(mysqlConn)
	repoCust   = mysql.NewCustomerMysql(mysqlConn)
	repoCoupon = mysql.NewCustomerMysql(mysqlConn)
)

func main() {
	r := mux.NewRouter()
	handlerUser := userhandler.NewUserHandler(ctx, repoUser)
	handlerCust := custhandler.NewCustomerHandler(ctx, repoCust)
	handlerCoupon := couponhandler.NewCouponHandler(ctx, repoCoupon)

	r.HandleFunc("/", ParamHandlerWithoutInput).Methods(http.MethodGet)
	r.HandleFunc("/usercreate", handlerUser.StoreDataUser).Methods(http.MethodPost)
	r.HandleFunc("/user", handlerUser.GetListUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handlerUser.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/userupdate", ParamHandlerWithoutInput).Methods(http.MethodPut)
	r.HandleFunc("/userdelete", ParamHandlerWithoutInput).Methods(http.MethodDelete)

	r.HandleFunc("/custcreate", handlerCust.StoreDataCustomer).Methods(http.MethodPost)
	r.HandleFunc("/cust", handlerCust.GetListCustomer).Methods(http.MethodGet)
	r.HandleFunc("/cust/{id}", handlerCust.GetCustomerById).Methods(http.MethodGet)
	//r.HandleFunc("/custupdate", ParamHandlerWithoutInput).Methods(http.MethodPut)
	//r.HandleFunc("/custdelete", ParamHandlerWithoutInput).Methods(http.MethodDelete)

	r.HandleFunc("/coupon", handlerCoupon.GetListCoupon).Methods(http.MethodGet)
	r.HandleFunc("/coupon/{id}", handlerCoupon.GetCouponById).Methods(http.MethodGet)
	//r.HandleFunc("/custupdate", ParamHandlerWithoutInput).Methods(http.MethodPut)
	//r.HandleFunc("/custdelete", ParamHandlerWithoutInput).Methods(http.MethodDelete)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)

	//dummyInventory := dummy.MakeDummyInventory()
	//for _, dumI := range dummyInventory {
	//	handler2.CreateInventory(ctx, dumI)
	//}
	//Inventories := handler2.ReadInventory(ctx)
	//fmt.Println(Inventories)
	//for _, inv := range Inventories {
	//	fmt.Println(inv.GetValueIdInv())
	//	fmt.Println(inv.GetValueProductNameInv())
	//	fmt.Println(inv.GetValueBrandInv())
	//	fmt.Println(inv.GetvaluePriceInv())
	//	fmt.Println(inv.GetValueCategoryInv())
	//	fmt.Println("____________________")
	//}

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
	//}
}

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
}
