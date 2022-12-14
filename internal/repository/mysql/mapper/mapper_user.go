package mapper

import (
	user2 "pc-shop-final-project/domain/entity/user"
	"pc-shop-final-project/internal/delivery/http/models"
)

//func DataBukuDbToEntity(dataDTO entity.DTOBuku) (*entity.Buku, error) {
//	buku, err := entity.NewBuku(dataDTO)
//	if err != nil {
//		return nil, err
//	}
//
//	return buku, nil
//}
//
//func DataPengarangDbToEntity(dataDTO entity.DTOPengarang) (*entity.Pengarang, error) {
//	pengarang, err := entity.NewPengarang(dataDTO)
//	if err != nil {
//		return nil, err
//	}
//
//	return pengarang, nil
//}
//
//func BukuEntityToModel(buku *entity.Buku) *models.ModelBuku {
//	date, _ := time.Parse("2006-01-02", buku.GetTahunTerbit())
//
//	return &models.ModelBuku{
//		IdPengarang: buku.GetIdPengarang(),
//		Judul:       buku.GetJudul(),
//		Category:    buku.GetCategory(),
//		TahunTerbit: date,
//		KodeBuku:    buku.GetKodeBuku(),
//	}
//}
//
//func BukuEntityToDbqStruct(buku *entity.Buku) []interface{} {
//	dbqStruct := dbq.Struct(BukuEntityToModel(buku))
//	return dbqStruct
//}

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
