package word

import (
	"database/sql"
	"fmt"

	"github.com/NureddinFarzaliyev/go-dict-cli/internal/util"
)

func AddWord(conn *sql.DB, dict string, word string, def string, tag string) (err error) {
	isValidDictName := util.IsValidTableName(dict)
	if !isValidDictName {
		return fmt.Errorf("%s is not a valid dictionary name", dict)
	}

	query := fmt.Sprintf("INSERT INTO %s (word, definition, tag) VALUES (?, ?, ?)", dict)
	stmt, err := conn.Prepare(query)

	if err != nil {
		return err
	}
	_, execErr := stmt.Exec(word, def, tag)

	if execErr != nil {
		return execErr
	} else {
		fmt.Printf("%s added to %s\n", word, dict)
		return nil
	}
}

func ListWords(conn *sql.DB, dict string) (err error) {
	isValidDictName := util.IsValidTableName(dict)

	if !isValidDictName {
		return fmt.Errorf("%s is not a valid dictionary name", dict)
	}

	query := fmt.Sprintf("SELECT * FROM %s", dict)
	rows, execErr := conn.Query(query)

	if execErr != nil {
		return execErr
	}

	defer rows.Close()
	util.PrintWordRows(rows)
	return rows.Err()
}

func RemoveWord(conn *sql.DB, dict string, word string) (err error) {
	isValidDictName := util.IsValidTableName(dict)

	if !isValidDictName {
		return fmt.Errorf("%s is not a valid dictionary name", dict)
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE word IS "%s"`, dict, word)
	_, errExec := conn.Exec(query)

	if errExec != nil {
		return errExec
	} else {
		fmt.Printf("%s deleted from %s\n", word, dict)
		return nil
	}
}
