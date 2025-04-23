package api_test

import (
	"bytes"
	"encoding/json"
	handlers "github.com/lahiruudayakumara/bet_settlement_engine/api/handler"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

// SetupTestStore clears the singleton and sets a fresh one for test.
func SetupTestStore() *store.UserStore {
	s := store.NewUserStore()
	store.OverrideStore(s) // Assume you implement this method for testing
	return s
}

func TestUserHandler_Success(t *testing.T) {
	user := &models.User{
		UserID:    "u100",
		Username:  "testuser",
		Email:     "test@example.com",
		Balance:   100.0,
		CreatedAt: time.Now(),
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handlers.UserHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", rec.Code)
	}

	var got models.User
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.UserID != user.UserID {
		t.Errorf("expected userID %s, got %s", user.UserID, got.UserID)
	}
}

func TestUserHandler_UserAlreadyExists(t *testing.T) {
	s := SetupTestStore()

	existing := &models.User{
		UserID:   "u101",
		Username: "exists",
		Email:    "exists@example.com",
		Balance:  100.0,
	}
	s.AddUser(existing)

	body, _ := json.Marshal(existing)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handlers.UserHandler(rec, req)

	if rec.Code != http.StatusConflict {
		t.Errorf("expected 409 Conflict, got %d", rec.Code)
	}
}

func TestUserHandler_InvalidEmail(t *testing.T) {
	user := &models.User{
		UserID:   "u102",
		Username: "bademail",
		Email:    "not-an-email",
		Balance:  50.0,
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handlers.UserHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %d", rec.Code)
	}
}

func TestUserHandler_NegativeBalance(t *testing.T) {
	user := &models.User{
		UserID:   "u103",
		Username: "neguser",
		Email:    "valid@example.com",
		Balance:  -50.0,
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handlers.UserHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request for negative balance, got %d", rec.Code)
	}
}
