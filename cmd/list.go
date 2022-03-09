/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"todolist/todo"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks to do",
	Long:  `List helps you to view all the tasks that has been been added to your application`,
	Run:   ListRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

}

func ListRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems()
	if err != nil {
		log.Printf("%v", err)

	}
	//fmt.Println(items)
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
	for _, i := range items {
		if i.Done == false {
			fmt.Fprintln(w, strconv.Itoa(i.Position)+"\t"+i.Text+"\t")
		}
	}
	w.Flush()
}
