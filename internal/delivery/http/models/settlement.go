package models

type ModelSettlement struct {
	//CUSTOMER	string
	CODE_TRANSACTION string `dbq:"CODE_TRANSACTION"`
	TOTAL_PRICE      int    `dbq:"TOTAL_PRICE"`
}

func GetSettlementTableName() string {
	return "settlement"
}

func GetSettlementFillableTable() []string {
	return []string{
		"CODE_TRANSACTION",
		"TOTAL_PRICE",
	}
}
