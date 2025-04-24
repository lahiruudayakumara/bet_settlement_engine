package scripts

import (
	"time"

	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

func SeedData(s *store.InMemoryStore) {
	// Users
	user := &models.User{
		UserID:    "user-001",
		Username:  "john_doe",
		Balance:   1000.00,
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}
	s.AddUser(user)

	// Events
	event := &models.Event{
		EventID:   "event-001",
		Name:      "Team A vs Team B",
		Status:    "upcoming",
		StartTime: time.Now().Add(2 * time.Hour),
		EndTime:   time.Now().Add(4 * time.Hour),
	}
	s.AddEvent(event)

	// Bets
	bet := &models.Bet{
		BetID:   "bet-001",
		UserID:  user.UserID,
		EventID: event.EventID,
		Amount:  100.00,
		Odds:    2.5,
		Status:  models.BetStatusPending,
		BetTime: time.Now(),
	}
	s.AddBet(bet)

	// Transactions
	//tx := &models.BetTransaction{
	//	TransactionID: "txn-001",
	//	BetID:         bet.BetID,
	//	UserID:        user.UserID,
	//	Amount:        bet.Amount,
	//	Type:          "debit",
	//	Timestamp:     time.Now(),
	//}
	//s.AddTransaction(tx)
	//
	//// Bet Results (optional, if already settled)
	//result := &models.BetResult{
	//	BetID:     bet.BetID,
	//	Outcome:   "win",
	//	Payout:    250.00,
	//	SettledAt: time.Now(),
	//}
	//s.AddResult(result)
}
