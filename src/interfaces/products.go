package interfaces

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"

	// "github.com/Irsad99/LectronicApp/src/input"
	"mime/multipart"
)

type ProductRepo interface {
	FindAll() (*models.Products, error)
	FindByID(id int) (*models.Product, error)
	SearchByName(name string) (*models.Products, error)
	SortByCategory(category string) (*models.Products, error)
	Add(data *models.Product) (*models.Product, error)
	Delete(id int) (*models.Product, error)
	Update(id int, data *models.Product) (*models.Product, error)
}

type ProductService interface {
	FindAll() (*helpers.Response, error)
	FindByID(id int) (*helpers.Response, error)
	SearchByName(name string) (*helpers.Response, error)
	SortByCategory(category string) (*helpers.Response, error)
	Add(data *input.InputProduct, file multipart.File, handle *multipart.FileHeader) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
	Update(id int, data *models.Product) (*helpers.Response, error)
	Upload(id int, file multipart.File, handle *multipart.FileHeader) (*helpers.Response, error)
}
