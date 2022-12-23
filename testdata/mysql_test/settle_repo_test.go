package mysql_test_test

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity"
	handler "pc-shop-final-project/internal/repository/controller"
	"testing"
)

func TestCreateSettleMysql(t *testing.T) {
	var (
		ctx  = context.Background()
		data = entity.User{}
	)

	users := handler.ReadUser(ctx)
	for _, user := range users {
		if user.GetValueIdUsr() == 1002 {
			data = *user
		}
	}

	settlement, err := entity.NewSettlement(entity.DTOSettlement{
		Id:         0,
		Code:       "",
		User:       data,
		Customer:   entity.Customer{},
		Product:    nil,
		Coupon:     entity.Coupon{},
		TotalPrice: 32000000,
		StatusTrns: true,
	})

	if err != nil {
		panic(err)
	}

	handler.CreateSettle(ctx, 1, "BB", 2, settlement)
}

func TestReadSettleMysql(t *testing.T) {
	var (
		ctx = context.Background()
	)

	settlements := handler.ReadSettle(ctx)

	for _, settlement := range settlements {
		fmt.Println(settlement)
	}
}

func TestDeleteSettleMysql(t *testing.T) {
	var (
		ctx  = context.Background()
		code string
	)

	settlements := handler.ReadSettle(ctx)

	for _, settlement := range settlements {
		code = settlement.GetValueCodeSett()
	}

	handler.DeleteSettle(ctx, code)
}

func TestUpdateMysql(t *testing.T) {
	var (
		ctx  = context.Background()
		code string
	)

	settlements := handler.ReadSettle(ctx)

	for _, settlement := range settlements {
		code = settlement.GetValueCodeSett()
	}

	handler.UpdateSettle(ctx, code)
}
