package models

import "time"

type Order struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	OrderStatus string    `json:"order_status"`
	TotalAmount float64   `json:"total_amount"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}
