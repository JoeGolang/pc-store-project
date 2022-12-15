package userhandler

import (
	"github.com/gorilla/mux"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (usr *handler.UserHandler) DeleteDataUser(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)

	dataUser, errGet := usr.repoUser.GetUserById(usr.ctx, vars["id"])
	if errGet != nil {
		respErr, _ := http_response.MapResponseUser(nil, http.StatusInternalServerError, errGet.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write(respErr)
		return
	}

	errDelete := usr.repoUser.DeleteUserById(usr.ctx, string(dataUser.GetValueIdUsr()))
	if errDelete != nil {
		respErr, _ := http_response.MapResponseUser(nil, http.StatusInternalServerError, errDelete.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	resp, errMap := http_response.MapResponseUser(nil, http.StatusOK, "SUCCESS DELETE DATA")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMap.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
	return
}
