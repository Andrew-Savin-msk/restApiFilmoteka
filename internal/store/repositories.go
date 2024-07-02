package store

import (
	actor "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/actor"
	user "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/user"
)

type UserRepository interface {
	Create(*user.User) error
	Find(int) (*user.User, error)
	FindByEmail(string) (*user.User, error)
}

type ActorRepository interface {
	Create(*actor.Actor) error
	Find(int) (*actor.Actor, error)
	Delete(int) (int, error)
	Overwright(*actor.Actor) error
	FindByNamePart(string) (*actor.Actor, error)
	GetAll() ([]*actor.Actor, error)
}
