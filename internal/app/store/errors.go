package store

import "errors"

var (
	ErrNotFound                 = errors.New("Entry not found")
	ErrIncorrectEmailOrPassword = errors.New("Incorrect email or password")
)
