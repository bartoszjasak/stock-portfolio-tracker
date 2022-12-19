package data

import (
	"time"
)

// var db *sql.DB

// func New(dbPool *sql.DB) Models {
// 	db = dbPool

// 	return Models{
// 		User: User{},
// 	}
// }

type Models struct {
	User        User
	Transaction Transaction
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type Transaction struct {
	ID       int       `json:"id"`
	Type     string    `json:"type"`
	Symbol   string    `json:"symbol"`
	Price    float32   `json:"price"`
	Quantity int       `json:"quantity"`
	Date     time.Time `json:"date"`
	UserId   int       `json:"user_id"`
}
