package repository

import (
	"errors"
	"supply-chain-monitor/models"
	"sync"
)

// UserRepository stores users in memory safely with a mutex.
type UserRepository struct {
	mu      sync.RWMutex
	users   map[string]models.User // key = username
	counter uint
}

// NewUserRepository creates a new empty UserRepository.
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]models.User),
	}
}

// Save adds a new user to the repository.
func (r *UserRepository) Save(user models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Username]; exists {
		return errors.New("username already exists")
	}

	r.counter++
	user.ID = r.counter
	r.users[user.Username] = user
	return nil
}

// FindByUsername retrieves a user by username.
func (r *UserRepository) FindByUsername(username string) (models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[username]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
