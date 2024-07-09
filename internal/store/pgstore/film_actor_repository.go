package pgstore

import (
	"database/sql"
	"fmt"
)

type FilmActorRepository struct {
	st *Store
}

// Doesn't forget rollback in any error case
func (far *FilmActorRepository) CreateConnections(tx *sql.Tx, actors []int, filmId int) error {
	stmt, err := tx.Prepare("INSERT INTO films_actors (film_id, actor_id) VALUES ($1, $2)")
	// No rollback due to no changes
	if err != nil {
		tErr := tx.Rollback()
		if tErr != nil {
			return fmt.Errorf("rollback error: %e, triggered by error: %e", tErr, err)
		}
		return err
	}

	for _, actorId := range actors {
		_, err := stmt.Exec(filmId, actorId)
		if err != nil {
			tErr := tx.Rollback()
			if tErr != nil {
				return fmt.Errorf("rollback error: %e, triggered by error: %e", tErr, err)
			}
			return err
		}
	}

	return nil
}
