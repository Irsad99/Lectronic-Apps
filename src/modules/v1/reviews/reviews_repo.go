package reviews

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"errors"

	"gorm.io/gorm"
)

type review_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *review_repo {
	return &review_repo{grm}
}

func (repo *review_repo) FindByID(id int) (*models.Review, error) {

	var review models.Review

	result := repo.db.First(&review, id)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &review, nil
}

func (repo *review_repo) Add(data *models.Review) (*models.Review, error) {

	var review models.Review

	result := repo.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambahkan data")
	}

	getdata := repo.db.First(&review, &data.ID_Review)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &review, nil
}