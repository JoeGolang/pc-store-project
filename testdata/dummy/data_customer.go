package dummy

import "pc-shop-final-project/domain/entity/customer"

func MakeDummyCustomer() []*customer.Customer {
	DataCustomer1 := &customer.DTOCustomer{
		Name: "Mina",
	}
	Customer1, err := customer.NewCustomer(*DataCustomer1, "BGR1")
	if err != nil {
		panic("wrong data customer")
	}

	DataCustomer2 := &customer.DTOCustomer{
		Name: "Momo",
	}
	Customer2, err := customer.NewCustomer(*DataCustomer2, "BGR1")
	if err != nil {
		panic("wrong data customer")
	}

	DataCustomer3 := &customer.DTOCustomer{
		Name: "Sana",
	}
	Customer3, err := customer.NewCustomer(*DataCustomer3, "BGR2")
	if err != nil {
		panic("wrong data customer")
	}

	listCustomer := make([]*customer.Customer, 0)
	listCustomer = append(listCustomer, Customer1, Customer2, Customer3)
	return listCustomer
}
