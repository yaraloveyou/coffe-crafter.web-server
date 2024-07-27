package webserver

import (
	"database/sql"
	"net/http"

	redisstore "github.com/yaraloveyou/coffe-crafter.web-server/internal/app/store/redis_store"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := connDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	rdb := redisstore.New(config.RedisAddr)
	server := newServer(store, rdb)

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
