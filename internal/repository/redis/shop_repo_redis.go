package redis

import (
	"context"
	"encoding/json"
	"errors"
	"pc-shop-final-project/domain/entity"
	"pc-shop-final-project/internal/repository/redis/mapper"

	_interface "pc-shop-final-project/domain/repository"

	"github.com/go-redis/redis/v8"
)

type RepoShopRedis struct {
	Conn           *redis.Client
	RepoUser       _interface.InterfaceUser
	RepoCustomer   _interface.InterfaceCustomer
	RepoSettleItem _interface.InterfaceSettlementItem
}

func NewRepoShopRedis(Conn *redis.Client, user _interface.InterfaceUser, customer _interface.InterfaceCustomer, setit _interface.InterfaceSettlementItem) *RepoShopRedis {
	return &RepoShopRedis{
		Conn:           Conn,
		RepoUser:       user,
		RepoCustomer:   customer,
		RepoSettleItem: setit,
	}
}

func (s *RepoShopRedis) GetAllRedis(ctx context.Context) (*entity.User, *entity.Customer, []*entity.SettlePurchase, error) {
	var (
		dataMapUser map[string]string
		dataMapCust map[string]string
		dataMapItem map[string]string
		checkErr    error
	)
	dataMapUser, checkErr = s.Conn.HGetAll(ctx, "myShopUser").Result()
	if checkErr != nil && checkErr != redis.Nil && dataMapUser == nil {
		return nil, nil, nil, errors.New("no data")
	}
	dataMapCust, checkErr = s.Conn.HGetAll(ctx, "myShopCust").Result()
	if checkErr != nil && checkErr != redis.Nil && dataMapCust == nil {
		return nil, nil, nil, errors.New("no data")
	}
	dataMapItem, checkErr = s.Conn.HGetAll(ctx, "myShopItem").Result()
	if checkErr != nil && checkErr != redis.Nil && dataMapItem == nil {
		return nil, nil, nil, errors.New("no data")
	}

	ActiveUser := &entity.User{}
	ActiveCust := &entity.Customer{}

	dataUser := &mapper.User{}
	for _, dataU := range dataMapUser {
		errUnmarshal := json.Unmarshal([]byte(dataU), dataUser)
		if errUnmarshal != nil {
			return nil, nil, nil, errors.New("unmarshal fail")
		}

		User, errFetch := entity.NewUser(&entity.DTOUser{
			Id:         dataUser.IdUser,
			Name:       dataUser.NameUser,
			OutletCode: dataUser.OutletCode,
			Status:     dataUser.StatusUser,
		})
		if errFetch != nil {
			return nil, nil, nil, errFetch
		}
		ActiveUser = User

		dataCust := &mapper.Customer{}
		for _, dataC := range dataMapCust {
			errUnmarshal := json.Unmarshal([]byte(dataC), dataCust)
			if errUnmarshal != nil {
				return nil, nil, nil, errors.New("unmarshal fail")
			}

			Cust, errFetch := entity.NewCustomer(&entity.DTOCustomer{
				UniqId:   dataCust.IdCustomer,
				Name:     dataCust.NameCustomer,
				JoinDate: dataCust.JoinDate,
			}, dataUser.OutletCode)
			if errFetch != nil {
				return nil, nil, nil, errFetch
			}
			ActiveCust = Cust
		}
	}

	SettleList := make([]*entity.SettlePurchase, 0)
	dataItem := &mapper.SettleItem{}
	for _, dataI := range dataMapItem {
		errUnmarshal := json.Unmarshal([]byte(dataI), dataItem)
		if errUnmarshal != nil {
			return nil, nil, nil, errors.New("unmarshal fail")
		}
		settleItem := entity.NewSettlePurchase(dataItem.IdItem, dataItem.Qty)
		SettleList = append(SettleList, settleItem)
	}

	return ActiveUser, ActiveCust, SettleList, nil
}

