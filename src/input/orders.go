package input

type OrderInput struct {
	ProductID uint64 `json:"product_id"`
	// TotalPrice int64    `json:"total_price"`
}

type OrderNotificationInput struct {
	OrderStatus string `json:"transaction_status" validate:"required"`
	OrderID     int    `json:"order_id" validate:"required"`
	PaymentType string `json:"payment_type" validate:"required"`
	FraudStatus string `json:"fraud_status" validate:"required"`
}
