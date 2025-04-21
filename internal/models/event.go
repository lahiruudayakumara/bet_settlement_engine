package models

import "time"

type Event struct {
	EventID   string    `json:"event_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Outcome   string    `json:"outcome"`
}
