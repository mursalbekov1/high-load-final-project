package models

import "time"

type CarItem struct {
	Id        int       `json:"id"`
	CartId    int       `json:"carId"`
	ProductId int       `json:"productId"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
