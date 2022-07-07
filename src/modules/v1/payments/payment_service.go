package payments

import (
	"os"

	"github.com/Irsad99/LectronicApp/src/interfaces"
	"github.com/veritrans/go-midtrans"
)

type service struct {
	repository        interfaces.OrderRepository
	productRepository interfaces.ProductRepository
}

func NewService(repository interfaces.OrderRepository, productRepository interfaces.ProductRepository) *service {
	return &service{repository, productRepository}
}

func (s *service) GetPaymentURL(orderID string, order *models.Order, user *models.User) (string, error) {

	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midclient.APIEnvType = midtrans.Sandbox

	coreGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	chargeReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(order.TotalPrice),
		},
	}

	snapTokenResp, err := coreGateway.GetToken(chargeReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

// func (s *service) ProcessPayment(input *input.OrderNotificationInput) (*helpers.Response, error) {
// 	order, err := s.repository.FindByID(input.OrderID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if input.PaymentType == "credit_card" && input.OrderStatus == "capture" && input.FraudStatus == "accept" {
// 		order.Status = "paid"
// 	} else if input.OrderStatus == "settlement" {
// 		order.Status = "paid"
// 	} else if input.OrderStatus == "deny" || input.OrderStatus == "expired" || input.OrderStatus == "cancel" {
// 		order.Status = "cancelled"
// 	}

// 	_, err = s.repository.Update(id, order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := helpers.New(order, 200, false)
// 	return res, nil
// }
