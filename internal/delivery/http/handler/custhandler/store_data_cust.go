package custhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	customer2 "pc-shop-final-project/domain/entity/customer"
	"pc-shop-final-project/internal/delivery/http/http_request"
)

func (cs *CustomerHandler) StoreDataCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestCustomer
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	store := "BGR1" //HARDCODED

	customer, err := customer2.NewCustomer(customer2.DTOCustomer{
		//UniqId:   req.UniqId,
		Name: req.Name,
		//JoinDate: req.JoinDate,
	}, store)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error build data"))
		return
	}

	errInsert := cs.repoCustomer.InsertDataCustomer(cs.ctx, customer)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInsert.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SUCCESS INSERT DATA")
	return
}
