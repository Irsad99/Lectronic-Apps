package interfaces

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
)

type UserRepo interface {
	GetAll() (*models.Users, error)
	GetId(id int) (*models.User, error)
	GetEmail(email string) (*models.User, error)
	AddUser(data *models.User) (*models.User, error)
	Update(id int, data *models.User) (*models.User, error)
	Delete(id int) (*models.User, error)
}

type UserService interface {
	GetAll() (*helpers.Response, error)
	GetId(id int) (*helpers.Response, error)
	GetEmail(email string) (*helpers.Response, error)
	AddUser(data *models.User) (*helpers.Response, error)
	Update(id int, data *models.User) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
}