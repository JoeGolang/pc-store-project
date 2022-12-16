package custhandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_request"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (cs *CustomerHandler) UpdateDataCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestUpdCustomer
		decoder = json.NewDecoder(r.Body)
		vars    = mux.Vars(r)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	dataCustomer, errGet := cs.repoCustomer.GetCustomerById(cs.ctx, vars["idcust"])
	if errGet != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errGet.Error()))
		return
	}

	dataCustomer.SetUpdateDataCust(req)

	errUpdate := cs.repoCustomer.UpdateCustomerById(cs.ctx, dataCustomer, vars["idcust"])
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errUpdate.Error()))
		return
	}

	resp, errMap := http_response.MapResponseCustomer(nil, http.StatusOK, "SUCCESS UPDATE DATA")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMap.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
	return
}
