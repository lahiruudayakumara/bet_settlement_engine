package store

import (
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"sync"
)

type EventStore struct {
	mu     sync.Mutex
	events map[string]*models.Event
}

// NewEventStore initializes a new in-memory store for events.
func NewEventStore() *EventStore {
	return &EventStore{
		events: make(map[string]*models.Event),
	}
}

// AddEvent adds a new event to the store.
func (s *EventStore) AddEvent(event *models.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event.EventID] = event
}

// GetEvent retrieves an event by its ID.
func (s *EventStore) GetEvent(eventID string) (*models.Event, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	event, exists := s.events[eventID]
	return event, exists
}

// GetAllEvents retrieves all events in the store.
func (s *EventStore) GetAllEvents() []*models.Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	var all []*models.Event
	for _, e := range s.events {
		all = append(all, e)
	}
	return all
}
