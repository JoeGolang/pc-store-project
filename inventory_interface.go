package _interface

import (
	"context"
	"pc-shop-final-project/domain/entity"
)

type InterfaceInventory interface {
	CreateInventory(ctx context.Context, inv *entity.Inventory) error
	ReadInventory(ctx context.Context) ([]*entity.Inventory, error)
	DeleteInventory(ctx context.Context, id int) error
}
