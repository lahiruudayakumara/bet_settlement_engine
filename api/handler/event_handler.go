package handlers

import (
	"encoding/json"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"net/http"
	"time"
)

// EventHandler handles the HTTP requests for events.
type EventHandler struct {
	service *services.EventService
}

// NewEventHandler initializes a new EventHandler.
func NewEventHandler(service *services.EventService) *EventHandler {
	return &EventHandler{
		service: service,
	}
}

// CreateEventHandler handles POST requests to create a new event.
func (h *EventHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var eventRequest struct {
		Name      string    `json:"name"`
		Status    string    `json:"status"`
		Outcome   string    `json:"outcome"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	}

	if err := json.NewDecoder(r.Body).Decode(&eventRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	event, err := h.service.CreateEvent(eventRequest.Name, eventRequest.Status, eventRequest.Outcome, eventRequest.StartTime, eventRequest.EndTime)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

// GetEventHandler handles GET requests to retrieve an event by ID.
func (h *EventHandler) GetEventHandler(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Query().Get("event_id")

	// Get event using the EventService
	event, err := h.service.GetEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(event)
}

// GetAllEventsHandler handles GET requests to retrieve all events.
func (h *EventHandler) GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := h.service.GetAllEvents()
	if err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}
