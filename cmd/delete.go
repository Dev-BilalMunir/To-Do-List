/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	todo "github.com/godev/todo-list/internal"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [todo ID]",
	Short: "Delete a task",
	Long: `Delete a task from the to-do list by specifying its ID.`,
	Example: `todo delete 1`,
	Args: cobra.MinimumNArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		todoList, err := todo.LoadTodo()
		if err != nil {
			fmt.Println("Error loading todo list:", err)
			return
		}

		for _, arg := range args {
			id, _ := strconv.Atoi(arg)
			// delete from the todo list
			for i, t := range todoList {
				if t.ID == id {
					todoList = append(todoList[:i], todoList[i+1:]...)
				}
			}
			todo.SaveTodo(todoList)
		}
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
