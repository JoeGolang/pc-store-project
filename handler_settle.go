package handler

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity"
	_interface "pc-shop-final-project/domain/repository"
	sqlConn "pc-shop-final-project/internal/config/database/mysql"
	"pc-shop-final-project/internal/repository/mysql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlConnection = sqlConn.InitMysqlDB()
	repoMysql       = mysql.NewSettleMysql(mysqlConnection)
	HandlerSettle   = NewSettleHandler(repoMysql)
)

type SettleInteractor struct {
	repository _interface.InterfaceSettlement
}

func NewSettleHandler(Repo _interface.InterfaceSettlement) *SettleInteractor {
	return &SettleInteractor{
		repository: Repo,
	}
}

func CreateSettle(ctx context.Context, idUser int, idCustomer int, idCoupon int, settle *entity.Settlement) {
	err := HandlerSettle.repository.CreateSettle(ctx, idUser, idCustomer, idCoupon, settle)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadSettle(ctx context.Context) []*entity.Settlement {
	settlements, err := HandlerSettle.repository.ReadSettle(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return settlements
}

func UpdateSettle(ctx context.Context, code string) {
	err := HandlerSettle.repository.UpdateSettle(ctx, code)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteSettle(ctx context.Context, code string) {
	err := HandlerSettle.repository.DeleteSettle(ctx, code)
	if err != nil {
		fmt.Println(err)
	}
}
