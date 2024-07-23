package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/model"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/store"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	user := &model.User{
		Email:             "test@test.com",
		EncryptedPassword: "password",
		Username:          "test",
	}
	err := s.User().Create(user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "test@test.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	user := &model.User{
		Email:             "test@test.com",
		EncryptedPassword: "password",
		Username:          "test",
	}
	s.User().Create(user)
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
