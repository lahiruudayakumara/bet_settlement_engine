package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lahiruudayakumara/bet_settlement_engine/internal/errors"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/services"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userService := services.NewUserService(store.GetStore())

	createdUser, err := userService.CreateUser(user.UserID, user.Username, user.Email, user.Balance)
	if err != nil {
		switch err {
		case errors.ErrUserAlreadyExists:
			http.Error(w, "User already exists", http.StatusConflict)
		case errors.ErrInvalidEmail:
			http.Error(w, "Invalid email format", http.StatusBadRequest)
		case errors.ErrNegativeBalance:
			http.Error(w, "Balance cannot be negative", http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	userStore := store.GetStore()
	users := userStore.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
