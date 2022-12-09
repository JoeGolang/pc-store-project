package value_object

import "fmt"

type ProductCategory struct {
	value int
}

func NewProductCategory(category int) (*ProductCategory, error) {
	pc := ProductCategory{value: category}
	err := pc.validate()
	if err != nil {
		return &ProductCategory{}, err
	}
	return &pc, nil
}

func ProductCategoryInt(category string) int {
	switch category {
	case "New Console":
		return 1
	case "New Game":
		return 2
	case "Second Game":
		return 3
	case "Accessories Console":
		return 4
	case "Service Console":
		return 5
	default:
		return 0
	}
}

func (ic *ProductCategory) validate() error {
	switch ic.value {
	case 1:
		return nil
	case 2:
		return nil
	case 3:
		return nil
	case 4:
		return nil
	case 5:
		return nil
	default:
		return fmt.Errorf("product category tidak ditemukan")
	}
}

func (ic *ProductCategory) ProductCategoryToString() string {
	switch ic.value {
	case 1:
		return "New Console"
	case 2:
		return "New Game"
	case 3:
		return "Second Game"
	case 4:
		return "Accessories Console"
	case 5:
		return "Service Console"
	default:
		return ""
	}
}
