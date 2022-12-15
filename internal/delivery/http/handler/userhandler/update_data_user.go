package userhandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_request"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (usr *UserHandler) UpdateDataUser(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestUser
		decoder = json.NewDecoder(r.Body)
		vars    = mux.Vars(r)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	dataUser, errGet := usr.repoUser.GetUserById(usr.ctx, vars["id"])
	if errGet != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errGet.Error()))
		return
	}

	dataUser.SetUpdateData(req)

	errUpdate := usr.repoUser.UpdateUserById(usr.ctx, dataUser, vars["id"])
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errUpdate.Error()))
		return
	}

	resp, errMap := http_response.MapResponseUser(nil, http.StatusOK, "SUCCESS UPDATE DATA")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMap.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
	return
}
