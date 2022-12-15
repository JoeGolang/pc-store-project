package mapper

import (
	"github.com/rocketlaunchr/dbq/v2"
	user2 "pc-shop-final-project/domain/entity/user"
	"pc-shop-final-project/internal/delivery/http/models"
)

func DataUserDbToEntity(dataDTO user2.DTOUser) (*user2.User, error) {
	user, err := user2.NewUser(dataDTO)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UserEntityToModel(usr *user2.User) *models.ModelUser {
	return &models.ModelUser{
		ID_USER:     usr.GetValueIdUsr(),
		NAME:        usr.GetValueNameUsr(),
		OUTLET_CODE: usr.GetValueOutletCodeUsr(),
		STATUS:      usr.GetValueStatusUsr(),
	}
}

func UserEntityToDbqStruct(user *user2.User) []interface{} {
	dbqStruct := dbq.Struct(UserEntityToModel(user))
	return dbqStruct
}

func UserModelToEntity(model *models.ModelUser) (*user2.User, error) {
	user, err := user2.NewUser(user2.DTOUser{
		Id:         model.ID_USER,
		Name:       model.NAME,
		OutletCode: model.OUTLET_CODE,
		Status:     model.STATUS,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
