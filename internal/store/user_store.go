package store

import (
	"errors"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/models"
	"sync"
)

type UserStore struct {
	mu    sync.Mutex
	users map[string]*models.User
}

var storeInstance *UserStore
var once sync.Once

// GetStore returns the singleton instance of UserStore.
func GetStore() *UserStore {
	once.Do(func() {
		storeInstance = &UserStore{
			users: make(map[string]*models.User),
		}
	})
	return storeInstance
}

func (s *UserStore) AddUser(user *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the user already exists
	if _, exists := s.users[user.UserID]; exists {
		return errors.New("user already exists")
	}

	// Add the user to the store
	s.users[user.UserID] = user
	return nil
}

func (s *UserStore) GetUser(userID string) (*models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, exists := s.users[userID]
	return user, exists
}

func (s *UserStore) GetAllUsers() []*models.User {
	s.mu.Lock()
	defer s.mu.Unlock()
	var all []*models.User
	for _, u := range s.users {
		all = append(all, u)
	}
	return all
}

func (s *UserStore) IsEmailExists(email string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, u := range s.users {
		if u.Email == email {
			return true
		}
	}
	return false
}

func OverrideStore(s *UserStore) {
	storeInstance = s
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]*models.User),
	}
}
