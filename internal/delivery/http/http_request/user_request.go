package http_request

type RequestUser struct {
	ID_USER     int    `json:"ID_USER"`
	NAME        string `json:"NAME"`
	OUTLET_CODE string `json:"OUTLET_CODE"`
	STATUS      string `json:"STATUS"`
}
