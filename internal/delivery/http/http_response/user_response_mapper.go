package http_response

import (
	"encoding/json"
	"pc-shop-final-project/domain/entity"
)

type Status struct {
	Code    int
	Message string
}

type ResponseUserJson struct {
	Id     int    `json:"ID_USER"`
	Name   string `json:"NAME"`
	Outlet string `json:"OUTLET_CODE"`
	Status string `json:"STATUS"`
}

type ReponseUsers struct {
	Status *Status
	Data   []*ResponseUserJson
}

type ReponseUser struct {
	Status *Status
	Data   *ResponseUserJson
}

func MapResponseListUser(dataUsers []*entity.User, code int, message string) ([]byte, error) {
	listRespUser := make([]*ResponseUserJson, 0)
	for _, dataUser := range dataUsers {
		respUser := &ResponseUserJson{
			Id:     dataUser.GetValueIdUsr(),
			Name:   dataUser.GetValueNameUsr(),
			Outlet: dataUser.GetValueOutletCodeUsr(),
			Status: dataUser.GetValueStatusUsr(),
		}
		listRespUser = append(listRespUser, respUser)
	}

	httpResponse := &ReponseUsers{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: listRespUser,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

func MapResponseUser(dataUser *entity.User, code int, message string) ([]byte, error) {
	var resp *ResponseUserJson
	if dataUser != nil {
		resp = &ResponseUserJson{
			Id:     dataUser.GetValueIdUsr(),
			Name:   dataUser.GetValueNameUsr(),
			Outlet: dataUser.GetValueOutletCodeUsr(),
			Status: dataUser.GetValueStatusUsr(),
		}
	}

	httpResponse := &ReponseUser{
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
