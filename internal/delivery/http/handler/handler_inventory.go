package handler

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity/inventory"
	_interface "pc-shop-final-project/domain/repository"
	"pc-shop-final-project/internal/repository/mysql"
)

var (
	repoInventoryMysql = mysql.NewInventoryMysql(mysqlConnection)
	HandlerInventory   = NewInventoryHandler(repoInventoryMysql)
)

type InventoryInteractor struct {
	repository _interface.InterfaceInventory
}

func NewInventoryHandler(Repo _interface.InterfaceInventory) *InventoryInteractor {
	return &InventoryInteractor{
		repository: Repo,
	}
}

func CreateInventory(ctx context.Context, inven *inventory.Inventory) {
	err := HandlerInventory.repository.CreateInventory(ctx, inven)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadInventory(ctx context.Context) []*inventory.Inventory {
	Inventories, err := HandlerInventory.repository.ReadInventory(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return Inventories
}

func DeleteInventory(ctx context.Context, code int) {
	err := HandlerInventory.repository.DeleteInventory(ctx, code)
	if err != nil {
		fmt.Println(err)
	}
}
