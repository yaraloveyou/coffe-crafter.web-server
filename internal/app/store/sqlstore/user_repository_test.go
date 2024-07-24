package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/model"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/store/sqlstore"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	user := model.TestUser(t)
	err := s.User().Create(user)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "test@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	user := model.TestUser(t)
	user.Email = email
	s.User().Create(user)
	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
