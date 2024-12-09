package models

import "time"

type ShoppingCart struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
