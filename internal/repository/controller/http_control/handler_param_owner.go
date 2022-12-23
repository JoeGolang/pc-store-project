package handlerHTTP

import (
	"fmt"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
	handler "pc-shop-final-project/internal/repository/controller"
	handler_redis "pc-shop-final-project/internal/repository/controller/redis_control"
)

func ParamOwnerGetCoupon(w http.ResponseWriter, r *http.Request) {
	userData := handler_redis.GetUserRedis(ctx)
	if userData.GetValueIdUsr() != 0 {
		if userData.GetValueStatusUsr() == "Employee" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "no access...")
		}
		if userData.GetValueStatusUsr() == "Owner" {

			w.WriteHeader(http.StatusOK)
			_, listCoupon := handler.ReadCoupon(ctx)

			response, errMap := http_response.MapResponseListCoupon(listCoupon, 200, "Success")
			if errMap != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(200)
			w.Write(response)

		}
	}
}

func ParamOwnerGetTransaction(w http.ResponseWriter, r *http.Request) {
	userData := handler_redis.GetUserRedis(ctx)
	if userData.GetValueIdUsr() != 0 {
		if userData.GetValueStatusUsr() == "Employee" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "no access...")
		}
		if userData.GetValueStatusUsr() == "Owner" {
			w.WriteHeader(http.StatusOK)
			listSettlement := handler.ReadSettle(ctx)
			responseSet, errMap := http_response.MapResponseListSettlement(listSettlement, 200, "Success")
			if errMap != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(200)
			w.Write(responseSet)

		}
	}
}
