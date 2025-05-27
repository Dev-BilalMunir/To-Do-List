/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	todo "github.com/godev/todo-list/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [todo]",
	Short: "Add a new todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskText := args[0]
		taskList, _ := todo.LoadTodo()
		newTask := todo.Todo{
			ID:     len(taskList) + 1,
			Text:   taskText,
			Status: false,
		}
		taskList = append(taskList, newTask)
		todo.SaveTodo(taskList)
		fmt.Println("✅ Todo added:", taskText)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
