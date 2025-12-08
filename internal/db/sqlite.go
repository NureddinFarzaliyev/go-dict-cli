package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/glebarez/go-sqlite"
)

func Connect() (*sql.DB, error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return nil, fmt.Errorf("get user home dir: %w", err)
	}

	dbPath := filepath.Join(home, ".local", "share", "go-dict-cli", "dict.db")

	err = os.MkdirAll(filepath.Dir(dbPath), 0755)
	if err != nil {
		return nil, fmt.Errorf("create db directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	var sqliteVersion string
	if err := db.QueryRow("select sqlite_version()").Scan(&sqliteVersion); err != nil {
		db.Close()
		return nil, fmt.Errorf("query sqlite_version %w", err)
	}

	return db, nil
}
