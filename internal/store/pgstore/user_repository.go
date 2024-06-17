package pgstore

import (
	"database/sql"

	model "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/user"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
)

type UserRepository struct {
	st *Store
}

func (r *UserRepository) Create(u *model.User) error {
	// Validate
	err := u.Validate()
	if err != nil {
		return err
	}

	// Encrypt
	err = u.Sequre()
	if err != nil {
		return err
	}

	// Save
	return r.st.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncPasswd,
	).Scan(&u.Id)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{
		Id: id,
	}
	err := r.st.db.QueryRow(
		"SELECT email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.Email,
		&u.EncPasswd,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{
		Email: email,
	}
	err := r.st.db.QueryRow(
		"SELECT id, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Id,
		&u.EncPasswd,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
