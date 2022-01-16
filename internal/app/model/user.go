package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int `json:"id"`
	Email             string `json:"email"`
	Passowrd          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
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

func (u *User) Sanitize() {
	u.Passowrd = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(pswd string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
