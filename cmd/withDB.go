package cmd

import (
	"database/sql"
	"fmt"

	"github.com/NureddinFarzaliyev/go-dict-cli/internal/db"
)

func WithDB(fn func(conn *sql.DB) error) {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println("error connecting:", err)
		return
	}
	defer conn.Close()

	if err := fn(conn); err != nil {
		fmt.Println("error:", err)
	}
}
