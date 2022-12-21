package handlerHTTP

import (
	"fmt"
	"net/http"
	"pc-shop-final-project/domain/entity"
	"pc-shop-final-project/internal/delivery/http/http_response"
	"strconv"

	"github.com/gorilla/mux"
	//"pc-shop-final-project/domain/entity"
	handler "pc-shop-final-project/internal/repository/controller"
	handler_redis "pc-shop-final-project/internal/repository/controller/redis_control"
)

func ParamDeleteIdUser(w http.ResponseWriter, r *http.Request) {
	var (
		user *entity.User
	)
	vars := mux.Vars(r)
	id := vars["ID"]
	intId, _ := strconv.Atoi(id)

	userData := handler_redis.GetUserRedis(ctx)
	if userData.GetValueIdUsr() == 0 {
		fmt.Fprintf(w, "LOGIN TO PROCEED...")
		fmt.Println("LOGIN TO PROCEED...")
		w.WriteHeader(http.StatusBadRequest)
	}

	if userData.GetValueIdUsr() != 0 {
		if userData.GetValueStatusUsr() == "Employee" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Only Owner can delete user")
			fmt.Println("Login as ", userData.GetValueStatusUsr())
			fmt.Println("cannot delete user")
		}
		if userData.GetValueStatusUsr() == "Owner" {
			//

			dataUsers := handler.ReadUser(ctx)
			for _, dataUser := range dataUsers {
				if dataUser.GetValueIdUsr() == intId {
					fmt.Println("WILL DELETE : ", intId)
					user = dataUser

					errDelete := handler.DeleteUser(ctx, user.GetValueIdUsr())
					if errDelete != nil {
						respErr, _ := http_response.MapResponseUser(nil, http.StatusInternalServerError, errDelete.Error())
						w.WriteHeader(http.StatusInternalServerError)
						w.Write(respErr)
						return
					}

					resp, errMap := http_response.MapResponseUser(nil, http.StatusOK, "SUCCESS DELETE DATA")
					if errMap != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte(errMap.Error()))
						return
					}

					w.WriteHeader(200)
					w.Write(resp)
					return
				}
			}
			if user == nil {
				fmt.Println("No User Found...", intId)
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "No User Found...", intId)
			}
		}
	}
}

func ParamDeleteIdCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		cust *entity.Customer
	)
	vars := mux.Vars(r)
	idcust := vars["ID"]

	userData := handler_redis.GetUserRedis(ctx)
	if userData.GetValueIdUsr() == 0 {
		fmt.Fprintf(w, "LOGIN TO PROCEED...")
		fmt.Println("LOGIN TO PROCEED...")
		w.WriteHeader(http.StatusBadRequest)
	}

	if userData.GetValueIdUsr() != 0 {
		if userData.GetValueStatusUsr() == "Employee" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Only Owner can delete user")
			fmt.Println("Login as ", userData.GetValueStatusUsr())
			fmt.Println("cannot delete user")
		}
		if userData.GetValueStatusUsr() == "Owner" {
			//

			dataCusts := handler.ReadCustomer(ctx)
			for _, dataCust := range dataCusts {
				if dataCust.GetValueUniqIdCust() == idcust {
					fmt.Println("WILL DELETE : ", idcust)
					cust = dataCust

					errDelete := handler.DeleteCustomer(ctx, cust.GetValueUniqIdCust())
					if errDelete != nil {
						respErr, _ := http_response.MapResponseUser(nil, http.StatusInternalServerError, errDelete.Error())
						w.WriteHeader(http.StatusInternalServerError)
						w.Write(respErr)
						return
					}

					resp, errMap := http_response.MapResponseUser(nil, http.StatusOK, "SUCCESS DELETE DATA")
					if errMap != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte(errMap.Error()))
						return
					}

					w.WriteHeader(200)
					w.Write(resp)
					return
				}
			}
			if cust == nil {
				fmt.Println("No Customer Found...", idcust)
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "No Customer Found...", idcust)
			}
		}
	}
}
