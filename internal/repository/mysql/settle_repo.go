package mysql

import (
	"context"
	"database/sql"
	"pc-shop-final-project/domain/entity/coupon"
	"pc-shop-final-project/domain/entity/customer"
	"pc-shop-final-project/domain/entity/inventory"
	"pc-shop-final-project/domain/entity/settlement"
	user2 "pc-shop-final-project/domain/entity/user"
	_interface "pc-shop-final-project/domain/repository"
	"time"
)

type SettleMysqlInteractor struct {
	db *sql.DB
}

func NewSettleMysql(db *sql.DB) _interface.InterfaceSettlement {
	return &SettleMysqlInteractor{db: db}
}

// CreateSettle implements _interface.InterfaceSettlement
func (set *SettleMysqlInteractor) CreateSettle(ctx context.Context, idUser int, idCustomer int, idCoupon int, settle *settlement.Settlement) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	setAutoInc := "ALTER TABLE settlement AUTO_INCREMENT=1"
	_, errMysql = set.db.Exec(setAutoInc)

	if errMysql != nil {
		return errMysql
	}

	insertQuery := "INSERT INTO settlement(USER, CUSTOMER, COUPON_ID, CODE_TRANSACTION, TOTAL_PRICE, STATUS_TRANSACTION) VALUES (?,?,?,?,?,?)"

	_, errMysql = set.db.Exec(insertQuery, idUser, idCustomer, idCoupon, settle.GetValueCodeSett(), settle.GetValueTotalPriceSett(), settle.GetValueStatusTrns())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// DeleteSettle implements _interface.InterfaceSettlement
func (set *SettleMysqlInteractor) DeleteSettle(ctx context.Context, code string) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteQuery := "DELETE FROM settlement WHERE CODE_TRANSACTION = ?"

	_, errMysql = set.db.Exec(deleteQuery, code)

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// ReadSettle implements _interface.InterfaceSettlement
func (set *SettleMysqlInteractor) ReadSettle(ctx context.Context) ([]*settlement.Settlement, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT ID, USER, CUSTOMER, COUPON_ID, CODE_TRANSACTION, TOTAL_PRICE, STATUS_TRANSACTION FROM settlement"
	rows, errMysql := set.db.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, errMysql
	}

	listSettlement := make([]*settlement.Settlement, 0)
	for rows.Next() {
		var (
			ID                 int
			USER               int
			CUSTOMER           int
			COUPON_ID          int
			CODE_TRANSACTION   string
			TOTAL_PRICE        int
			STATUS_TRANSACTION bool
		)

		errScan := rows.Scan(&ID, &USER, &CUSTOMER, &COUPON_ID, &CODE_TRANSACTION, &TOTAL_PRICE, &STATUS_TRANSACTION)
		if errScan != nil {
			return nil, errScan
		}

		set, errFetch := settlement.NewSettlement(settlement.DTOSettlement{
			Id:         ID,
			Code:       CODE_TRANSACTION,
			User:       user2.User{},
			Customer:   customer.Customer{},
			Product:    []inventory.DTOInventory{},
			Coupon:     coupon.Coupon{},
			TotalPrice: TOTAL_PRICE,
			StatusTrns: STATUS_TRANSACTION,
		})

		if errFetch != nil {
			return nil, errFetch
		}

		listSettlement = append(listSettlement, set)
	}
	defer rows.Close()

	return listSettlement, nil
}

// UpdateSettle implements _interface.InterfaceSettlement
func (set *SettleMysqlInteractor) UpdateSettle(ctx context.Context, code string) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	updateQuery := "UPDATE settlement SET STATUS_TRANSACTION = true WHERE CODE_TRANSACTION = ?"

	_, errMysql = set.db.Exec(updateQuery, code)

	if errMysql != nil {
		return errMysql
	}
	return nil
}
