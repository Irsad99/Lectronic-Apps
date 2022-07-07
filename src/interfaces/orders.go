package interfaces

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
)

type OrderRepository interface {
	FindAll() (*models.Orders, error)
	FindByID(id string) (*models.Order, error)
	FindByUserID(id string) (*models.Order, error)
	Save(order *models.Order) (*models.Order, error)
	Update(order *models.Order) (*models.Order, error)
	Delete(id string) error
}

type OrderService interface {
	FindAll() (*helpers.Response, error)
	FindByID(id string) (*helpers.Response, error)
	FindByUserID(id string) (*helpers.Response, error)
	Create(order *models.Order) (*helpers.Response, error)
	Update(id string, input *models.Order) (*helpers.Response, error)
	Delete(id string) (*helpers.Response, error)
}
