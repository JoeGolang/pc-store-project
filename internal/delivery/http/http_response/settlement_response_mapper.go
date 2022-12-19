package http_response

import (
	"encoding/json"
	"pc-shop-final-project/domain/entity"
)

type ResponseSettlementJson struct {
	ID                 int    `json:"ID"`
	COUPON_ID          string `json:"COUPON_ID"`
	CODE_TRANSACTION   string `json:"CODE_TRANSACTION"`
	TOTAL_PRICE        int    `json:"TOTAL_PRICE"`
	STATUS_TRANSACTION bool   `json:"STATUS_TRANSACTION"`
}

type CustomSettReponseCollection struct {
	Status *Status
	Data   []*ResponseSettlementJson
}

type CustomSettReponseSingle struct {
	Status *Status
	Data   *ResponseSettlementJson
}

func MapResponseListSettlement(dataSettlement []*entity.Settlement, code int, message string) ([]byte, error) {
	listRespSettlement := make([]*ResponseSettlementJson, 0)
	for _, dataSett := range dataSettlement {
		respSett := &ResponseSettlementJson{
			ID:               dataSett.GetValueIdSett(),
			CODE_TRANSACTION: dataSett.GetValueCodeSett(),
			TOTAL_PRICE:      dataSett.GetValueTotalPriceSett(),
		}
		listRespSettlement = append(listRespSettlement, respSett)
	}

	httpResponse := &CustomSettReponseCollection{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: listRespSettlement,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

func MapResponseSettlement(dataSettlement *entity.Settlement, code int, message string) ([]byte, error) {
	var resp *ResponseSettlementJson
	if dataSettlement != nil {
		resp = &ResponseSettlementJson{
			ID:               dataSettlement.GetValueIdSett(),
			CODE_TRANSACTION: dataSettlement.GetValueCodeSett(),
			TOTAL_PRICE:      dataSettlement.GetValueTotalPriceSett(),
		}
	}

	httpResponse := &CustomSettReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: resp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
