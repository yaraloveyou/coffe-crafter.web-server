package store

import (
	"time"

	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/model"
)

type UserRepository struct {
	store *Store
}

// Create
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password, username, last_activity) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Email,
		u.EncryptedPassword,
		u.Username,
		time.Now(),
	).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

// FindByEmail
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password, username, last_activity FROM users WHERE email = $1",
		email,
	).Scan(&user.ID,
		&user.Email,
		&user.Username,
		&user.EncryptedPassword,
		&user.LastActivity,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
