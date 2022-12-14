package repository

import (
	"context"
	"pc-shop-final-project/domain/entity/inventory"
)

type InterfaceInventory interface {
	GetListInventory(ctx context.Context) ([]*inventory.Inventory, error)
	//CreateInventory(ctx context.Context) ([]*inventory.Inventory, error)
	//ReadInventory(ctx context.Context) ([]*inventory.Inventory, error)
	//DeleteInventory(ctx context.Context, id int) error
}
