package model

import "time"

type User struct {
	ID                int
	Email             string
	EncryptedPassword string
	Username          string
	LastActivity      time.Time
}
