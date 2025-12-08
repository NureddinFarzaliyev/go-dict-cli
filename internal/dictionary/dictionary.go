package dictionary

import (
	"database/sql"
	"fmt"

	"github.com/NureddinFarzaliyev/go-dict-cli/internal/util"
)

func ListDictionaries(conn *sql.DB) (err error) {
	query := `
		SELECT name
		FROM sqlite_master
		WHERE type = 'table'
			AND name NOT LIKE 'sqlite_%'`

	rows, err := conn.Query(query)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}
		fmt.Println(name)
	}

	return rows.Err()
}

func CreateDictionary(conn *sql.DB, dictName string) (err error) {
	isValid := util.IsValidTableName(dictName)
	if isValid {
		query := fmt.Sprintf(`
			CREATE TABLE %s (
				word varchar(255) NOT NULL UNIQUE, 
				definition varchar(255) NOT NULL,
				tag varchar(255)
			)`, dictName)

		_, err := conn.Exec(query)

		if err != nil {
			return err
		} else {
			fmt.Println("Created new dictionary:", dictName)
			return nil
		}
	} else {
		return fmt.Errorf("Cannot create table that starts with numbers or include special characters.")
	}
}

func DeleteDictionary(conn *sql.DB, dictName string) (err error) {
	isValid := util.IsValidTableName(dictName)
	if !isValid {
		return fmt.Errorf("Cannot use table that starts with numbers or include special characters.")
	}

	query := fmt.Sprintf("DROP TABLE %s", dictName)
	_, exErr := conn.Exec(query)

	if exErr != nil {
		return exErr
	}

	fmt.Println("Deleted dictionary", dictName)
	return nil
}
