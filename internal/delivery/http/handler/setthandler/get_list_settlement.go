package setthandler

import (
	"net/http"
	"pc-shop-final-project/internal/delivery/http/http_response"
)

func (st *SettHandler) GetListSettlement(w http.ResponseWriter, r *http.Request) {
	listSettlement, err := st.repoSett.ReadSettle(st.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := http_response.MapResponseListSettlement(listSettlement, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
