package interfaces

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
)

type PaymentService interface {
	GetPaymentURL(orderID uint64, order *models.Order, user *models.User) (string, error)
	ProcessPayment(input *input.OrderNotificationInput) (*helpers.Response, error)
}
