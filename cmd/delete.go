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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a dictionary",
	Long: `Delete a dictionary by its name

- Usage:
dict delete <dict-name>`,
	Run: func(cmd *cobra.Command, args []string) {
		WithDB(func(conn *sql.DB) (err error) {
			if len(args) == 1 {
				return dictionary.DeleteDictionary(conn, args[0])
			}
			fmt.Println(cmd.Long)
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
