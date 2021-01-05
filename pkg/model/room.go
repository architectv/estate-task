package model

type Room struct {
	Id          int    `json:"room_id" db:"id"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
}
