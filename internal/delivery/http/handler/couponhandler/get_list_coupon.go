package couponhandler

import (
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (cp *CouponHandler) GetListCoupon(w http.ResponseWriter, r *http.Request) {
	listCoupon, _, err := cp.repoCoupon.GetListCoupon(cp.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := http_response.MapResponseListCoupon(listCoupon, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