func (s *RepoShopRedis) GetUserRedis(ctx context.Context) (*entity.User, error) {
	var (
		dataMapUser map[string]string
		checkErr    error
	)
	dataMapUser, checkErr = s.Conn.HGetAll(ctx, "myShopUser").Result()
	if checkErr != nil && checkErr != redis.Nil && dataMapUser == nil {
		return nil, errors.New("no data")
	}

	ActiveUser := &entity.User{}

	dataUser := &mapper.User{}
	for _, dataU := range dataMapUser {
		errUnmarshal := json.Unmarshal([]byte(dataU), dataUser)
		if errUnmarshal != nil {
			return nil, errors.New("unmarshal fail")
		}

		User, errFetch := entity.NewUser(&entity.DTOUser{
			Id:         dataUser.IdUser,
			Name:       dataUser.NameUser,
			OutletCode: dataUser.OutletCode,
			Status:     dataUser.StatusUser,
		})
		if errFetch != nil {
			return nil, errFetch
		}
		ActiveUser = User
	}
	return ActiveUser, nil
}

func (s *RepoShopRedis) GetUserCustomerRedis(ctx context.Context) (*entity.User, *entity.Customer, error) {
	var (
		dataMapUser map[string]string
		dataMapCust map[string]string
		checkErr    error
	)
	dataMapUser, checkErr = s.Conn.HGetAll(ctx, "myShopUser").Result()
	if checkErr != nil && checkErr != redis.Nil && dataMapUser == nil {
		return nil, nil, errors.New("no data")
	}
	dataMapCust, checkErr = s.Conn.HGetAll(ctx, "myShopCust").Result()
	if checkErr != nil && checkErr != redis.Nil && dataMapCust == nil {
		return nil, nil, errors.New("no data")
	}

	ActiveUser := &entity.User{}
	ActiveCust := &entity.Customer{}

	dataUser := &mapper.User{}
	for _, dataU := range dataMapUser {
		errUnmarshal := json.Unmarshal([]byte(dataU), dataUser)
		if errUnmarshal != nil {
			return nil, nil, errors.New("unmarshal fail")
		}

		User, errFetch := entity.NewUser(&entity.DTOUser{
			Id:         dataUser.IdUser,
			Name:       dataUser.NameUser,
			OutletCode: dataUser.OutletCode,
			Status:     dataUser.StatusUser,
		})
		if errFetch != nil {
			return nil, nil, errFetch
		}
		ActiveUser = User

		dataCust := &mapper.Customer{}
		for _, dataC := range dataMapCust {
			errUnmarshal := json.Unmarshal([]byte(dataC), dataCust)
			if errUnmarshal != nil {
				return nil, nil, errors.New("unmarshal fail")
			}

			Cust, errFetch := entity.NewCustomer(&entity.DTOCustomer{
				UniqId:   dataCust.IdCustomer,
				Name:     dataCust.NameCustomer,
				JoinDate: dataCust.JoinDate,
			}, dataUser.OutletCode)
			if errFetch != nil {
				return nil, nil, errFetch
			}
			ActiveCust = Cust
		}
	}
	return ActiveUser, ActiveCust, nil
}

func (s RepoShopRedis) StoreUser(ctx context.Context, dataU *entity.User) error {
	_, err := s.Conn.HSet(ctx, "myShopUser", 0, mapper.MapSetUserToString(dataU)).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s RepoShopRedis) StoreCustomer(ctx context.Context, dataC *entity.Customer) error {
	_, err := s.Conn.HSet(ctx, "myShopCust", 0, mapper.MapSetCustToString(dataC)).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s RepoShopRedis) StoreSetIt(ctx context.Context, dataI []*entity.SettlePurchase) error {
	SetIt := mapper.MapSetSettleItemToStringSlice(dataI)
	_, err := s.Conn.HMSet(ctx, "myShopItem", SetIt).Result()
	if err != nil {
		return err
	}
	return nil
}
