/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/NureddinFarzaliyev/go-dict-cli/internal/dictionary"
	"github.com/NureddinFarzaliyev/go-dict-cli/internal/word"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List dictionaries or the content of a dictionary",
	Long: `List all of the previously created dictionaries or list the content of a specific dictionary.

Example:
	List all dictionaries: dict list 
	List the content of a dictionary: dict list <dict-name>`,
	Run: func(cmd *cobra.Command, args []string) {
		WithDB(func(conn *sql.DB) error {
			if len(args) == 0 {
				return dictionary.ListDictionaries(conn)
			} else if len(args) == 1 {
				return word.ListWords(conn, args[0])
			} else {
				fmt.Println(cmd.Long)
				return nil
			}
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
