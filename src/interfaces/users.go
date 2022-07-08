package interfaces

import (
	"mime/multipart"

	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
)

type UserRepo interface {
	GetAll() (*models.Users, error)
	GetId(id int) (*models.User, error)
	GetEmail(email string) (*models.User, error)
	AddUser(data *models.User) (*models.User, error)
	Update(id uint, data *models.User) (*models.User, error)
	Delete(id int) (*models.User, error)
}

type UserService interface {
	GetAll() (*helpers.Response, error)
	GetId(id int) (*helpers.Response, error)
	GetEmail(email string) (*helpers.Response, error)
	AddUser(input *input.RegisterInput) (*helpers.Response, error)
	Update(id uint, data *models.User) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
	Verify(token string) (*helpers.Response, error)
	Upload(id int, file multipart.File, handle *multipart.FileHeader) (*helpers.Response, error)
}
