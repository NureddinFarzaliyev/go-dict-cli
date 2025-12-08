package search

import (
	"database/sql"
	"fmt"

	"github.com/NureddinFarzaliyev/go-dict-cli/internal/util"
)

func SearchWordInDictionary(conn *sql.DB, dict string, word string, tag bool) (err error) {
	isValidDictName := util.IsValidTableName(dict)
	if !isValidDictName {
		return fmt.Errorf("%s is not a valid dictionary name", dict)
	}

	var query string

	if tag {
		query = fmt.Sprintf("SELECT * FROM %s WHERE tag LIKE ?", dict)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s WHERE word LIKE ?", dict)
	}

	rows, err := conn.Query(query, "%"+word+"%")

	if err != nil {
		return err
	}

	defer rows.Close()
	util.PrintWordRows(rows)
	return rows.Err()
}
