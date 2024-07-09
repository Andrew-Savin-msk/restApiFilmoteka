package pgstore

import (
	"fmt"

	film "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/film"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
)

type FilmRepository struct {
	st *Store
}

func (f *FilmRepository) CreateAndConnectActors(film *film.Film, actors []int) error {

	tx, err := f.st.db.Begin()
	if err != nil {
		// TODO: Think about special Err statment to detect internal issues during server running
		return err
	}

	err = tx.QueryRow(
		"INSERT INTO films (name, description, release_date, assesment) VALUES ($1, $2, $3, $4) RETURNING id",
		film.Name,
		film.Desc,
		film.Date,
		film.Assesment,
	).Scan(&film.Id)
	if err != nil {
		tErr := tx.Rollback()
		if tErr != nil {
			return fmt.Errorf("rollback error: %e, triggered by error: %e", tErr, err)
		}
		return err
	}

	err = f.st.FilmActor().CreateConnections(tx, actors, film.Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (f *FilmRepository) Delete(id int) (int, error) {
	res, err := f.st.db.Exec(
		"DELETE FROM films WHERE id = $1",
		id,
	)
	if err != nil {
		return -1, err
	}

	am, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	if am == 0 {
		return -1, store.ErrRecordNotFound
	}

	return int(am), nil
}
