package dummy

import "pc-shop-final-project/domain/entity"

func MakeDummyCustomer() []*entity.Customer {
	DataCustomer1 := &entity.DTOCustomer{
		Name: "Mina",
	}
	Customer1, err := entity.NewCustomer(DataCustomer1, "BGR1")
	if err != nil {
		panic("wrong data customer")
	}

	DataCustomer2 := &entity.DTOCustomer{
		Name: "Momo",
	}
	Customer2, err := entity.NewCustomer(DataCustomer2, "BGR1")
	if err != nil {
		panic("wrong data customer")
	}

	DataCustomer3 := &entity.DTOCustomer{
		Name: "Sana",
	}
	Customer3, err := entity.NewCustomer(DataCustomer3, "BGR2")
	if err != nil {
		panic("wrong data customer")
	}

	listCustomer := make([]*entity.Customer, 0)
	listCustomer = append(listCustomer, Customer1, Customer2, Customer3)
	return listCustomer
}
