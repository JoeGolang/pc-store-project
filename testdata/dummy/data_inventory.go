package dummy

import "pc-shop-final-project/domain/entity/inventory"

func MakeDummyInventory() []*inventory.Inventory {
	DataInven1 := &inventory.DTOInventory{
		Id:          100001,
		ProductName: "Playstation 5",
		Brand:       "Sony",
		Price:       11000000,
		Category:    "New Console",
	}
	Inven1, err := inventory.NewInventory(*DataInven1)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven2 := &inventory.DTOInventory{
		Id:          100002,
		ProductName: "Playstation 4",
		Brand:       "Sony",
		Price:       5500000,
		Category:    "New Console",
	}
	Inven2, err := inventory.NewInventory(*DataInven2)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven3 := &inventory.DTOInventory{
		Id:          100003,
		ProductName: "Nintendo Switch",
		Brand:       "Nintendo",
		Price:       4000000,
		Category:    "New Console",
	}
	Inven3, err := inventory.NewInventory(*DataInven3)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven4 := &inventory.DTOInventory{
		Id:          200001,
		ProductName: "GTA Collection PS5",
		Brand:       "Rockstar Games",
		Price:       1200000,
		Category:    "New Game",
	}
	Inven4, err := inventory.NewInventory(*DataInven4)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven5 := &inventory.DTOInventory{
		Id:          200002,
		ProductName: "FIFA 2022 PS4",
		Brand:       "EA",
		Price:       700000,
		Category:    "New Game",
	}
	Inven5, err := inventory.NewInventory(*DataInven5)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven6 := &inventory.DTOInventory{
		Id:          200003,
		ProductName: "RESIDENT EVIL 7 PS5",
		Brand:       "CAPCOM",
		Price:       750000,
		Category:    "New Game",
	}
	Inven6, err := inventory.NewInventory(*DataInven6)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven7 := &inventory.DTOInventory{
		Id:          300001,
		ProductName: "RESIDENT EVIL 7 PS5",
		Brand:       "CAPCOM",
		Price:       400000,
		Category:    "Second Game",
	}
	Inven7, err := inventory.NewInventory(*DataInven7)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven8 := &inventory.DTOInventory{
		Id:          300002,
		ProductName: "Story of Seasons : Mineral Town Wii",
		Brand:       "Natsume",
		Price:       250000,
		Category:    "Second Game",
	}
	Inven8, err := inventory.NewInventory(*DataInven8)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven9 := &inventory.DTOInventory{
		Id:          400001,
		ProductName: "Joystick PS5",
		Brand:       "Sony",
		Price:       880000,
		Category:    "Accessories Console",
	}
	Inven9, err := inventory.NewInventory(*DataInven9)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven10 := &inventory.DTOInventory{
		Id:          400002,
		ProductName: "Cable HDMI High-Speed",
		Brand:       "Vention",
		Price:       70000,
		Category:    "Accessories Console",
	}
	Inven10, err := inventory.NewInventory(*DataInven10)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven11 := &inventory.DTOInventory{
		Id:          400003,
		ProductName: "Cable HDMI Ultra",
		Brand:       "Vention",
		Price:       190000,
		Category:    "Accessories Console",
	}
	Inven11, err := inventory.NewInventory(*DataInven11)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven12 := &inventory.DTOInventory{
		Id:          500001,
		ProductName: "Upgrade SSD M.2 Sata",
		Brand:       "Service",
		Price:       900000,
		Category:    "Service Console",
	}
	Inven12, err := inventory.NewInventory(*DataInven12)
	if err != nil {
		panic("wrong data inventory")
	}

	DataInven13 := &inventory.DTOInventory{
		Id:          500002,
		ProductName: "Service & Cleaning",
		Brand:       "Service",
		Price:       150000,
		Category:    "Service Console",
	}
	Inven13, err := inventory.NewInventory(*DataInven13)
	if err != nil {
		panic("wrong data inventory")
	}

	listInven := make([]*inventory.Inventory, 0)
	listInven = append(listInven, Inven1, Inven2, Inven3, Inven4, Inven5, Inven6, Inven7, Inven8, Inven9, Inven10, Inven11, Inven12, Inven13)
	return listInven
}
