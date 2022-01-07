package model_test

import (
	"testing"

	"github.com/cseltol/go-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct{
		name string
		u func() *model.User
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
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Passowrd = ""

				return u
			},
			isValid: false,
		},
		{
			name: "with encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Passowrd = ""
				u.EncryptedPassword = "encrypted_password"

				return u
			},
			isValid: true,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid email"

				return u
			},
			isValid: false,
		},
		{
			name: "invalid password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Passowrd = "short"
				
				return u
			},
			isValid: false,
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

func TestUser_BeforeUserCreation(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.BeforeUserCreation())
	assert.NotEmpty(t, u.EncryptedPassword)
}
