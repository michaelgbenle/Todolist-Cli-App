/*
Copyright © 2022 Gbenle Michael <Michaelgbenle@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"todolist/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to the list",
	Long:  `add will create a new todo item to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		AddRun(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func AddRun(args string) string {
	var items []todo.Item
	items, err := todo.ReadItems()
	if err != nil {
		log.Printf("%v", err)
	}

	items = append(items, todo.Item{1, args, false})
	err = todo.SaveItems(items)
	if err != nil {
		log.Printf("%v", err)
	}
	fmt.Println("Added")
	return ""
}
