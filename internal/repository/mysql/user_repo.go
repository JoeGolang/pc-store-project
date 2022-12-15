package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	user2 "pc-shop-final-project/domain/entity/user"
	"pc-shop-final-project/internal/delivery/http/models"
	"pc-shop-final-project/internal/repository/mysql/mapper"
	"time"
)

type UserMysqlInteractor struct {
	db *sql.DB
}

func (usr *UserMysqlInteractor) UpdateUserByKode(ctx context.Context, dataUser *user2.User, kodeUser string) error {
	//TODO implement me
	panic("implement me")
}

// DeleteUser implements _interface.InterfaceUser
func (usr *UserMysqlInteractor) DeleteUserById(ctx context.Context, id string) error {
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

func NewUserMysql(db *sql.DB) *UserMysqlInteractor {
	return &UserMysqlInteractor{db: db}
}

func (usr *UserMysqlInteractor) InsertDataUser(ctx context.Context, user *user2.User) error {
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

// ReadUser implements _interface.InterfaceUser
func (usr *UserMysqlInteractor) GetListUser(ctx context.Context) ([]*user2.User, error) {
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

	listUser := make([]*user2.User, 0)
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

		user, errFetch := user2.NewUser(user2.DTOUser{
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

// ReadUser by ID implements _interface.InterfaceUser
func (usr *UserMysqlInteractor) GetUserById(ctx context.Context, id string) (*user2.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	queryUser := fmt.Sprintf("SELECT * FROM %s WHERE ID_USER = ?", models.GetUserTableName())

	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.ModelUser{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	resultUser, err := dbq.Q(ctx, usr.db, queryUser, opts, id)

	if err != nil {
		return nil, err
	}

	if resultUser == nil {
		return nil, errors.New("USER TIDAK DITEMUKAN")
	}

	user, errMap := mapper.UserModelToEntity(resultUser.(*models.ModelUser))

	if errMap != nil {
		return nil, errMap
	}

	return user, nil
}
