package mysql

import (
	"context"
	"database/sql"
	"pc-shop-final-project/domain/entity"
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
func (set *SettleMysqlInteractor) CreateSettle(ctx context.Context, idUser int, idCustomer string, idCoupon int, settle *entity.Settlement) error {
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
func (set *SettleMysqlInteractor) ReadSettle(ctx context.Context) ([]int, []string, []int, []*entity.Settlement, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT ID, USER, CUSTOMER, COUPON_ID, CODE_TRANSACTION, TOTAL_PRICE, STATUS_TRANSACTION FROM settlement"
	rows, errMysql := set.db.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, nil, nil, nil, errMysql
	}

	listSettlement := make([]*entity.Settlement, 0)
	users := make([]int, 0)
	custs := make([]string, 0)
	coups := make([]int, 0)
	for rows.Next() {
		var (
			ID                 int
			USER               int
			CUSTOMER           string
			COUPON_ID          int
			CODE_TRANSACTION   string
			TOTAL_PRICE        int
			STATUS_TRANSACTION bool
		)

		errScan := rows.Scan(&ID, &USER, &CUSTOMER, &COUPON_ID, &CODE_TRANSACTION, &TOTAL_PRICE, &STATUS_TRANSACTION)
		if errScan != nil {
			return nil, nil, nil, nil, errScan
		}

		user, cust, coup, set := entity.FetchSettlement(USER, CUSTOMER, COUPON_ID, entity.DTOSettlement{
			Id:         ID,
			Code:       CODE_TRANSACTION,
			TotalPrice: TOTAL_PRICE,
			StatusTrns: STATUS_TRANSACTION,
		})

		listSettlement = append(listSettlement, set)
		users = append(users, user)
		custs = append(custs, cust)
		coups = append(coups, coup)
	}
	defer rows.Close()

	return users, custs, coups, listSettlement, nil
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
