package routes

import (
	"github.com/gorilla/mux"
	"github.com/lahiruudayakumara/bet_settlement_engine/api/handler"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

func RegisterBetTransactionRoutes(r *mux.Router) {
	betTransactionStore := store.NewBetTransactionStore()
	betTransactionService := services.NewBetTransactionService(betTransactionStore)
	betTransactionHandler := handlers.NewBetTransactionHandler(betTransactionService)

	betTransactionRouter := r.PathPrefix("/bet_transaction").Subrouter()
	betTransactionRouter.HandleFunc("/", betTransactionHandler.CreateBetTransactionHandler).Methods("POST")
}
