package models

import "time"

type WishListItem struct {
	Id         int       `json:"id"`
	WishlistId int       `json:"wishlist_id"`
	ProductId  int       `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
}
