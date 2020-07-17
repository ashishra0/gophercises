package cmd

import (
		"fmt"
		"github.com/ashishra0/gophercises/task/db"
		"github.com/spf13/cobra"
		"strings"
)

var addCmd = &cobra.Command{
		Use: "add",
		Short: "Adds a task to your task list.",
		Run: func(command *cobra.Command, args []string) {
				task := strings.Join(args, " ")
				_, err := db.CreateTask(task)
				if err != nil {
						fmt.Println("Something went wrong: ", err.Error())
						return
				}
				fmt.Printf("The task \"%s\" was added to your list. \n", task)
		},
}

func init() {
		RootCmd.AddCommand(addCmd)
}