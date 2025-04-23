package services

import (
	"fmt"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
	"time"
)

type BetTransactionService struct {
	betTransactionStore *store.BetTransactionStore
}

func NewBetTransactionService(betTransactionStore *store.BetTransactionStore) *BetTransactionService {
	return &BetTransactionService{
		betTransactionStore: betTransactionStore,
	}
}

// CreateBetTransaction creates a bet transaction when a bet is placed or settled.
func (s *BetTransactionService) CreateBetTransaction(bet *models.BetTransaction, transactionType string, amount float64) (*models.BetTransaction, error) {
	transaction := &models.BetTransaction{
		TransactionID: fmt.Sprintf("%d-%d", time.Now().Unix(), time.Now().Nanosecond()),
		BetID:         bet.BetID,
		UserID:        bet.UserID,
		Amount:        amount,
		Type:          transactionType,
		Timestamp:     time.Now(),
	}

	s.betTransactionStore.AddBetTransaction(transaction)
	return transaction, nil
}
