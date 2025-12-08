package util

import (
	"database/sql"
	"fmt"
	"regexp"
)

func IsValidTableName(name string) bool {
	valid := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return valid.MatchString(name)
}

func PrintWordRows(rows *sql.Rows) {
	for rows.Next() {
		var word string
		var definition string
		var tag string

		rows.Scan(&word, &definition, &tag)
		fmt.Println(word, definition, tag)
	}
}
