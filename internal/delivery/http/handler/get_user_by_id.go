package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (usr *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataUser, err := usr.repoUser.GetUserById(usr.ctx, vars["id"])

	response, err := http_response.MapResponseUser(dataUser, 200, "Success")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
