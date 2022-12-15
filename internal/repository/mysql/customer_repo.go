package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"pc-shop-final-project/domain/entity/customer"
	"pc-shop-final-project/internal/delivery/http/models"
	"pc-shop-final-project/internal/repository/mysql/mapper"
	"time"
)

type CustomerMysqlInteractor struct {
	db *sql.DB
}

func NewCustomerMysql(db *sql.DB) *CustomerMysqlInteractor {
	return &CustomerMysqlInteractor{db: db}
}

// CreateCustomer implements _interface.InterfaceCustomer
func (cst *CustomerMysqlInteractor) InsertDataCustomer(ctx context.Context, cust *customer.Customer) error {
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
func (cst *CustomerMysqlInteractor) GetListCustomer(ctx context.Context) ([]*customer.Customer, error) {
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

// ReadUser by ID implements _interface.InterfaceCustomer
func (cst *CustomerMysqlInteractor) GetCustomerById(ctx context.Context, id string) (*customer.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	queryCustomer := fmt.Sprintf("SELECT * FROM %s WHERE UNIQ_ID = ?", models.GetCustomerTableName())

	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.ModelCustomer{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	resultCustomer, err := dbq.Q(ctx, cst.db, queryCustomer, opts, id)

	if err != nil {
		return nil, err
	}

	if resultCustomer == nil {
		return nil, errors.New("CUSTOMER TIDAK DITEMUKAN")
	}

	customer, errMap := mapper.CustomerModelToEntity(resultCustomer.(*models.ModelCustomer))

	if errMap != nil {
		return nil, errMap
	}

	return customer, nil
}

//func (usr *UserMysqlInteractor) UpdateUserById(ctx context.Context, dataUser *user2.User, idUser string) error {
//	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
//	defer cancel()
//
//	query := fmt.Sprintf("UPDATE %s SET ID_USER = ?, NAME = ?, OUTLET_CODE = ?, STATUS = ? "+
//		"WHERE ID_USER = '%s'", models.GetUserTableName(), idUser)
//
//	_, err := dbq.E(ctx, usr.db, query, nil, mapper.UserEntityToDbqStruct(dataUser))
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
