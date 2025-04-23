package routes

import (
	"github.com/gorilla/mux"
	"github.com/lahiruudayakumara/bet_settlement_engine/api/handler"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

func RegisterEventRoutes(r *mux.Router) {
	eventStore := store.NewEventStore()
	eventService := services.NewEventService(eventStore)

	eventHandler := handlers.NewEventHandler(eventService)

	r.HandleFunc("/events", eventHandler.CreateEventHandler).Methods("POST")
	r.HandleFunc("/events/{event_id}", eventHandler.GetEventHandler).Methods("GET")
	r.HandleFunc("/events", eventHandler.GetAllEventsHandler).Methods("GET")
}
