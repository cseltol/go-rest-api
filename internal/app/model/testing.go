package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email: "user@example.com",
		Passowrd: "password",
	}
}
