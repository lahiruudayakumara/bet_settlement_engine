package store_test

import (
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
	"testing"
)

func TestUserStore(t *testing.T) {
	us := store.NewUserStore()

	user := &models.User{
		UserID:   "test1",
		Username: "Test User",
		Email:    "test@example.com",
		Balance:  100.0,
	}

	us.AddUser(user)

	got, exists := us.GetUser("test1")
	if !exists {
		t.Fatal("expected user to exist")
	}

	if got.Email != "test@example.com" {
		t.Errorf("expected email %s, got %s", "test@example.com", got.Email)
	}
}
