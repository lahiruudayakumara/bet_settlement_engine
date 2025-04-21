package models

import "time"

type BetTransaction struct {
	TransactionID string    `json:"transaction_id"`
	BetID         string    `json:"bet_id"`
	UserID        string    `json:"user_id"`
	Amount        float64   `json:"amount"`
	Type          string    `json:"type"`
	Timestamp     time.Time `json:"timestamp"`
}
