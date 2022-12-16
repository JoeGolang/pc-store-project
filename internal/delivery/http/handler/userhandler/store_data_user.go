package userhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	user2 "pc-shop-final-project/domain/entity/user"
	"pc-shop-final-project/internal/delivery/http/http_request"
)

func (usr *UserHandler) StoreDataUser(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestUser
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	fmt.Println(req)

	user, err := user2.NewUser(user2.DTOUser{
		Id:         req.Id,
		Name:       req.Name,
		OutletCode: req.Outlet,
		Status:     req.Status,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error build data"))
		return
	}

	errInsert := usr.repoUser.InsertDataUser(usr.ctx, user)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInsert.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SUCCESS INSERT DATA")
	return
}
