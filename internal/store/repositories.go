package store

import "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/models/user"

type UserRepository interface {
	Create(*user.User) error
	Find(int) (*user.User, error)
}
