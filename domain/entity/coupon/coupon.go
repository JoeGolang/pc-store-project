package entity

import (
	"strconv"
	"strings"
	"time"
)

type Coupon struct {
	id              int
	idCustomer      string
	uniqCoupon      []UniqCoupon
	generateDate    time.Time
	prevShopRevenue int
	active          bool
}

type UniqCoupon struct {
	idCoupon int
	uniqId   string
	status   bool
}

type DTOCoupon struct {
	Id         int
	IdCustomer string
	UniqCoupon []UniqCoupon
	GenDate    string
	Revenue    int
	Active     bool
}

type DTOUniqCoupon struct {
	Id     int
	UniqId string
	Status bool
}

func AutoGenerateCoupon(revenue int, idCust string) *Coupon {
	var (
		uniqCoupon *UniqCoupon
	)

	charsName := []rune(idCust)
	charName := string(charsName[0]) + string(charsName[1]) + string(charsName[2])
	upCharName := strings.ToUpper(charName)

	generateTime := time.Now()

	coupons := make([]UniqCoupon, 0)

	if revenue > 6000000 {
		newId := "BASIC-" + upCharName + "1" + generateTime.Format("020106150405")
		uniqCoupon = &UniqCoupon{
			uniqId: newId,
			status: true,
		}
		coupons = append(coupons, *uniqCoupon)
	}
	if revenue > 13000000 {
		newId := "PREMI-" + upCharName + "2" + generateTime.Format("020106150405")
		uniqCoupon = &UniqCoupon{
			uniqId: newId,
			status: true,
		}
		coupons = append(coupons, *uniqCoupon)
	}
	if revenue > 25000000 {
		newId := "ULTI-" + upCharName + "3" + generateTime.Format("020106150405")
		uniqCoupon = &UniqCoupon{
			uniqId: newId,
			status: true,
		}
		coupons = append(coupons, *uniqCoupon)
	}
	if revenue <= 6000000 {
		coupons = nil
	}

	return &Coupon{
		idCustomer:      idCust,
		uniqCoupon:      coupons,
		generateDate:    generateTime,
		prevShopRevenue: revenue,
		active:          true,
	}
}

func (cpn *Coupon) GetValueIdCpn() int {
	return cpn.id
}

func (cpn *Coupon) GetValueIdCustomerCpn() string {
	return cpn.idCustomer
}

func (cpn *Coupon) GetValueUniqIdCpn() []UniqCoupon {
	return cpn.uniqCoupon
}

func (cpn *Coupon) GetValueGenDateCpn() string {
	return cpn.generateDate.Format("2006-01-02 15:04:05.")
}

func (cpn *Coupon) GetValueRevenueCpn() int {
	return cpn.prevShopRevenue
}

func (cpn *Coupon) GetValueActiveCpn() bool {
	exp, _ := strconv.Atoi(cpn.generateDate.AddDate(0, 1, 0).Format("20060102150405"))
	now, _ := strconv.Atoi(time.Now().Format("20060102150405"))
	if exp < now {
		cpn.active = true
	} else {
		cpn.active = false
	}
	return cpn.active
}

func (cpn *Coupon) SetUseCouponBasic(uniqCode string) *Coupon {
	for _, coupon := range cpn.uniqCoupon {
		if coupon.uniqId == uniqCode {
			coupon.status = false
		}
	}
	return &Coupon{}
}

func (cpn *Coupon) SetUseCouponPremi(uniqCode string) *Coupon {
	for _, coupon := range cpn.uniqCoupon {
		if coupon.uniqId == uniqCode {
			coupon.status = false
		}
	}
	return &Coupon{}
}

func (cpn *Coupon) SetUseCouponUlti(uniqCode string) *Coupon {
	for _, coupon := range cpn.uniqCoupon {
		if coupon.uniqId == uniqCode {
			coupon.status = false
		}
	}
	return &Coupon{}
}

func (UC *UniqCoupon) GetValueIdCoupon() int {
	return UC.idCoupon
}

func (UC *UniqCoupon) GetValueUniqIdCoupon() string {
	return UC.uniqId
}

func (UC *UniqCoupon) GetValueStatusCoupon() bool {
	return UC.status
}

func FetchCoupon(cpn *DTOCoupon) *Coupon {
	timeGenerateDate, _ := time.Parse("2006-01-02 15:04:05.", cpn.GenDate)

	return &Coupon{
		id:              cpn.Id,
		idCustomer:      cpn.IdCustomer,
		uniqCoupon:      []UniqCoupon{},
		generateDate:    timeGenerateDate,
		prevShopRevenue: cpn.Revenue,
		active:          cpn.Active,
	}
}

func FetchUniqCoupon(UC *DTOUniqCoupon) *UniqCoupon {
	return &UniqCoupon{
		idCoupon: UC.Id,
		uniqId:   UC.UniqId,
		status:   UC.Status,
	}
}
