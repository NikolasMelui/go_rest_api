package store

import "github.com/nikolasmelui/go_rest_api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id;",
		u.Email, u.Password,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil

}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
