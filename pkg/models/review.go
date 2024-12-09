package models

import "time"

//id, product_id, user_id, rating, comment, created_at, updated_at

type Review struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id"`
	UserId    int       `json:"user_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
