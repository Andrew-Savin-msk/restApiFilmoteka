package pgstore

import (
	"database/sql"
	"fmt"

	actor "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/actor"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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

func (a *ActorRepository) Delete(id int) (int, error) {
	res, err := a.st.db.Exec(
		"DELETE FROM actors WHERE id = $1",
		id,
	)

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
		"UPDATE actors SET gender = $1, birthdate = $2, name = $3 WHERE id = $4",
		act.Gen,
		act.Birthdate,
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

func (s *ActorRepository) OverwrightFields(id int, fields map[string]interface{}) error {

	allowedKeys := []string{"id", "name", "gender", "birthdate"}
	err := validation.Validate(
		fields,
		validation.Map(
			validation.Key("name", validation.Length(1, 150)),
		),
		validation.By(validation.RuleFunc(ValidateMapFields(allowedKeys))),
	)
	if err != nil {
		return err
	}

	var data string
	for k, v := range fields {
		data = fmt.Sprintf("%s %s = %v,", data, k, v)
	}

	res, err := s.st.db.Exec(
		fmt.Sprintf("UPDATE actors SET %s WHERE id = $1", data[:len(data)-1]),
		id,
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
