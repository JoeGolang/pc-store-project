package http_response

import (
	"encoding/json"
	"pc-shop-final-project/domain/entity"
)

type ResponseCouponJson struct {
	UNIQ_ID    string `json:"UNIQ_ID"`
	STATUS_USE bool   `json:"STATUS_USE"`
}

type CustomCouponReponseCollection struct {
	Status *Status
	Data   []*ResponseCouponJson
}

type CustomCouponReponseSingle struct {
	Status *Status
	Data   *ResponseCouponJson
}

func MapResponseListCoupon(dataCoupon []*entity.UniqCoupon, code int, message string) ([]byte, error) {
	listRespCoupon := make([]*ResponseCouponJson, 0)
	for _, dataCoupon := range dataCoupon {
		respCoupon := &ResponseCouponJson{
			UNIQ_ID:    dataCoupon.GetValueUniqIdCoupon(),
			STATUS_USE: dataCoupon.GetValueStatusCoupon(),
		}
		listRespCoupon = append(listRespCoupon, respCoupon)
	}

	httpResponse := &CustomCouponReponseCollection{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: listRespCoupon,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

func MapResponseCoupon(dataCoupon *entity.UniqCoupon, code int, message string) ([]byte, error) {
	var resp *ResponseCouponJson
	if dataCoupon != nil {
		resp = &ResponseCouponJson{
			UNIQ_ID:    dataCoupon.GetValueUniqIdCoupon(),
			STATUS_USE: dataCoupon.GetValueStatusCoupon(),
		}
	}

	httpResponse := &CustomCouponReponseSingle{
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
