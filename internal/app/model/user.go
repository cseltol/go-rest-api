package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string
	Passowrd          string
	EncryptedPassword string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u, 
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Passowrd, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 32)),
	)
}

func (u *User) BeforeUserCreation() error {
	if len(u.Passowrd) > 0 {
		enc, err := encryptString(u.Passowrd)
		if err != nil {
			return nil
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func encryptString(pswd string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
