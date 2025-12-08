/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/NureddinFarzaliyev/go-dict-cli/internal/dictionary"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List dictionaries or the content of a dictionary",
	Long: `This command is used to list all of the previously created dictionaries or list the content of a specific dictionary.

- Usage:
dict list - lists all dictionaries
dict list <dict-name> - lists content of the <dict-name>`,
	Run: func(cmd *cobra.Command, args []string) {
		WithDB(func(conn *sql.DB) error {
			if len(args) == 0 {
				dictionary.ListDictionaries(conn)
				return nil
			} else if len(args) == 1 {
				fmt.Println("display content of", args[0])
				return nil
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
