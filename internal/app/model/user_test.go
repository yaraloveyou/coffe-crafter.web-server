package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/model"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty username",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Username = ""
				return user
			},
			isValid: false,
		},
		{
			name: "format username",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Username = "ya"
				return user
			},
			isValid: false,
		},
		{
			name: "empty email",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Email = ""
				return user
			},
			isValid: false,
		},
		{
			name: "format email",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Email = "test@test"
				return user
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				return user
			},
			isValid: false,
		},
		{
			name: "with encrypted password",
			u: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				user.EncryptedPassword = "password"
				return user
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
