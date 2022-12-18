package setthandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_request"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (cs *SettHandler) UpdateDataSettlement(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestSettlement
		decoder = json.NewDecoder(r.Body)
		vars    = mux.Vars(r)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	dataSettlement, errGet := cs.repoSett.ReadSettlementById(cs.ctx, vars["idsett"])
	if errGet != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errGet.Error()))
		return
	}

	dataSettlement.SetUpdateSettlementData(req)

	errUpdate := cs.repoSett.UpdateSettle(cs.ctx, vars["idsett"])
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errUpdate.Error()))
		return
	}

	resp, errMap := http_response.MapResponseSettlement(nil, http.StatusOK, "SUCCESS UPDATE DATA")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMap.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
	return
}
