package mysql

import (
	"context"
	"database/sql"
	"pc-shop-final-project/domain/entity/inventory"
	"time"
)

type InventoryMysqlInteractor struct {
	db *sql.DB
}

func NewInventoryMysql(db *sql.DB) *InventoryMysqlInteractor {
	return &InventoryMysqlInteractor{
		db: db,
	}
}

// CreateInventory implements _interface.InterfaceInventory
func (inven *InventoryMysqlInteractor) CreateInventory(ctx context.Context, inv *inventory.Inventory) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	setAutoInc := "ALTER TABLE inventory AUTO_INCREMENT=1"
	_, errMysql = inven.db.Exec(setAutoInc)

	if errMysql != nil {
		return errMysql
	}

	insertQuery := "INSERT INTO inventory(ID_PRODUCT, PRODUCT_NAME, BRAND, PRICE, CATEGORY) VALUES (?,?,?,?,?)"

	_, errMysql = inven.db.Exec(insertQuery, inv.GetValueIdInv(), inv.GetValueProductNameInv(), inv.GetValueBrandInv(), inv.GetvaluePriceInv(), inv.GetValueCategoryInv())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// DeleteInventory implements _interface.InterfaceInventory
func (inven *InventoryMysqlInteractor) DeleteInventory(ctx context.Context, id int) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteQuery := "DELETE FROM inventory WHERE ID_PRODUCT = ?"

	_, errMysql = inven.db.Exec(deleteQuery, id)

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// ReadInventory implements _interface.InterfaceInventory
func (inven *InventoryMysqlInteractor) ReadInventory(ctx context.Context) ([]*inventory.Inventory, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT ID_PRODUCT, PRODUCT_NAME, BRAND, PRICE, CATEGORY FROM inventory"
	rows, errMysql := inven.db.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, errMysql
	}

	listInventory := make([]*inventory.Inventory, 0)
	for rows.Next() {
		var (
			ID_PRODUCT   int
			PRODUCT_NAME string
			BRAND        string
			PRICE        int
			CATEGORY     string
		)

		errScan := rows.Scan(&ID_PRODUCT, &PRODUCT_NAME, &BRAND, &PRICE, &CATEGORY)
		if errScan != nil {
			return nil, errScan
		}

		inventory, errFetch := inventory.NewInventory(inventory.DTOInventory{
			Id:          ID_PRODUCT,
			ProductName: PRODUCT_NAME,
			Brand:       BRAND,
			Price:       PRICE,
			Category:    CATEGORY,
		})

		if errFetch != nil {
			return nil, errFetch
		}

		listInventory = append(listInventory, inventory)
	}
	defer rows.Close()

	return listInventory, nil
}
