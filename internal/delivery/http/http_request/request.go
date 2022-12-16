package http_request

type RequestUser struct {
	Id     int    `json:"ID_USER"`
	Name   string `json:"NAME"`
	Outlet string `json:"OUTLET_CODE"`
	Status string `json:"STATUS"`
}

type RequestCustomer struct {
	//UniqId   string `json:"UNIQ_ID"`
	Name string `json:"NAME"`
	//JoinDate string `json:"JOIN_DATE"`
}

type RequestUpdCustomer struct {
	UniqId string `json:"UNIQ_ID"`
	Name   string `json:"NAME"`
	//JoinDate string `json:"JOIN_DATE"`
}
