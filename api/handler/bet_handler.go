package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"log"
	"net/http"
)

// BetHandler struct to hold the BetService
type BetHandler struct {
	BetService *services.BetService
}

// NewBetHandler creates a new BetHandler with a BetService instance
func NewBetHandler(service *services.BetService) *BetHandler {
	return &BetHandler{BetService: service}
}

// CreateBetHandler creates a new bet.
func (h *BetHandler) CreateBetHandler(w http.ResponseWriter, r *http.Request) {
	type reqBody struct {
		UserID  string  `json:"user_id"`
		EventID string  `json:"event_id"`
		Amount  float64 `json:"amount"`
		Odds    float64 `json:"odds"`
	}

	var body reqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	bet, err := h.BetService.PlaceBet(body.UserID, body.EventID, body.Amount, body.Odds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(bet)
}

// GetBetHandler retrieves a specific bet by its ID.
func (h *BetHandler) GetBetHandler(w http.ResponseWriter, r *http.Request) {
	betID := mux.Vars(r)["betID"]
	log.Printf("Received request for bet with ID: %s", betID)

	// Retrieve the bet from BetStore
	bet, exists := h.BetService.BetStore.GetBet(betID)
	if !exists {
		log.Printf("Bet with ID %s not found", betID)
		http.Error(w, "bet not found", http.StatusNotFound)
		return
	}

	// Return the bet as a response
	log.Printf("Found bet: %+v", bet)
	json.NewEncoder(w).Encode(bet)
}

// SettleBetHandler settles a bet based on its ID and the 'won' query parameter.
func (h *BetHandler) SettleBetHandler(w http.ResponseWriter, r *http.Request) {
	betID := mux.Vars(r)["betID"]
	won := r.URL.Query().Get("won") == "true"

	bet, err := h.BetService.SettleBet(betID, won)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(bet)
}

// CancelBetHandler cancels an existing bet based on its ID.
func (h *BetHandler) CancelBetHandler(w http.ResponseWriter, r *http.Request) {
	betID := mux.Vars(r)["betID"]

	bet, err := h.BetService.CancelBet(betID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(bet)
}
