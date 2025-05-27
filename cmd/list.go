/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	todo "github.com/godev/todo-list/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		todoList, _ := todo.LoadTodo()
		if len(todoList) == 0 {
			fmt.Println("📭 No tasks yet!")
			return
		}
		fmt.Println("📝 To-Do List:")
		for _, t := range todoList {
			status := "❌"
			if t.Status {
				status = "✅"
			}
			fmt.Printf("%d. %s [%s]\n", t.ID, t.Text, status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
