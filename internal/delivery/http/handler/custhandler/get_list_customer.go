package custhandler

import (
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (cs *CustomerHandler) GetListCustomer(w http.ResponseWriter, r *http.Request) {
	listCustomer, err := cs.repoCustomer.GetListCustomer(cs.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := http_response.MapResponseListCustomer(listCustomer, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
