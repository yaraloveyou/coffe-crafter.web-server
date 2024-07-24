package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository == nil {
		s.UserRepository = &UserRepository{
			store: s,
		}
	}

	return s.UserRepository
}
