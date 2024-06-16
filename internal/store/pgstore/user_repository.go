package pgstore

import "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/models/user"

type UserRepository struct {
	st *Store
}

func (r *UserRepository) Create(req *user.User) error {
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
	r.st.db.QueryRow("INSERT INTO users (email, en)")
	return nil
}

func (r *UserRepository) Find(id int) (*user.User, error) {

}
