package handlerHTTP

import (
	"fmt"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
	"strconv"

	"github.com/gorilla/mux"
	//"pc-shop-final-project/domain/entity"
	handler "pc-shop-final-project/internal/repository/controller"
	handler_redis "pc-shop-final-project/internal/repository/controller/redis_control"
)

func ParamDeleteIdUser(w http.ResponseWriter, r *http.Request) {
	var flag bool
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
			dataUsers := handler.ReadUser(ctx)
			for _, dataUser := range dataUsers {
				if dataUser.GetValueIdUsr() == intId {
					flag = true
					fmt.Println("WILL DELETE : ", intId)

					errDelete := handler.DeleteUser(ctx, dataUser.GetValueIdUsr())
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
			if !flag {
				fmt.Println("No match ID found...")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintln(w, "No match ID found...")
			}
		}
	}
}

func ParamDeleteIdCustomer(w http.ResponseWriter, r *http.Request) {
	var flag bool
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
			dataCusts := handler.ReadCustomer(ctx)
			for _, dataCust := range dataCusts {
				if dataCust.GetValueUniqIdCust() == idcust {
					flag = true
					fmt.Println("WILL DELETE : ", idcust)

					errDelete := handler.DeleteCustomer(ctx, dataCust.GetValueUniqIdCust())
					if errDelete != nil {
						respErr, _ := http_response.MapResponseUser(nil, http.StatusInternalServerError, errDelete.Error())
						w.WriteHeader(http.StatusInternalServerError)
						w.Write(respErr)
						return
					}

					resp, errMap := http_response.MapResponseCustomer(nil, http.StatusOK, "SUCCESS DELETE DATA")
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
		}
		if !flag {
			fmt.Println("No match ID found...")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "No match ID found...")
		}
	}
}
