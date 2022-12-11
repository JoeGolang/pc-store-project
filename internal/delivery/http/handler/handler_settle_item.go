package handler

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity/settlement"
	_interface "pc-shop-final-project/domain/repository"
	"pc-shop-final-project/internal/repository/mysql"
)

var (
	repoSettleItemMysql = mysql.NewSettlePurchaseMysql(mysqlConnection)
	HandlerSettleItem   = NewSettleItemHandler(repoSettleItemMysql)
)

type SettleItemInteractor struct {
	repository _interface.InterfaceSettlementItem
}

func NewSettleItemHandler(Repo _interface.InterfaceSettlementItem) *SettleItemInteractor {
	return &SettleItemInteractor{
		repository: Repo,
	}
}

func CreateSettleItem(ctx context.Context, items []*settlement.SettlePurchase) {
	err := HandlerSettleItem.repository.CreatePurchase(ctx, items)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadSettleItem(ctx context.Context) []*settlement.SettlePurchase {
	SI, err := HandlerSettleItem.repository.ReadPurchase(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return SI
}

func DeleteSettleitem(ctx context.Context, codeItem string) {
	err := HandlerSettleItem.repository.DeletePurchase(ctx, codeItem)
	if err != nil {
		fmt.Println(err)
	}
}
