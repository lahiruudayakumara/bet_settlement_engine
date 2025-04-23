package handlers

import (
	"encoding/json"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"net/http"
)

type BetTransactionHandler struct {
	service *services.BetTransactionService
}

func NewBetTransactionHandler(service *services.BetTransactionService) *BetTransactionHandler {
	return &BetTransactionHandler{service: service}
}

// CreateBetTransactionHandler handles POST requests for creating bet transactions.
func (h *BetTransactionHandler) CreateBetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var betTransaction models.BetTransaction
	if err := json.NewDecoder(r.Body).Decode(&betTransaction); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.CreateBetTransaction(&betTransaction, betTransaction.Type, betTransaction.Amount)
	if err != nil {
		http.Error(w, "failed to create transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}
