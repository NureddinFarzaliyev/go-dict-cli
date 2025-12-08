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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new dictionary",
	Long: `Create a new dictionary with default word, definition and tag fields.

Example:
  dict create <dict-name>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			WithDB(func(conn *sql.DB) error {
				return dictionary.CreateDictionary(conn, args[0])
			})
		} else {
			fmt.Println(cmd.Long)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
