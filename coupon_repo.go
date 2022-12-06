package mysql

import (
	"context"
	"database/sql"
	"pc-shop-final-project/domain/entity"
	_interface "pc-shop-final-project/domain/repository"
	"time"
)

type CouponMysqlInteractor struct {
	db *sql.DB
}

func NewCouponMysql(db *sql.DB) _interface.InterfaceCoupon {
	return &CouponMysqlInteractor{db: db}
}

// CreateCoupon implements _interface.InterfaceCoupon
func (cpn *CouponMysqlInteractor) CreateCoupon(ctx context.Context, coupon *entity.Coupon, uniqcoupon []*entity.UniqCoupon) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	setAutoInc := "ALTER TABLE coupon AUTO_INCREMENT=1"
	_, errMysql = cpn.db.Exec(setAutoInc)

	if errMysql != nil {
		return errMysql
	}

	insertQueryCoupon := "INSERT INTO coupon(ID_COUPON, ID_CUSTOMER, GENERATE_DATE, REVENUE, ACTIVE_DATE_STATUS) VALUES (?,?,?,?,?)"

	_, errMysql = cpn.db.Exec(insertQueryCoupon, coupon.GetValueIdCpn(), coupon.GetValueIdCustomerCpn(), coupon.GetValueGenDateCpn(), coupon.GetValueRevenueCpn(), coupon.GetValueActiveCpn())

	if errMysql != nil {
		return errMysql
	}

	for _, uniqCpn := range coupon.GetValueUniqIdCpn() {
		insertQueryUC := "INSERT INTO code_coupon(ID_COUPON, UNIQUE_ID, ACTIVE_USE_STATUS) VALUES (?,?,?)"

		_, errMysql = cpn.db.Exec(insertQueryUC, uniqCpn.GetValueIdCoupon(), uniqCpn.GetValueUniqIdCoupon(), uniqCpn.GetValueStatusCoupon())
	}

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// DeleteCoupon implements _interface.InterfaceCoupon
func (cpn *CouponMysqlInteractor) DeleteCoupon(ctx context.Context, id int) error {
	var (
		errMysqlA error
		errMysqlB error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteQuery := "DELETE FROM coupon WHERE ID_COUPON = ?"

	_, errMysqlA = cpn.db.Exec(deleteQuery, id)

	if errMysqlA != nil {
		return errMysqlA
	}

	deleteQueryCode := "DELETE FROM code_coupon WHERE ID_COUPON = ?"

	_, errMysqlB = cpn.db.Exec(deleteQueryCode, id)

	if errMysqlB != nil {
		return errMysqlB
	}

	return nil
}

// ReadCoupon implements _interface.InterfaceCoupon
func (cpn *CouponMysqlInteractor) ReadCoupon(ctx context.Context) ([]*entity.Coupon, []*entity.UniqCoupon, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQueryUniq := "SELECT ID_COUPON, UNIQUE_ID, ACTIVE_USE_STATUS FROM code_coupon"
	rows, errMysql := cpn.db.QueryContext(ctx, sqlQueryUniq)
	if errMysql != nil {
		return nil, nil, errMysql
	}

	listUniqCoupon := make([]*entity.UniqCoupon, 0)
	for rows.Next() {
		var (
			ID_COUPON         int
			UNIQ_ID           string
			ACTIVE_USE_STATUS bool
		)

		errScan := rows.Scan(&ID_COUPON, &UNIQ_ID, &ACTIVE_USE_STATUS)
		if errScan != nil {
			return nil, nil, errScan
		}

		uniqCoupon := entity.FetchUniqCoupon(&entity.DTOUniqCoupon{
			Id:     ID_COUPON,
			UniqId: UNIQ_ID,
			Status: ACTIVE_USE_STATUS,
		})

		listUniqCoupon = append(listUniqCoupon, uniqCoupon)
	}

	sqlQuery := "SELECT ID_COUPON, ID_CUSTOMER, GENERATE_DATE, REVENUE, ACTIVE_DATE_STATUS FROM coupon"
	rows, errMysql = cpn.db.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, nil, errMysql
	}

	listCoupon := make([]*entity.Coupon, 0)
	for rows.Next() {
		var (
			ID_COUPON          int
			ID_CUSTOMER        string
			GENERATE_DATE      string
			REVENUE            int
			ACTIVE_DATE_STATUS bool
		)

		errScan := rows.Scan(&ID_COUPON, &ID_CUSTOMER, &GENERATE_DATE, &REVENUE, &ACTIVE_DATE_STATUS)
		if errScan != nil {
			return nil, nil, errScan
		}

		listMatchUniqCoupon := make([]entity.UniqCoupon, 0)

		for _, matchCoupon := range listUniqCoupon {
			if matchCoupon.GetValueIdCoupon() == ID_COUPON {
				listMatchUniqCoupon = append(listMatchUniqCoupon, *matchCoupon)
			}
		}

		coupon := entity.FetchCoupon(&entity.DTOCoupon{
			Id:         ID_COUPON,
			IdCustomer: ID_CUSTOMER,
			UniqCoupon: listMatchUniqCoupon,
			GenDate:    GENERATE_DATE,
			Revenue:    REVENUE,
			Active:     ACTIVE_DATE_STATUS,
		})

		listCoupon = append(listCoupon, coupon)
	}
	defer rows.Close()

	return listCoupon, listUniqCoupon, nil
}

// UpdateCoupon implements _interface.InterfaceCoupon
func (cpn *CouponMysqlInteractor) UpdateCoupon(ctx context.Context, uniqId string) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	updateQuery := "UPDATE code_coupon SET ACTIVE_USE_STATUS = false WHERE UNIQUE_ID = ?"

	_, errMysql = cpn.db.Exec(updateQuery, uniqId)

	if errMysql != nil {
		return errMysql
	}
	return nil
}
