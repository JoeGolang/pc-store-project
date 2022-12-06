package entity

import (
	"strconv"
	"time"
)

type CustomerRevenue struct {
	idCustomer      string
	revenue         int
	dateTransaction time.Time
}

func NewCustomerRevenue(idCust string, settle []*Settlement) []*CustomerRevenue {
	var (
		timeTransactionDate time.Time
		revenue             int
	)

	custRev := make([]*CustomerRevenue, 0)

	for _, set := range settle {
		if set.customer.uniqId == idCust {
			settleCode := []rune(set.code)
			dateSettle := string(settleCode[8]) + string(settleCode[9]) + string(settleCode[10]) + string(settleCode[11]) + string(settleCode[12]) + string(settleCode[13])
			timeTransactionDate, _ = time.Parse("020106", dateSettle)
			revenue = set.totalPrice
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

	return custRev
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
