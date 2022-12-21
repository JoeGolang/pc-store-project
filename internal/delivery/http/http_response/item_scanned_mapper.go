package http_response

import (
	"encoding/json"
	"pc-shop-final-project/domain/entity"
)

type ResponseItemJson struct {
	IdProduct  int    `json:"id_product"`
	Qty        int    `json:"qty"`
	ProdDetail string `json:"prod_detail"`
	Price      int    `json:"price"`
}

type CustomItemReponseCollection struct {
	Status *Status
	Data   []*ResponseItemJson
	Total  int
}

type CustomItemReponseSingle struct {
	Status *Status
	Data   *ResponseItemJson
}

func MapResponseListItem(dataItem []*entity.SettlePurchase, product string, price int, total int, code int, message string) ([]byte, error) {
	listresp := make([]*ResponseItemJson, 0)
	for _, dataItem := range dataItem {
		respItem := &ResponseItemJson{
			IdProduct:  dataItem.GetValueIdProduct(),
			Qty:        dataItem.GetValueQtyProduct(),
			ProdDetail: product,
			Price:      price,
		}
		listresp = append(listresp, respItem)
	}

	httpResponse := &CustomItemReponseCollection{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data:  listresp,
		Total: total,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
