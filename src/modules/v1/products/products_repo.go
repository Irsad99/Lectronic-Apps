package products

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"errors"

	"gorm.io/gorm"
)

type product_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *product_repo {
	return &product_repo{grm}
}

func (repo *product_repo) FindAll() (*models.Products, error) {

	var product models.Products

	result := repo.db.Order("id_product desc").Find(&product)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &product, nil
}

func (repo *product_repo) FindByID(id int) (*models.Product, error) {

	var product models.Product

	result := repo.db.First(&product, id)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &product, nil
}

func (repo *product_repo) SearchByName(name string) (*models.Products, error) {

	var product models.Products

	result := repo.db.Where(`products."name" LIKE ?`, "%"+name+"%").Find(&product)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &product, nil
}

func (repo *product_repo) SortByCategory(category string) (*models.Products, error) {

	var product models.Products

	result := repo.db.Order("id_product desc").Where("category = ?", category).Find(&product)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &product, nil
}

func (repo *product_repo) Add(data *models.Product) (*models.Product, error) {

	var product models.Product

	result := repo.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambahkan data")
	}

	getdata := repo.db.First(&product, &data.Id_Product)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &product, nil
}

func (repo *product_repo) Delete(id int) (*models.Product, error) {

	var product models.Product

	getdata := repo.db.First(&product, id)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	result := repo.db.Delete(&models.Product{}, id)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	return &product, nil
}

func (repo *product_repo) Update(id int, data *models.Product) (*models.Product, error) {

	var product models.Product

	result := repo.db.Model(&models.Product{}).Where("id_product = ?", id).Updates(&models.Product{Name: data.Name, Price: data.Price, Category: data.Category, Description: data.Description, Stock: data.Stock, Image: data.Image, Sold: data.Sold})

	if result.Error != nil {
		return nil, errors.New("gagal meng-update data")
	}

	getData := repo.db.First(&product, id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &product, nil
}