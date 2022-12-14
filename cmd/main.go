package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/handler"
	"pc-shop-final-project/internal/repository/mysql"
	mysqlCon "pc-shop-final-project/pkg/mysql_connection"
)

var (
	ctx       = context.Background()
	mysqlConn = mysqlCon.InitMysqlDB()
	repoUser  = mysql.NewUserMysql(mysqlConn)
)

func main() {
	r := mux.NewRouter()
	handlerUser := handler.NewUserHandler(ctx, repoUser)

	r.HandleFunc("/", ParamHandlerWithoutInput).Methods(http.MethodGet)
	//r.HandleFunc("/create-inventory", handlerInventory.Store)
	r.HandleFunc("/list-user", handlerUser.GetListUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handlerUser.GetListUser).Methods(http.MethodGet)

	http.HandleFunc("/test", ParamHandlerWithoutInput)
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
