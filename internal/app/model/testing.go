package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "test@test.com",
		Password: "password",
		Username: "test",
	}
}
