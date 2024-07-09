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

	// TODO: Change to Exec statment
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

func (a *ActorRepository) Delete(id int) (int, error) {
	res, err := a.st.db.Exec(
		"DELETE FROM actors WHERE id = $1",
		id,
	)
	if err != nil {
		return -1, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	if rows == 0 {
		return -1, store.ErrRecordNotFound
	}
	return id, nil
}

func (a *ActorRepository) Overwright(act *actor.Actor) error {
	res, err := a.st.db.Exec(
		"UPDATE actors SET "+
			"gender = CASE WHEN gender <> $1 AND $1 <> '' THEN $1 ELSE gender END, "+
			"birthdate = CASE WHEN birthdate <> $2 AND $3 IS NOT FALSE THEN $2 ELSE birthdate END, "+
			"name = CASE WHEN name <> $4 AND $4 <> '' THEN $4 ELSE name END "+
			"WHERE id = $5",
		act.Gen,
		act.Birthdate,
		act.Birthdate.IsZero(),
		act.Name,
		act.Id,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return store.ErrRecordNotFound
	}

	return nil
}

// TODO: Actors must be returned with their films
func (a *ActorRepository) GetAll() ([]*actor.Actor, error) {
	actors := []*actor.Actor{}
	rows, err := a.st.db.Query(
		"SELECT id, name, gender, birthdate FROM actors LIMIT 20",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var actor actor.Actor
		err := rows.Scan(&actor.Id, &actor.Name, &actor.Gen, &actor.Birthdate)
		if err != nil {
			return nil, err
		}
		actors = append(actors, &actor)
	}
	return actors, nil
}
