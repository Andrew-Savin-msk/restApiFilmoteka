package pgstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestActorMap(t *testing.T) map[string]interface{} {

	return map[string]interface{}{
		"name":      "John",
		"gender":    "man",
		"birthdate": time.Now(),
	}
}

func TestStore(t *testing.T, dbURL string) (*sql.DB, func(...string)) {

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
