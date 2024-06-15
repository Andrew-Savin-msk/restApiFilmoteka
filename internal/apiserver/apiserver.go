package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/config"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store/pgstore"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	// _ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	ErrDbTypeUnknown = fmt.Errorf("unknown database name")
)

func Start(cfg *config.Config) error {

	st, err := loadStore(cfg.DbPath, cfg.DbType, cfg.SchemaPath)
	if err != nil {
		return fmt.Errorf("unable to init db. ended with error: %s", err)
	}
	defer st.Close()

	srv := server{
		mux:    newMux(),
		logger: setLog(cfg.LogLevel),
		store:  st,
	}

	err = http.ListenAndServe(cfg.Port, srv.mux)
	if err != nil {
		return fmt.Errorf("unable to start listening port. ended with error: %s", err)
	}
	return nil
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	// Place for setting endpoints

	return mux
}

func setLog(level string) *logrus.Logger {
	log := logrus.New()
	switch strings.ToLower(level) {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	}
	return log
}

func loadStore(url, sType, migrationsPath string) (store.Store, error) {
	switch strings.ToLower(sType) {
	case "postgres", "psql", "pg4":
		return loadPg(url, migrationsPath)
	}
	return nil, ErrDbTypeUnknown
}

func loadPg(url, migrationsPath string) (store.Store, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("open: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// migrations, err := migrate.New(migrationsPath, url)
	// if err != nil {
	// 	return nil, fmt.Errorf("migrate: %v", err)
	// }

	// err = migrations.Force(20240608145718)

	// if err != nil {
	// 	return nil, fmt.Errorf("migrate.force: %v", err)
	// }

	// err = migrations.Up()
	// if err != nil {
	// 	return nil, fmt.Errorf("migrate.up: %v", err)
	// }

	return pgstore.New(db), nil
}

// func isValidDb(db *sql.DB) (bool, error) {
// 	tables := ""
// 	for _, table := range tables {
// 		var exists bool
// 		query := fmt.Sprintf(`SELECT EXISTS (
//             SELECT FROM information_schema.tables
//             WHERE table_schema = 'public'
//             AND table_name = '%s'
//         )`, table)

// 		err := db.QueryRow(query).Scan(&exists)
// 		if err != nil {
// 			return false, err
// 		}

// 		if !exists {
// 			return false, nil
// 		}
// 	}
// 	return true, nil
// }
