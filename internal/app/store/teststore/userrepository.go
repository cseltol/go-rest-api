package teststore

import (
	"github.com/cseltol/go-rest-api/internal/app/model"
	"github.com/cseltol/go-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeUserCreation(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

func (r *UserRepository) FindByID(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if email == u.Email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
