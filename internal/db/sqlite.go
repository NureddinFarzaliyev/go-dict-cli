package db

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./dict.db")

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

	// fmt.Println("- SQL:", sqliteVersion)

	return db, nil
}
