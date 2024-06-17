package pgstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, dbURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

	return db, func(s ...string) {
		if len(s) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", strings.Join(s, ", ")))
		}
	}
}
