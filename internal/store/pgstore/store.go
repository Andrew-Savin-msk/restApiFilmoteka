package pgstore

import (
	"database/sql"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
)

type Store struct {
	db *sql.DB
	ur *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s Store) Close() {
	s.db.Close()
}

func (s *Store) User() store.UserRepository {
	if s.ur == nil {
		s.ur = &UserRepository{
			st: s,
		}
	}
	return s.ur
}
