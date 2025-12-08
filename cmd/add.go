/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/NureddinFarzaliyev/go-dict-cli/internal/word"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add record to dictionary",
	Long: `Add new record consisting of the word, definition and optionally a tag to your existing dictionary.

- Usage:
add <dict-name> -w <word> -d <definition> -t <tag>`,
	Run: func(cmd *cobra.Command, args []string) {
		w, _ := cmd.Flags().GetString("w")
		d, _ := cmd.Flags().GetString("d")
		t, _ := cmd.Flags().GetString("t")

		if len(args) == 1 && w != "" && d != "" {
			WithDB(func(conn *sql.DB) error {
				return word.AddWord(conn, args[0], w, d, t)
			})
		} else {
			fmt.Println(cmd.Long)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().String("w", "", "Word to be added to the dictionary")
	addCmd.Flags().String("d", "", "Definition of the word")
	addCmd.Flags().String("t", "", "Tag for the record (optional)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
