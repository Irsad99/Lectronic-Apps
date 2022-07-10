package interfaces

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
)

type OrderRepository interface {
	FindAll() (*models.Orders, error)
	FindByID(id int) (*models.Order, error)
	FindByUserID(id int) (*models.Orders, error)
	Save(order *models.Order) (*models.Order, error)
	Update(order *models.Order) (*models.Order, error)
	Delete(id int) error
}

type OrderService interface {
	FindAll() (*helpers.Response, error)
	FindByID(id int) (*helpers.Response, error)
	FindByUserID(id int) (*helpers.Response, error)
	Create(id uint64, input *input.OrderInput) (*helpers.Response, error)
	Update(id int, input *models.Order) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
}
