package models

import "time"

type Category struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	ParentId  int       `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
