package orders

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type service struct {
	repository interfaces.OrderRepository
}

func NewService(repository interfaces.OrderRepository) *service {
	return &service{repository: repository}
}

func (s *service) FindAll() (*helpers.Response, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return helpers.New(orders, 200, true), nil
}

func (s *service) FindByID(id string) (*helpers.Response, error) {
	order, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return helpers.New(order, 200, true), nil
}

func (s *service) FindByUserID(id string) (*helpers.Response, error) {
	order, err := s.repository.FindByUserID(id)
	if err != nil {
		return nil, err
	}

	return helpers.New(order, 200, true), nil
}

func (s *service) Create(order *models.Order) (*helpers.Response, error) {
	order, err := s.repository.Save(order)
	if err != nil {
		return nil, err
	}

	return helpers.New(order, 200, true), nil
}

func (s *service) Update(id string, input *models.Order) (*helpers.Response, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	data.Status = input.Status

	order, err := s.repository.Update(input)
	if err != nil {
		return nil, err
	}

	return helpers.New(order, 200, true), nil
}

func (s *service) Delete(id string) (*helpers.Response, error) {
	err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	return helpers.New("Successfully delete order", 200, true), nil
}
