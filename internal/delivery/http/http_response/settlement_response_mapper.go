package http_response

import (
	"encoding/json"
	"pc-shop-final-project/domain/entity/settlement"
)

type ResponseSettlementJson struct {
	//CUSTOMER         int `json:"CUSTOMER"`
	CODE_TRANSACTION string `json:"CODE_TRANSACTION"`
	TOTAL_PRICE      int    `json:"TOTAL_PRICE"`
}

type CustomSettReponseCollection struct {
	Status *Status
	Data   []*ResponseSettlementJson
}

type CustomSettReponseSingle struct {
	Status *Status
	Data   *ResponseSettlementJson
}

func MapResponseListSettlement(dataSettlement []*settlement.Settlement, code int, message string) ([]byte, error) {
	listRespSettlement := make([]*ResponseSettlementJson, 0)
	for _, dataSett := range dataSettlement {
		respSettlement := &ResponseSettlementJson{
			//CUSTOMER:   dataSett.GetValueCustomerSett(),
			CODE_TRANSACTION: dataSett.GetValueCodeSett(),
			TOTAL_PRICE:      dataSett.GetValueTotalPriceSett(),
		}
		listRespSettlement = append(listRespSettlement, respSettlement)
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

func MapResponseSettlement(dataSettlement *settlement.Settlement, code int, message string) ([]byte, error) {
	var resp *ResponseSettlementJson
	if dataSettlement != nil {
		resp = &ResponseSettlementJson{
			//CUSTOMER:   dataSettlement.GetValueCustomerSett(),
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
