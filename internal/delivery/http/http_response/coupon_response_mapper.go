package http_response

import (
	"encoding/json"
	"pc-shop-final-project/domain/entity"
)

type ResponseCouponJson struct {
	ID_COUPON     int    `json:"ID_COUPON"`
	ID_CUSTOMER   string `json:"ID_CUSTOMER"`
	GENERATE_DATE string `json:"GENERATE_DATE"`
	REVENUE       int    `json:"REVENUE"`
}

type CustomCouponReponseCollection struct {
	Status *Status
	Data   []*ResponseCouponJson
}

type CustomCouponReponseSingle struct {
	Status *Status
	Data   *ResponseCouponJson
}

func MapResponseListCoupon(dataCoupon []*entity.Coupon, code int, message string) ([]byte, error) {
	listRespCoupon := make([]*ResponseCouponJson, 0)
	for _, dataCoupon := range dataCoupon {
		respCoupon := &ResponseCouponJson{
			ID_COUPON:     dataCoupon.GetValueIdCpn(),
			ID_CUSTOMER:   dataCoupon.GetValueIdCustomerCpn(),
			GENERATE_DATE: dataCoupon.GetValueGenDateCpn(),
			REVENUE:       dataCoupon.GetValueRevenueCpn(),
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

func MapResponseCoupon(dataCoupon *entity.Coupon, code int, message string) ([]byte, error) {
	var resp *ResponseCouponJson
	if dataCoupon != nil {
		resp = &ResponseCouponJson{
			ID_COUPON:     dataCoupon.GetValueIdCpn(),
			ID_CUSTOMER:   dataCoupon.GetValueIdCustomerCpn(),
			GENERATE_DATE: dataCoupon.GetValueGenDateCpn(),
			REVENUE:       dataCoupon.GetValueRevenueCpn(),
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
