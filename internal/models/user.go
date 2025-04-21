package models

import "time"

type User struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Balance   float64   `json:"balance"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
