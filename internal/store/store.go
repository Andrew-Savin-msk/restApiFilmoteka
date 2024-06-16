package store

type Store interface {
	Close()
	User() UserRepository
}
