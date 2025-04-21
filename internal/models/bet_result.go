package models

import "time"

type BetResult struct {
	BetID     string    `json:"bet_id"`
	Outcome   string    `json:"outcome"`
	Payout    float64   `json:"payout"`
	SettledAt time.Time `json:"settled_at"`
}
