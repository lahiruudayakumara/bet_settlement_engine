package services

import (
	"errors"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
	"time"
)

// BetResultService handles business logic for BetResult.
type BetResultService struct {
	betResultStore *store.BetResultStore
}

// NewBetResultService creates a new instance of BetResultService.
func NewBetResultService(store *store.BetResultStore) *BetResultService {
	return &BetResultService{betResultStore: store}
}

// AddBetResult adds a new BetResult to the store.
func (s *BetResultService) AddBetResult(betID, outcome string, payout float64, settledAt time.Time) error {
	// Check if the bet already exists
	if _, exists := s.betResultStore.GetBetResult(betID); exists {
		return errors.New("bet result already exists")
	}

	// Create a new BetResult instance
	betResult := &models.BetResult{
		BetID:     betID,
		Outcome:   outcome,
		Payout:    payout,
		SettledAt: settledAt,
	}

	// Store the BetResult
	s.betResultStore.AddBetResult(betResult)
	return nil
}

// GetBetResult retrieves a BetResult by its BetID.
func (s *BetResultService) GetBetResult(betID string) (*models.BetResult, error) {
	betResult, exists := s.betResultStore.GetBetResult(betID)
	if !exists {
		return nil, errors.New("bet result not found")
	}
	return betResult, nil
}
