package http_response

import (
	"encoding/json"
	"pc-shop-final-project/domain/entity"
)

type ResponseCustomerJson struct {
	UNIQ_ID   string `json:"UNIQ_ID"`
	NAME      string `json:"NAME"`
	JOIN_DATE string `json:"JOIN_DATE"`
}

type CustomCustReponseCollection struct {
	Status *Status
	Data   []*ResponseCustomerJson
}

type CustomCustReponseSingle struct {
	Status *Status
	Data   *ResponseCustomerJson
}

func MapResponseListCustomer(dataCustomer []*entity.Customer, code int, message string) ([]byte, error) {
	listRespCustomer := make([]*ResponseCustomerJson, 0)
	for _, dataCust := range dataCustomer {
		respCustomer := &ResponseCustomerJson{
			UNIQ_ID:   dataCust.GetValueUniqIdCust(),
			NAME:      dataCust.GetValueNameCust(),
			JOIN_DATE: dataCust.GetValueJoinDateCust(),
		}
		listRespCustomer = append(listRespCustomer, respCustomer)
	}

	httpResponse := &CustomCustReponseCollection{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: listRespCustomer,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

func MapResponseCustomer(dataCustomer *entity.Customer, code int, message string) ([]byte, error) {
	var resp *ResponseCustomerJson
	if dataCustomer != nil {
		resp = &ResponseCustomerJson{
			UNIQ_ID:   dataCustomer.GetValueUniqIdCust(),
			NAME:      dataCustomer.GetValueNameCust(),
			JOIN_DATE: dataCustomer.GetValueJoinDateCust(),
		}
	}

	httpResponse := &CustomCustReponseSingle{
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
