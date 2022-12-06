package mysql

import (
	"context"
	"database/sql"
	"pc-shop-final-project/domain/entity"
	_interface "pc-shop-final-project/domain/repository"
	"time"
)

type SettleItemMysqlInteractor struct {
	db *sql.DB
}

func NewSettlePurchaseMysql(db *sql.DB) _interface.InterfaceSettlementItem {
	return &SettleItemMysqlInteractor{db: db}
}

// CreatePurchase implements _interface.InterfaceSettlementItem
func (set *SettleItemMysqlInteractor) CreatePurchase(ctx context.Context, purchase []*entity.SettlePurchase) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	setAutoInc := "ALTER TABLE settlement_purchased AUTO_INCREMENT=1"
	_, errMysql = set.db.Exec(setAutoInc)

	if errMysql != nil {
		return errMysql
	}

	for _, purc := range purchase {
		insertQuery := "INSERT INTO settlement_purchased(ID_SETTLEMENT, ID_ITEM, QTY) VALUES (?,?,?)"

		_, errMysql = set.db.Exec(insertQuery, purc.GetValueIdProduct(), purc.GetValueIdSettlement())

		if errMysql != nil {
			return errMysql
		}
	}
	return nil
}

// DeletePurchase implements _interface.InterfaceSettlementItem
func (set *SettleItemMysqlInteractor) DeletePurchase(ctx context.Context, codeItem string) error {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteQuery := "DELETE FROM settlement_purchased WHERE CODE_TRANSACTION = ?"

	_, errMysql = set.db.Exec(deleteQuery, codeItem)

	if errMysql != nil {
		return errMysql
	}

	return nil
}

// ReadPurchase implements _interface.InterfaceSettlementItem
func (set *SettleItemMysqlInteractor) ReadPurchase(ctx context.Context) ([]*entity.SettlePurchase, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT ID_SETTLEMENT, ID_ITEM, QTY FROM settlement_purchased"
	rows, errMysql := set.db.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, errMysql
	}

	listSettleP := make([]*entity.SettlePurchase, 0)
	for rows.Next() {
		var (
			ID_SETTLEMENT int
			ID_ITEM       int
			QTY           int
		)

		errScan := rows.Scan(&ID_SETTLEMENT, &ID_ITEM, &QTY)
		if errScan != nil {
			return nil, errScan
		}

		setP := entity.NewSettlePurchase(ID_SETTLEMENT, ID_ITEM, QTY)

		listSettleP = append(listSettleP, setP)
	}
	defer rows.Close()

	return listSettleP, nil
}
