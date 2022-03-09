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
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	var items []todo.Item

	itemsArr, err := todo.ReadItems()
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range itemsArr {
		items = append(items, todo.Item{1, x.Text, false})
	}
	items = append(items, todo.Item{1, args[0], false})

	err = todo.SaveItems(items)
	if err != nil {
		fmt.Errorf("%v", err)
	}

}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
