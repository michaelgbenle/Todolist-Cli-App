/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"todolist/todo"

	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark a task as not done",
	Run:   undoneRun,
}

func init() {
	rootCmd.AddCommand(undoneCmd)

}

func undoneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems()
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}
	if i > 0 && i <= len(items) {
		items[i-1].Done = false
		todo.SaveItems(items)
	}
	fmt.Printf("%q %v\n", items[i-1].Text, "Not done")
}
