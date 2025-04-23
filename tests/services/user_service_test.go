package services_test

import (
	"testing"
	"time"

	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

func TestCreateUser(t *testing.T) {
	userStore := store.NewUserStore()
	service := services.NewUserService(userStore)

	userID := "u1"
	username := "alice"
	email := "alice@example.com"
	balance := 500.0

	user, err := service.CreateUser(userID, username, email, balance)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.Username != username {
		t.Errorf("expected username %s, got %s", username, user.Username)
	}
	if user.CreatedAt.After(time.Now()) {
		t.Error("user creation time should not be in the future")
	}
}
