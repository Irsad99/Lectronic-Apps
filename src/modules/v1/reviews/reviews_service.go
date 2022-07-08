package reviews

import (
	// "github.com/Irsad99/LectronicApp/src/database/gorm/models"

	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
	// "github.com/asaskevich/govalidator"
)

type review_service struct {
	repo interfaces.ReviewRepo
}

func NewService(svc interfaces.ReviewRepo) *review_service {
	return &review_service{svc}
}

func (svc *review_service) FindByID(id int) (*helpers.Response, error) {

	result, err := svc.repo.FindByID(id)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *review_service) Add(id int, data *input.InputReview) (*helpers.Response, error) {

	var review models.Review

	review.ProductID = data.ProductID
	review.UserID = id
	review.Comment = data.Comment
	review.Rating = float32(data.Rating)

	result, err := svc.repo.Add(&review)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}