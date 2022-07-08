package products

import (
	// "github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
	// "github.com/asaskevich/govalidator"
)

type product_service struct {
	repo interfaces.ProductRepo
}

func NewService(svc interfaces.ProductRepo) *product_service {
	return &product_service{svc}
}

func (svc *product_service) FindAll() (*helpers.Response, error) {

	result, err := svc.repo.FindAll()
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *product_service) FindByID(id int) (*helpers.Response, error) {

	result, err := svc.repo.FindByID(id)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *product_service) SearchByName(name string) (*helpers.Response, error) {

	result, err := svc.repo.SearchByName(name)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *product_service) SortByCategory(category string) (*helpers.Response, error) {

	result, err := svc.repo.SortByCategory(category)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *product_service) Add(data *input.InputProduct) (*helpers.Response, error) {

	// var product models.Product
	var product models.Product

	product.Name = data.Name
	product.Price = data.Price
	product.Category = data.Category
	product.Description = data.Description
	product.Image = data.Image
	product.Stock = data.Stock

	result, err := svc.repo.Add(&product)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *product_service) Delete(id int) (*helpers.Response, error) {

	// _, err := govalidator.ToInt(id)
	// if err != nil {
	// 	res := response.ResponseJSON(400, "Id yang anda masukan salah")
	// 	res.Message = err.Error()
	// 	return res, nil
	// }

	result, err := svc.repo.Delete(id)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}

func (svc *product_service) Update(id int, data *models.Product) (*helpers.Response, error) {

	// _, err := govalidator.ToInt(id)
	// if err != nil {
	// 	res := response.ResponseJSON(400, "Id yang anda masukan salah")
	// 	res.Message = err.Error()
	// 	return res, nil
	// }

	result, err := svc.repo.Update(id, data)
	if err != nil {
		res := helpers.New(result, 400, true)
		return res, nil
	}

	res := helpers.New(result, 200, false)
	return res, nil
}