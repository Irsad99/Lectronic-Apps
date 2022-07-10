package payments

import (
	"os"
	"strconv"

	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
	"github.com/veritrans/go-midtrans"
)

type service struct {
	orderRepository   interfaces.OrderRepository
	productRepository interfaces.ProductRepo
}

func NewService(orderRepository interfaces.OrderRepository, productRepository interfaces.ProductRepo) *service {
	return &service{orderRepository, productRepository}
}

func (s *service) GetPaymentURL(orderID uint64, order *models.Order, user *models.User) (string, error) {
	oid := strconv.FormatUint(orderID, 10)

	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midclient.APIEnvType = midtrans.Sandbox

	coreGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	chargeReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Fullname,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  oid,
			GrossAmt: order.TotalPrice,
		},
	}

	snapTokenResp, err := coreGateway.GetToken(chargeReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

func (s *service) ProcessPayment(input *input.OrderNotificationInput) (*helpers.Response, error) {
	order, err := s.orderRepository.FindByID(input.OrderID)
	if err != nil {
		return nil, err
	}

	if input.PaymentType == "credit_card" && input.OrderStatus == "capture" && input.FraudStatus == "accept" {
		order.Status = "paid"
	} else if input.OrderStatus == "settlement" {
		order.Status = "paid"
	} else if input.OrderStatus == "deny" || input.OrderStatus == "expired" || input.OrderStatus == "cancel" {
		order.Status = "cancelled"
	}

	_, err = s.orderRepository.Update(order)
	if err != nil {
		return nil, err
	}

	res := helpers.New(order, 200, false)
	return res, nil
}
