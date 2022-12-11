package settle_repo_test

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity/settlement"
	"pc-shop-final-project/internal/delivery/http/handler"
	"testing"
)

func TestCreateSettleMysql(t *testing.T) {
	var (
		ctx = context.Background()
	)

	settlement, err := settlement.NewSettlement(settlement.DTOSettlement{
		TotalPrice: 32000000,
	})

	if err != nil {
		panic(err)
	}

	handler.CreateSettle(ctx, 1, 1, 1, settlement)
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
