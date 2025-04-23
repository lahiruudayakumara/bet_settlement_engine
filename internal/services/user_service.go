package services

import (
	"fmt"
	"regexp"
	"time"

	"github.com/lahiruudayakumara/bet_settlement_engine/internal/errors"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
)

type UserService struct {
	userStore *store.UserStore
}

func NewUserService(userStore *store.UserStore) *UserService {
	return &UserService{userStore: userStore}
}

func (s *UserService) CreateUser(userID, username, email string, balance float64) (*models.User, error) {
	if _, exists := s.userStore.GetUser(userID); exists {
		return nil, errors.ErrUserAlreadyExists
	}

	if !isValidEmail(email) {
		return nil, errors.ErrInvalidEmail
	}

	if s.userStore.IsEmailExists(email) {
		return nil, fmt.Errorf("email %s is already in use", email)
	}

	if balance < 0 {
		return nil, errors.ErrNegativeBalance
	}

	user := &models.User{
		UserID:    userID,
		Username:  username,
		Email:     email,
		Balance:   balance,
		CreatedAt: time.Now(),
	}

	s.userStore.AddUser(user)
	return user, nil
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
