package store

import (
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"sync"
)

type BetTransactionStore struct {
	mu              sync.Mutex
	betTransactions map[string]*models.BetTransaction
}

func NewBetTransactionStore() *BetTransactionStore {
	return &BetTransactionStore{
		betTransactions: make(map[string]*models.BetTransaction),
	}
}

// AddBetTransaction adds a new transaction to the store.
func (s *BetTransactionStore) AddBetTransaction(transaction *models.BetTransaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.betTransactions[transaction.TransactionID] = transaction
}

// GetBetTransaction retrieves a transaction by its ID.
func (s *BetTransactionStore) GetBetTransaction(transactionID string) (*models.BetTransaction, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	transaction, exists := s.betTransactions[transactionID]
	return transaction, exists
}

// GetAllTransactions returns all transactions.
func (s *BetTransactionStore) GetAllTransactions() map[string]*models.BetTransaction {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.betTransactions
}
