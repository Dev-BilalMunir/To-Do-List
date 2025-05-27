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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [todo ID]",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		todoList, _ := todo.LoadTodo()
		for i, t := range todoList {
			if t.ID == id {
				todoList[i].Status = true
				todo.SaveTodo(todoList)
				fmt.Println("✅ Task marked as done:", t.Text)
				return
			}
		}
		fmt.Println("⚠️ Task not found")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
