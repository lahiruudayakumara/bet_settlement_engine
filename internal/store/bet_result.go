package store

import (
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"sync"
)

type BetResultStore struct {
	mu         sync.Mutex
	betResults map[string]*models.BetResult
}

func NewBetResultStore() *BetResultStore {
	return &BetResultStore{betResults: make(map[string]*models.BetResult)}
}

// AddBetResult adds a new BetResult to the store
func (s *BetResultStore) AddBetResult(betResult *models.BetResult) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.betResults[betResult.BetID] = betResult
}

// GetBetResult retrieves a BetResult from the store by BetID
func (s *BetResultStore) GetBetResult(betID string) (*models.BetResult, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	betResult, exists := s.betResults[betID]
	return betResult, exists
}
