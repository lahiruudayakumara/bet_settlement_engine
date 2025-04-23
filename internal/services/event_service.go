package services

import (
	"errors"
	"time"

	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

// EventService manages events and handles business logic.
type EventService struct {
	store *store.EventStore
}

// NewEventService initializes a new EventService.
func NewEventService(store *store.EventStore) *EventService {
	return &EventService{
		store: store,
	}
}

// CreateEvent creates a new event and adds it to the store.
func (s *EventService) CreateEvent(name, status, outcome string, startTime, endTime time.Time) (*models.Event, error) {
	eventID := generateEventID() // Assume generateEventID is implemented
	event := &models.Event{
		EventID:   eventID,
		Name:      name,
		Status:    status,
		StartTime: startTime,
		EndTime:   endTime,
		Outcome:   outcome,
	}
	s.store.AddEvent(event)
	return event, nil
}

// GetEvent retrieves an event by its ID.
func (s *EventService) GetEvent(eventID string) (*models.Event, error) {
	event, exists := s.store.GetEvent(eventID)
	if !exists {
		return nil, errors.New("event not found")
	}
	return event, nil
}

// GetAllEvents retrieves all events.
func (s *EventService) GetAllEvents() ([]*models.Event, error) {
	events := s.store.GetAllEvents()
	return events, nil
}

// generateEventID generates a unique ID for the event.
func generateEventID() string {
	// This can be replaced with your ID generation logic (e.g., using UUIDs)
	return time.Now().Format("20060102150405")
}
