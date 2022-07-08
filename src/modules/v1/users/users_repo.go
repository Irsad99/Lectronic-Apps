package users

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

// GET ALL USER
func (repo *user_repo) GetAll() (*models.Users, error) {

	var users models.Users
	result := repo.db.Order("id_user desc").Find(&users)

	if result.Error != nil {
		return nil, errors.New("error getting all data")
	}
	return &users, nil
}

// GET USER BY ID
func (repo *user_repo) GetId(id int) (*models.User, error) {
	var users models.User

	result := repo.db.First(&users, id)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &users, nil
}

// GET USER BY EMAIL
func (repo *user_repo) GetEmail(email string) (*models.User, error) {
	var user models.User
	result := repo.db.Where("email = ?", email).Find(&user)

	if result.Error != nil {
		return nil, errors.New("failed getting email")
	}
	return &user, nil
}

// ADD USER aka CREATE USER aka REGISTER
func (repo *user_repo) AddUser(data *models.User) (*models.User, error) {
	var users models.User

	result := repo.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambahkan data")
	}

	getdata := repo.db.First(&users, &data.IdUser)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &users, nil
}

// UPDATE USER
func (repo *user_repo) Update(id uint, data *models.User) (*models.User, error) {
	var users models.User

	result := repo.db.Model(&models.User{}).Where("vehicle_id = ?", id).Updates(&models.User{Fullname: data.Fullname, Address: data.Address, Image: data.Image, Birthdate: data.Birthdate, Phone: data.Phone, Gender: data.Gender, Verified: data.Verified})

	if result.Error != nil {
		return nil, errors.New("gagal meng-update data")
	}

	getData := repo.db.First(&users, id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &users, nil
}

// DELETE USER
func (repo *user_repo) Delete(id int) (*models.User, error) {
	var users models.User

	getdata := repo.db.First(&users, id)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	result := repo.db.Delete(&models.Product{}, id)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	return &users, nil
}
