package inventory

import (
	"errors"
	"pc-shop-final-project/domain/value_object"
)

type Inventory struct {
	id          int
	productName string
	brand       string
	price       int
	category    *value_object.ProductCategory
}

type DTOInventory struct {
	Id          int
	ProductName string
	Brand       string
	Price       int
	Category    string
}

func NewInventory(inven *DTOInventory) (*Inventory, error) {
	if inven.ProductName == "" {
		return nil, errors.New("product name cannot be empty")
	}
	if inven.Brand == "" {
		return nil, errors.New("brand cannot be empty")
	}
	if inven.Category == "" {
		return nil, errors.New("product category cannot be empty")
	}

	category, errInven := value_object.NewProductCategory(value_object.ProductCategoryInt(inven.Category))
	if errInven != nil {
		return nil, errInven
	}

	return &Inventory{
		id:          inven.Id,
		productName: inven.ProductName,
		brand:       inven.Brand,
		price:       inven.Price,
		category:    category,
	}, nil
}

func (inv *Inventory) GetValueIdInv() int {
	return inv.id
}

func (inv *Inventory) GetValueProductNameInv() string {
	return inv.productName
}

func (inv *Inventory) GetValueBrandInv() string {
	return inv.brand
}

func (inv *Inventory) GetvaluePriceInv() int {
	return inv.price
}

func (inv *Inventory) GetValueCategoryInv() string {
	return inv.category.ProductCategoryToString()
}
