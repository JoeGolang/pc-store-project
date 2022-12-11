package mysql

import (
	"context"
	"database/sql"
	"pc-shop-final-project/domain/entity/customer"
	_interface "pc-shop-final-project/domain/repository"
	"time"
)

type CustomerMysqlInteractor struct {
	db *sql.DB
}

func NewCustomerMysql(db *sql.DB) _interface.InterfaceCustomer {
	return &CustomerMysqlInteractor{db: db}
}

// CreateCustomer implements _interface.InterfaceCustomer
func (cst *CustomerMysqlInteractor) CreateCustomer(ctx context.Context, cust *customer.Customer) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	setAutoInc := "ALTER TABLE customer AUTO_INCREMENT=1"
	_, errMysql = cst.db.Exec(setAutoInc)

	if errMysql != nil {
		return errMysql
	}

	insertQuery := "INSERT INTO customer(UNIQ_ID, NAME, JOIN_DATE) VALUES (?,?,?)"

	_, errMysql = cst.db.Exec(insertQuery, cust.GetValueUniqIdCust(), cust.GetValueNameCust(), cust.GetValueJoinDateCust())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// DeleteCustomer implements _interface.InterfaceCustomer
func (cst *CustomerMysqlInteractor) DeleteCustomer(ctx context.Context, uniqId string) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteQuery := "DELETE FROM customer WHERE UNIQ_ID = ?"

	_, errMysql = cst.db.Exec(deleteQuery, uniqId)

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// ReadCustomer implements _interface.InterfaceCustomer
func (cst *CustomerMysqlInteractor) ReadCustomer(ctx context.Context) ([]*customer.Customer, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT UNIQ_ID, NAME, JOIN_DATE FROM customer"
	rows, errMysql := cst.db.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, errMysql
	}

	listCustomer := make([]*customer.Customer, 0)
	for rows.Next() {
		var (
			UNIQ_ID   string
			NAME      string
			JOIN_DATE string
		)

		errScan := rows.Scan(&UNIQ_ID, &NAME, &JOIN_DATE)
		if errScan != nil {
			return nil, errScan
		}

		customer, errFetch := customer.NewCustomer(customer.DTOCustomer{
			UniqId:   UNIQ_ID,
			Name:     NAME,
			JoinDate: JOIN_DATE,
		}, "manggadua")

		if errFetch != nil {
			return nil, errFetch
		}

		listCustomer = append(listCustomer, customer)
	}
	defer rows.Close()

	return listCustomer, nil
}
