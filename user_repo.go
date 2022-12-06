package mysql

import (
	"context"
	"database/sql"
	"pc-shop-final-project/domain/entity"
	_interface "pc-shop-final-project/domain/repository"
	"time"
)

type UserMysqlInteractor struct {
	db *sql.DB
}

func NewUserMysql(db *sql.DB) _interface.InterfaceUser {
	return &UserMysqlInteractor{db: db}
}

// CreateUser implements _interface.InterfaceUser
func (usr *UserMysqlInteractor) CreateUser(ctx context.Context, user *entity.User) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	setAutoInc := "ALTER TABLE user AUTO_INCREMENT=1"
	_, errMysql = usr.db.Exec(setAutoInc)

	if errMysql != nil {
		return errMysql
	}

	insertQuery := "INSERT INTO user(ID_USER, NAME, OUTLET_CODE, STATUS) VALUES (?,?,?,?)"

	_, errMysql = usr.db.Exec(insertQuery, user.GetValueIdUsr(), user.GetValueNameUsr(), user.GetValueOutletCodeUsr(), user.GetValueStatusUsr())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// DeleteUser implements _interface.InterfaceUser
func (usr *UserMysqlInteractor) DeleteUser(ctx context.Context, id int) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteQuery := "DELETE FROM user WHERE ID_USER = ?"

	_, errMysql = usr.db.Exec(deleteQuery, id)

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// ReadUser implements _interface.InterfaceUser
func (usr *UserMysqlInteractor) ReadUser(ctx context.Context) ([]*entity.User, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT ID_USER, NAME, OUTLET_CODE, STATUS FROM user"
	rows, errMysql := usr.db.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, errMysql
	}

	listUser := make([]*entity.User, 0)
	for rows.Next() {
		var (
			ID_USER     int
			NAME        string
			OUTLET_CODE string
			STATUS      string
		)

		errScan := rows.Scan(&ID_USER, &NAME, &OUTLET_CODE, &STATUS)
		if errScan != nil {
			return nil, errScan
		}

		user, errFetch := entity.NewUser(&entity.DTOUser{
			Id:         ID_USER,
			Name:       NAME,
			OutletCode: OUTLET_CODE,
			Status:     STATUS,
		})

		if errFetch != nil {
			return nil, errFetch
		}

		listUser = append(listUser, user)
	}
	defer rows.Close()

	return listUser, nil
}
