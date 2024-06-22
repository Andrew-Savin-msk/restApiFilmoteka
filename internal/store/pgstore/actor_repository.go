package pgstore

import (
	"database/sql"

	actor "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/actor"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
)

type ActorRepository struct {
	st *Store
}

func (a *ActorRepository) Create(act *actor.Actor) error {
	err := act.Validate()
	if err != nil {
		return err
	}

	return a.st.db.QueryRow(
		"INSERT INTO actors (name, gender, birthdate) VALUES ($1, $2, $3) RETURNING id",
		act.Name,
		act.Gen,
		act.Birthdate,
	).Scan(
		&act.Id,
	)
}

func (a *ActorRepository) Find(id int) (*actor.Actor, error) {
	act := &actor.Actor{
		Id: id,
	}

	err := a.st.db.QueryRow(
		"SELECT gender, birthdate, name FROM actors WHERE id = $1",
		id,
	).Scan(
		&act.Gen,
		&act.Birthdate,
		&act.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return act, nil
}
