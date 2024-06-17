package store

import model "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/user"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
