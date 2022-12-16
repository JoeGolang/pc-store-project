package models

type ModelUser struct {
	Id     int    `dbq:"ID_USER"`
	Name   string `dbq:"NAME"`
	Outlet string `dbq:"OUTLET_CODE"`
	Status string `dbq:"STATUS"`
}

func GetUserTableName() string {
	return "user"
}

func GetUserFillableTable() []string {
	return []string{
		"ID_USER",
		"NAME",
		"OUTLET_CODE",
		"STATUS",
	}
}
