package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
)

// BetResultHandler struct holds a reference to the BetResultService.
type BetResultHandler struct {
	service *services.BetResultService
}

// NewBetResultHandler creates a new BetResultHandler.
func NewBetResultHandler(service *services.BetResultService) *BetResultHandler {
	return &BetResultHandler{service: service}
}

// AddBetResultHandler handles the POST request to add a BetResult.
func (h *BetResultHandler) AddBetResultHandler(w http.ResponseWriter, r *http.Request) {
	var betResult models.BetResult
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&betResult)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.AddBetResult(betResult.BetID, betResult.Outcome, betResult.Payout, betResult.SettledAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(betResult)
}

// GetBetResultHandler handles the GET request to retrieve a BetResult.
func (h *BetResultHandler) GetBetResultHandler(w http.ResponseWriter, r *http.Request) {
	betID := r.URL.Query().Get("bet_id")
	if betID == "" {
		http.Error(w, "Bet ID is required", http.StatusBadRequest)
		return
	}

	betResult, err := h.service.GetBetResult(betID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(betResult)
}
