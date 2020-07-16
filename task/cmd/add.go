package cmd

import (
		"fmt"
		"github.com/spf13/cobra"
		"strings"
)

var addCmd = &cobra.Command{
		Use: "add",
		Short: "Adds a task to your task list.",
		Run: func(command *cobra.Command, args []string) {
				task := strings.Join(args, " ")
				fmt.Printf("The task \"%s\" was added to your list. \n", task)
		},
}

func init() {
		RootCmd.AddCommand(addCmd)
}