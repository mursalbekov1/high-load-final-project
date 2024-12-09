package models

import "time"

type Payment struct {
	Id            int       `json:"id"`
	OrderId       int       `json:"order_id"`
	PaymentMethod string    `json:"payment_method"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
