package orders

import (
	"strconv"
	"time"

	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type service struct {
	repository     interfaces.OrderRepository
	userRepo       interfaces.UserRepo
	productRepo    interfaces.ProductRepo
	paymentService interfaces.PaymentService
}

func NewService(repository interfaces.OrderRepository, userRepo interfaces.UserRepo, paymentService interfaces.PaymentService, productRepo interfaces.ProductRepo) *service {
	return &service{repository, userRepo, productRepo, paymentService}
}

func (s *service) FindAll() (*helpers.Response, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return helpers.New(orders, 200, true), nil
}

func (s *service) FindByID(id int) (*helpers.Response, error) {
	order, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return helpers.New(order, 200, true), nil
}

func (s *service) FindByUserID(id int) (*helpers.Response, error) {
	order, err := s.repository.FindByUserID(id)
	if err != nil {
		return nil, err
	}

	return helpers.New(order, 200, true), nil
}

func (s *service) Create(id uint64, input *input.OrderInput) (*helpers.Response, error) {
	var order models.Order

	product, err := s.productRepo.FindByID(int(input.ProductID))
	if err != nil {
		return nil, err
	}

	price, _ := strconv.Atoi(product.Price)

	order.ProductID = input.ProductID
	order.UserID = id
	order.Status = "pending"
	order.TotalPrice = int64(price)
	order.PaidAt = time.Now()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	user, err := s.userRepo.GetId(int(id))
	if err != nil {
		return nil, err
	}

	paymentUrl, err := s.paymentService.GetPaymentURL(order.ID, &order, user)
	if err != nil {
		return nil, err
	}

	order.PaymentURL = paymentUrl

	newOrder, err := s.repository.Save(&order)
	if err != nil {
		return nil, err
	}

	return helpers.New(newOrder, 200, true), nil
}

func (s *service) Update(id int, input *models.Order) (*helpers.Response, error) {
	_, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	order, err := s.repository.Update(input)
	if err != nil {
		return nil, err
	}

	return helpers.New(order, 200, true), nil
}

func (s *service) Delete(id int) (*helpers.Response, error) {
	err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	return helpers.New("Successfully delete order", 200, true), nil
}
