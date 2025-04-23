package api

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lahiruudayakumara/bet_settlement_engine/api/handler"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateEventHandler(t *testing.T) {
	// Setup
	eventStore := store.NewEventStore()
	eventService := services.NewEventService(eventStore)
	eventHandler := handlers.NewEventHandler(eventService)

	// Create a test HTTP request to create an event
	eventRequest := struct {
		Name      string    `json:"name"`
		Status    string    `json:"status"`
		Outcome   string    `json:"outcome"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	}{
		Name:      "Test Event",
		Status:    "active",
		Outcome:   "ongoing",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
	}

	body, _ := json.Marshal(eventRequest)
	req, err := http.NewRequest("POST", "/events", bytes.NewReader(body))
	assert.NoError(t, err)

	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Setup the router and route
	r := mux.NewRouter()
	r.HandleFunc("/events", eventHandler.CreateEventHandler).Methods("POST")

	// Serve HTTP
	r.ServeHTTP(rr, req)

	// Assert that the response status code is 201 Created
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Unmarshal the response body into the event model
	var createdEvent models.Event
	err = json.NewDecoder(rr.Body).Decode(&createdEvent)
	assert.NoError(t, err)

	// Assert that the event fields are correctly set
	assert.Equal(t, "Test Event", createdEvent.Name)
	assert.Equal(t, "active", createdEvent.Status)
	assert.Equal(t, "ongoing", createdEvent.Outcome)
}

func TestGetEventHandler(t *testing.T) {
	// Setup
	eventStore := store.NewEventStore()
	eventService := services.NewEventService(eventStore)
	eventHandler := handlers.NewEventHandler(eventService)

	// Create a test event to add to the store
	event := &models.Event{
		EventID:   "123",
		Name:      "Test Event",
		Status:    "active",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
		Outcome:   "ongoing",
	}
	eventStore.AddEvent(event)

	// Create a test HTTP request to get the event
	req, err := http.NewRequest("GET", "/events?event_id=123", nil)
	assert.NoError(t, err)

	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Setup the router and route
	r := mux.NewRouter()
	r.HandleFunc("/events", eventHandler.GetEventHandler).Methods("GET")

	// Serve HTTP
	r.ServeHTTP(rr, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Unmarshal the response body into the event model
	var retrievedEvent models.Event
	err = json.NewDecoder(rr.Body).Decode(&retrievedEvent)
	assert.NoError(t, err)

	// Assert that the retrieved event matches the one we added
	assert.Equal(t, event.EventID, retrievedEvent.EventID)
}

func TestGetAllEventsHandler(t *testing.T) {
	// Setup
	eventStore := store.NewEventStore()
	eventService := services.NewEventService(eventStore)
	eventHandler := handlers.NewEventHandler(eventService)

	// Create some test events
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

	// Create a test HTTP request to get all events
	req, err := http.NewRequest("GET", "/events", nil)
	assert.NoError(t, err)

	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Setup the router and route
	r := mux.NewRouter()
	r.HandleFunc("/events", eventHandler.GetAllEventsHandler).Methods("GET")

	// Serve HTTP
	r.ServeHTTP(rr, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Unmarshal the response body into the event model slice
	var events []models.Event
	err = json.NewDecoder(rr.Body).Decode(&events)
	assert.NoError(t, err)

	// Assert that we have 2 events in the response
	assert.Len(t, events, 2)
}
