package store

import (
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"sync"
)

type BetStore struct {
	mu   sync.Mutex
	bets map[string]*models.Bet
}

func NewBetStore() *BetStore {
	return &BetStore{bets: make(map[string]*models.Bet)}
}

func (s *BetStore) AddBet(bet *models.Bet) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bets[bet.BetID] = bet
}

func (store *BetStore) GetBet(betID string) (*models.Bet, bool) { // Correct return type
	bet, exists := store.bets[betID]
	if !exists {
		return nil, false
	}
	return bet, true // Return *models.Bet, not *BetStore
}
