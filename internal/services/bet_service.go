package services

import (
	"fmt"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
	"time"
)

type BetService struct {
	BetStore  *store.BetStore
	UserStore *store.UserStore
}

func NewBetService(betStore *store.BetStore, userStore *store.UserStore) *BetService {
	return &BetService{
		BetStore:  betStore,
		UserStore: userStore,
	}
}

func (s *BetService) PlaceBet(userID, eventID string, amount, odds float64) (*models.Bet, error) {
	user, exists := s.UserStore.GetUser(userID)
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	if user.Balance < amount {
		return nil, fmt.Errorf("insufficient balance")
	}

	bet := &models.Bet{
		BetID:   fmt.Sprintf("%d", time.Now().UnixNano()),
		UserID:  userID,
		EventID: eventID,
		Amount:  amount,
		Odds:    odds,
		Status:  models.BetStatusPending,
		BetTime: time.Now(),
	}

	user.Balance -= amount
	s.BetStore.AddBet(bet)
	s.UserStore.AddUser(user)

	return bet, nil
}

func (s *BetService) SettleBet(betID string, won bool) (*models.Bet, error) {
	bet, exists := s.BetStore.GetBet(betID)
	if !exists {
		return nil, fmt.Errorf("bet not found")
	}
	if bet.Status != models.BetStatusPending {
		return nil, fmt.Errorf("bet is already settled or cancelled")
	}

	user, exists := s.UserStore.GetUser(bet.UserID)
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	if won {
		bet.Status = models.BetStatusWon
		user.Balance += bet.Amount * bet.Odds
	} else {
		bet.Status = models.BetStatusLost
	}

	bet.SettlementTime = time.Now()
	s.UserStore.AddUser(user)
	return bet, nil
}

func (s *BetService) CancelBet(betID string) (*models.Bet, error) {
	bet, exists := s.BetStore.GetBet(betID)
	if !exists {
		return nil, fmt.Errorf("bet not found")
	}
	if bet.Status != models.BetStatusPending {
		return nil, fmt.Errorf("cannot cancel non-pending bet")
	}

	user, exists := s.UserStore.GetUser(bet.UserID)
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	bet.Status = models.BetStatusCancelled
	user.Balance += bet.Amount
	bet.SettlementTime = time.Now()

	s.UserStore.AddUser(user)
	return bet, nil
}
