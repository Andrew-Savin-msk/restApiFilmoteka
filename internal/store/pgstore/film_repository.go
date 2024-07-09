package pgstore

import (
	film "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/film"
)

type FilmRepository struct {
	st *Store
}

func (f *FilmRepository) Create(film *film.Film) error {
	err := f.st.db.QueryRow(
		"INSERT INTO films (name, description, release_date, assesment) VALUES ($1, $2, $3, $4) RETURNING id",
		film.Name,
		film.Desc,
		film.Date,
		film.Assesment,
	).Scan(&film.Id)
	if err != nil {
		return err
	}

	return nil
}
