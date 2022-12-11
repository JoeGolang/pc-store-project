package repository

import (
	"context"
	"pc-shop-final-project/domain/entity/inventory"
)

type InterfaceInventory interface {
	CreateInventory(ctx context.Context, inv *inventory.Inventory) error
	ReadInventory(ctx context.Context) ([]*inventory.Inventory, error)
	DeleteInventory(ctx context.Context, id int) error
}
