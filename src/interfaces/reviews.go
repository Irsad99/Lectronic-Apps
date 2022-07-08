package interfaces

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/helpers"
)

type ReviewRepo interface {
	FindByID(id int) (*models.Review, error)
	Add(data *models.Review) (*models.Review, error)
}

type ReviewService interface {
	FindByID(id int) (*helpers.Response, error)
	Add(id int, data *input.InputReview) (*helpers.Response, error)
}