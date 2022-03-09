/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"todolist/todo"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up done tasks",
	Run:   cleanedup,
}

func init() {
	rootCmd.AddCommand(cleanupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func cleanedup(cmd *cobra.Command, args []string) {
	//type Item struct {
	//	Position int
	//	Text     string
	//	Done     bool
	//	Clean    bool
	//}
	//open new slice of Items
	var newSlice []todo.Item
	items, err := todo.ReadItems()

	if err != nil {
		log.Fatal(err)
	}
	for _, val := range items {
		if val.Done == true {
			continue
		} else {
			newSlice = append(newSlice, val)
		}
	}
	os.OpenFile("mytodolist.csv", os.O_TRUNC, 0644)
	todo.SaveItems(newSlice)
	fmt.Println("Cleaned ")

	//if i > 0 && i <= len(items) {
	//	items[i-1].Clean = true
	//	todo.SaveItems(items)
	//}
	//fmt.Printf("%q %v\n", items[i-1].Text, "Cleaned")
}
