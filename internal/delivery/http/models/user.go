package models

type ModelUser struct {
	ID_USER     int    `dbq:"ID_USER"`
	NAME        string `dbq:"NAME"`
	OUTLET_CODE string `dbq:"OUTLET_CODE"`
	STATUS      string `dbq:"STATUS"`
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
