package webserver

import (
	"database/sql"
	"net/http"

	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := connDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	server := newServer(store)

	return http.ListenAndServe(config.BindAddr, server)
}

func connDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
