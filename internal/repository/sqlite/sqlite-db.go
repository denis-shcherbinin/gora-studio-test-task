package sqlite

import (
	"database/sql"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDb(cfg *config.SqliteConfig) (*sql.DB, error) {
	db, err := sql.Open(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
