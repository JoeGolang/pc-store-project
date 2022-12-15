package userhandler

import (
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (us *UserHandler) GetListUser(w http.ResponseWriter, r *http.Request) {
	listUser, err := us.repoUser.GetListUser(us.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := http_response.MapResponseListUser(listUser, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
