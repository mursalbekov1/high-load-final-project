package models

import "time"

type Wishlist struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}