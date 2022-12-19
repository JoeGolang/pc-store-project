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

func CreateSettle(ctx context.Context, idUser int, idCustomer string, idCoupon int, settle *entity.Settlement) {
	err := HandlerSettle.repository.CreateSettle(ctx, idUser, idCustomer, idCoupon, settle)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadSettle(ctx context.Context) []*entity.Settlement {
	var (
		x entity.User
		y entity.Customer
		z entity.Coupon
		i int
	)
	usr, cst, cpn, settlements, err := HandlerSettle.repository.ReadSettle(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	setts := make([]*entity.Settlement, 0)
	users := ReadUser(ctx)
	custs := ReadCustomer(ctx)
	coups, _ := ReadCoupon(ctx)
	setIt := ReadSettleItem(ctx)

	for _, sett := range settlements {
		for _, user := range users {
			if user.GetValueIdUsr() == usr[i] {
				x = *user
			}
		}
		for _, cust := range custs {
			if cust.GetValueUniqIdCust() == cst[i] {
				y = *cust
			}
		}
		for _, coup := range coups {
			if coup.GetValueIdCpn() == cpn[i] {
				z = *coup
			}
		}
		dataItem := make([]*entity.SettlePurchase, 0)
		for _, SI := range setIt {
			if SI.GetValueCodeSettlement() == sett.GetValueCodeSett() {
				dataItem = append(dataItem, SI)
			}
		}

		settle, err := entity.NewSettlement(entity.DTOSettlement{
			Id:         sett.GetValueIdSett(),
			Code:       sett.GetValueCodeSett(),
			User:       x,
			Customer:   y,
			Product:    dataItem,
			Coupon:     z,
			TotalPrice: sett.GetValueTotalPriceSett(),
			StatusTrns: sett.GetValueStatusTrns(),
		})
		if err != nil {
			fmt.Println(err)
			return nil
		}
		setts = append(setts, settle)
		i++
	}

	return setts
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
