/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/NureddinFarzaliyev/go-dict-cli/internal/search"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a word or a tag in dictionary.",
	Long: `Search for a word or for a tag in specific dictionary.

Example:
	Word search: dict search <dict-name> <word>
	Tag search: dict search <dict-name> <word> -t`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			WithDB(func(conn *sql.DB) (err error) {
				tag, _ := cmd.Flags().GetBool("t")
				return search.SearchWordInDictionary(conn, args[0], args[1], tag)
			})
		} else {
			fmt.Println(cmd.Long)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().BoolP("t", "t", false, "Use this flag if you want to search based on tag instead of word")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
