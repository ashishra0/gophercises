package cmd

import (
	"fmt"
		"strconv"

		"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
				id, err := strconv.Atoi(arg)
				if err != nil {
						fmt.Println("Unable to parse: ", arg)
						fmt.Println("Please pass a valid argument")
				} else {
						ids = append(ids, id)
				}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
