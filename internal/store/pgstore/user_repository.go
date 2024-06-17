package pgstore

import (
	"database/sql"

	model "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/user"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
)

type UserRepository struct {
	st *Store
}

func (r *UserRepository) Create(req *model.User) error {
	// Validate
	err := req.Validate()
	if err != nil {
		return err
	}

	// Encrypt
	err = req.Sequre()
	if err != nil {
		return err
	}

	// Save
	return r.st.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		req.Email,
		req.EncPasswd,
	).Scan(&req.Id)
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
