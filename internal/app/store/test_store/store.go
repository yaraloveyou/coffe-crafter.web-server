package test_store

import (
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/model"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
