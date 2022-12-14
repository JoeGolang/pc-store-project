package handler

import (
	"context"
	_interface "pc-shop-final-project/domain/repository"
)

//var (
//	repoInventoryMysql = mysql.NewInventoryMysql(mysqlConnection)
//	HandlerInventory   = NewInventoryHandler(repoInventoryMysql)
//)

type InventoryHandler struct {
	ctx           context.Context
	repoInventory _interface.InterfaceInventory
}

//type InventoryInteractor struct {
//	repository _interface.InterfaceInventory
//}

func NewInventoryHandler(ctx context.Context, RepoInv _interface.InterfaceInventory) *InventoryHandler {
	return &InventoryHandler{
		ctx:           ctx,
		repoInventory: RepoInv,
	}
}

//func CreateInventory(ctx context.Context, inven *inventory.Inventory) {
//	err := HandlerInventory.repository.CreateInventory(ctx, inven)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//func ReadInventory(ctx context.Context) []*inventory.Inventory {
//	Inventories, err := HandlerInventory.repository.ReadInventory(ctx)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	return Inventories
//}

//func DeleteInventory(ctx context.Context, code int) {
//	err := HandlerInventory.repository.DeleteInventory(ctx, code)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
