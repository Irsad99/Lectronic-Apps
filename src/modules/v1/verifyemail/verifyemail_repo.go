package verifyemail

import (
	"errors"

	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *user_repo {
	return &user_repo{grm}
}

func (repo *user_repo) FindByToken(token string) (*models.Token, error) {
	var verify models.Token

	result := repo.db.Where("token = ?", token).First(&verify)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &verify, nil
}

func (repo *user_repo) InsertToken(data *models.Token) (*models.Token, error) {
	var verify models.Token

	result := repo.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambahkan token")
	}

	getdata := repo.db.First(&verify, &data.ID)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &verify, nil
}

func (repo *user_repo) DeleteToken(token string) (*models.Token, error) {
	var verify models.Token

	del := repo.db.Where("token = ?", token).Delete(&verify)

	if del.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	return &verify, nil
}
