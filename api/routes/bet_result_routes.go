package routes

import (
	"github.com/gorilla/mux"
	"github.com/lahiruudayakumara/bet_settlement_engine/api/handler"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

func RegisterBetResultRoutes(r *mux.Router) {
	betResultStore := store.NewBetResultStore()
	betResultService := services.NewBetResultService(betResultStore)
	betResultHandler := handlers.NewBetResultHandler(betResultService)

	betResultRouter := r.PathPrefix("/bet_result").Subrouter()

	betResultRouter.HandleFunc("/", betResultHandler.AddBetResultHandler).Methods("POST")
	
	betResultRouter.HandleFunc("/", betResultHandler.GetBetResultHandler).Methods("GET")
}
