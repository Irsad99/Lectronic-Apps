package models

import "time"

type Order struct {
	ID          uint64     `json:"id"`
	ProductID   uint64     `json:"product_id"`
	UserID      uint64     `json:"user_id"`
	Status      string     `json:"status"`
	TotalPrice  int64      `json:"total_price"`
	PaymentURL  string     `json:"payment_url"`
	PaidAt      time.Time  `json:"paid_at"`
	DeliveredAt *time.Time `json:"delivered_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Orders []Order
