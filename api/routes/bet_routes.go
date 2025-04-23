package routes

import (
	"github.com/gorilla/mux"
	"github.com/lahiruudayakumara/bet_settlement_engine/api/handler"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

func RegisterBetRoutes(r *mux.Router) {
	betStore := store.NewBetStore()
	userStore := store.GetStore()
	betService := services.NewBetService(betStore, userStore)

	betHandler := handlers.NewBetHandler(betService)

	betRouter := r.PathPrefix("/bet").Subrouter()

	betRouter.HandleFunc("/", betHandler.CreateBetHandler).Methods("POST")
	betRouter.HandleFunc("/{betID}", betHandler.GetBetHandler).Methods("GET")
	betRouter.HandleFunc("/{betID}/settle", betHandler.SettleBetHandler).Methods("PUT")
	betRouter.HandleFunc("/{betID}/cancel", betHandler.CancelBetHandler).Methods("PUT")
}
