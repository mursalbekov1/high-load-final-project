package models

import "time"

type OrderItem struct {
	Id        int       `json:"id"`
	OrderId   int       `json:"order_id"`
	ProductId int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
