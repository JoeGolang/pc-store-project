package custhandler

import (
	"github.com/gorilla/mux"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (cs *CustomerHandler) DeleteDataCustomer(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)

	dataCustomer, errGet := cs.repoCustomer.GetCustomerById(cs.ctx, vars["idcust"])
	if errGet != nil {
		respErr, _ := http_response.MapResponseCustomer(nil, http.StatusInternalServerError, errGet.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write(respErr)
		return
	}

	errDelete := cs.repoCustomer.DeleteCustomer(cs.ctx, string(dataCustomer.GetValueUniqIdCust()))
	if errDelete != nil {
		respErr, _ := http_response.MapResponseCustomer(nil, http.StatusInternalServerError, errDelete.Error())
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
