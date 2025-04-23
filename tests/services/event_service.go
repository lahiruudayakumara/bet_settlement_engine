package services

import (
	"testing"
	"time"

	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	eventStore := store.NewEventStore()
	eventService := services.NewEventService(eventStore)

	// Create a test event using the service
	event, err := eventService.CreateEvent("Test Event", "active", "ongoing", time.Now(), time.Now().Add(1*time.Hour))
	assert.NoError(t, err)

	// Assert that the event was created correctly
	assert.NotNil(t, event)
	assert.Equal(t, "Test Event", event.Name)
	assert.Equal(t, "active", event.Status)
	assert.Equal(t, "ongoing", event.Outcome)
}

func TestGetEvent(t *testing.T) {
	eventStore := store.NewEventStore()
	eventService := services.NewEventService(eventStore)

	// Create and add an event
	event := &models.Event{
		EventID:   "123",
		Name:      "Test Event",
		Status:    "active",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
		Outcome:   "ongoing",
	}
	eventStore.AddEvent(event)

	// Get the event using the service
	retrievedEvent, err := eventService.GetEvent("123")
	assert.NoError(t, err)

	// Assert that the retrieved event matches the expected event
	assert.Equal(t, event.EventID, retrievedEvent.EventID)
}

func TestGetAllEvents(t *testing.T) {
	eventStore := store.NewEventStore()
	eventService := services.NewEventService(eventStore)

	// Create and add some events
	event1 := &models.Event{
		EventID:   "123",
		Name:      "Test Event 1",
		Status:    "active",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
		Outcome:   "ongoing",
	}
	event2 := &models.Event{
		EventID:   "124",
		Name:      "Test Event 2",
		Status:    "completed",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
		Outcome:   "success",
	}
	eventStore.AddEvent(event1)
	eventStore.AddEvent(event2)

	// Get all events using the service
	events, err := eventService.GetAllEvents()
	assert.NoError(t, err)

	// Assert that we have 2 events
	assert.Len(t, events, 2)
}
