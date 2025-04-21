package models

import "time"

type Bet struct {
	BetID          string    `json:"bet_id"`
	UserID         string    `json:"user_id"`
	EventID        string    `json:"event_id"`
	Amount         float64   `json:"amount"`
	Odds           float64   `json:"odds"`
	Status         string    `json:"status"`
	BetTime        time.Time `json:"bet_time"`
	SettlementTime time.Time `json:"settlement_time"`
}

const (
	BetStatusPending   = "pending"
	BetStatusWon       = "won"
	BetStatusLost      = "lost"
	BetStatusCancelled = "cancelled"
)
