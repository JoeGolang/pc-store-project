package models

type ModelCustomer struct {
	UNIQ_ID   string `dbq:"UNIQ_ID"`
	NAME      string `dbq:"NAME"`
	JOIN_DATE string `dbq:"JOIN_DATE"`
}

func GetCustomerTableName() string {
	return "customer"
}

func GetCustomerFillableTable() []string {
	return []string{
		"UNIQ_ID",
		"NAME",
		"JOIN_DATE",
	}
}
