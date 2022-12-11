package custRevenue

import (
	"pc-shop-final-project/domain/entity/settlement"
	"strconv"
	"time"
)

type CustomerRevenue struct {
	idCustomer      string
	revenue         int
	dateTransaction time.Time
}

type DTOCustomerRevenue struct {
	IdCustomer      string
	Revenue         int
	DateTransaction time.Time
}

func NewCustomerRevenue(idCust string, settle []*settlement.Settlement) ([]*CustomerRevenue, []*CustomerRevenue) {
	var (
		timeTransactionDate time.Time
		revenue             int
	)

	custRev := make([]*CustomerRevenue, 0)

	for _, set := range settle {
		if set.GetValueCustomerSett().GetValueUniqIdCust() == idCust {
			settleCode := []rune(set.GetValueCodeSett())
			dateSettle := string(settleCode[8]) + string(settleCode[9]) + string(settleCode[10]) + string(settleCode[11]) + string(settleCode[12]) + string(settleCode[13])
			timeTransactionDate, _ = time.Parse("020106", dateSettle)
			revenue = set.GetValueTotalPriceSett()
		}
	}

	cr := &CustomerRevenue{
		idCustomer:      idCust,
		revenue:         revenue,
		dateTransaction: timeTransactionDate,
	}

	exp, _ := strconv.Atoi(timeTransactionDate.AddDate(0, 1, 0).Format("20060102"))
	now, _ := strconv.Atoi(time.Now().Format("20060102"))

	if exp < now {
		custRev = append(custRev, cr)
	}

	return custRev, nil
}

func (cr *CustomerRevenue) GetValueIdCustCR() string {
	return cr.idCustomer
}

func (cr *CustomerRevenue) GetValueRevenueCR() int {
	return cr.revenue
}

func (cr *CustomerRevenue) GetValueDateTransactionCR() string {
	return cr.dateTransaction.Format("2006-01-02 15:04:05.")
}
