package mapper

import (
	"github.com/rocketlaunchr/dbq/v2"
	"pc-shop-final-project/domain/entity/customer"
	"pc-shop-final-project/internal/delivery/http/models"
)

func DataCustomerDbToEntity(dataDTO customer.DTOCustomer) (*customer.Customer, error) {
	customer, err := customer.NewCustomer(dataDTO, "BOGOR")
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func CustomerEntityToModel(cst *customer.Customer) *models.ModelCustomer {
	return &models.ModelCustomer{
		UNIQ_ID:   cst.GetValueUniqIdCust(),
		NAME:      cst.GetValueNameCust(),
		JOIN_DATE: cst.GetValueJoinDateCust(),
	}
}

func CustomerEntityToDbqStruct(customer *customer.Customer) []interface{} {
	dbqStruct := dbq.Struct(CustomerEntityToModel(customer))
	return dbqStruct
}

func CustomerModelToEntity(model *models.ModelCustomer) (*customer.Customer, error) {
	customer, err := customer.NewCustomer(customer.DTOCustomer{
		UniqId:   model.UNIQ_ID,
		Name:     model.NAME,
		JoinDate: model.JOIN_DATE,
	}, "BOGOR")

	if err != nil {
		return nil, err
	}

	return customer, nil
}
