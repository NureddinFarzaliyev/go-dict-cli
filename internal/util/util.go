package util

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"
)

func IsValidTableName(name string) bool {
	valid := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return valid.MatchString(name)
}

func PrintWordRows(rows *sql.Rows) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprint(w, "────────────\t")
	fmt.Fprint(w, "────────────\t")
	fmt.Fprint(w, "────────────\t")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "%s\t", "Word")
	fmt.Fprintf(w, "%s\t", "Definition")
	fmt.Fprintf(w, "%s\t", "Tag")
	fmt.Fprintln(w)
	fmt.Fprint(w, "────────────\t")
	fmt.Fprint(w, "────────────\t")
	fmt.Fprint(w, "────────────\t")
	fmt.Fprintln(w)

	for rows.Next() {
		var word string
		var definition string
		var tag string

		rows.Scan(&word, &definition, &tag)
		fmt.Fprintf(w, "%v\t", word)
		fmt.Fprintf(w, "%v\t", definition)
		fmt.Fprintf(w, "%v\t", tag)
		fmt.Fprintln(w)
		// fmt.Println(word, definition, tag)
	}

	w.Flush()
}

func PrintDictionaryList(rows *sql.Rows) error {
	fmt.Println("Existing Dictionaries:")

	fmt.Println("────────────")

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}
		fmt.Println(name)
	}

	fmt.Println("────────────")

	return nil
}
