package http_request

type RequestUser struct {
	ID_USER     int    `json:"ID_USER"`
	NAME        string `json:"NAME"`
	OUTLET_CODE string `json:"OUTLET_CODE"`
	STATUS      string `json:"STATUS"`
}

type RequestCustomer struct {
	UNIQ_ID   string `json:"UNIQ_ID"`
	NAME      string `json:"NAME"`
	JOIN_DATE string `json:"JOIN_DATE"`
}
