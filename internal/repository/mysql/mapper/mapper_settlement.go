package mapper

import (
	"github.com/rocketlaunchr/dbq/v2"
	settlement2 "pc-shop-final-project/domain/entity/settlement"
	"pc-shop-final-project/internal/delivery/http/models"
)

func DataSettlementDbToEntity(dataDTO settlement2.DTOSettlement) (*settlement2.Settlement, error) {
	settlement, err := settlement2.NewSettlement(dataDTO)
	if err != nil {
		return nil, err
	}

	return settlement, nil
}

func SettlementEntityToModel(set *settlement2.Settlement) *models.ModelSettlement {
	return &models.ModelSettlement{
		CODE_TRANSACTION: set.GetValueCodeSett(),
		TOTAL_PRICE:      set.GetValueTotalPriceSett(),
	}
}

func SettlementEntityToDbqStruct(settlement *settlement2.Settlement) []interface{} {
	dbqStruct := dbq.Struct(SettlementEntityToModel(settlement))
	return dbqStruct
}

func SettlementModelToEntity(model *models.ModelSettlement) (*settlement2.Settlement, error) {
	settlement, err := settlement2.NewSettlement(settlement2.DTOSettlement{
		Code:       model.CODE_TRANSACTION,
		TotalPrice: model.TOTAL_PRICE,
	})

	if err != nil {
		return nil, err
	}

	return settlement, nil
}
