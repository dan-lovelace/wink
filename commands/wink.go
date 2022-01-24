package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Response struct {
	Err error
}

func Execute(args []string) Response {
	fmt.Println("Executed with args", args)
	var resp Response

	rootCmd := &cobra.Command{
		Use:   "wink",
		Short: "Wink is a simple time tracker",
		Long:  "An easy to use time tracking utility built with ♥ in Go",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(projectCommand())

	err := rootCmd.Execute()
	resp.Err = err

	return resp
}
