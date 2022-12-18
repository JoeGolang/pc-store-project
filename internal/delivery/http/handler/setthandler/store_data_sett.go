package setthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	settlement2 "pc-shop-final-project/domain/entity/settlement"
	"pc-shop-final-project/internal/delivery/http/http_request"
)

func (st *SettHandler) StoreDataSettlement(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestSettlement
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	settlement, err := settlement2.NewSettlement(settlement2.DTOSettlement{
		Id:   req.Id,
		Code: req.Code,
		//User: req.User,
		//Customer: req.Customer,
		//Product:    req.,
		//Coupon:     req.Coupon,
		TotalPrice: req.TotalPrice,
		StatusTrns: req.StatusTrns,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error build data"))
		return
	}

	errInsert := st.repoSett.CreateSettle(st.ctx, settlement)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInsert.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SUCCESS INSERT DATA")
	return
}
