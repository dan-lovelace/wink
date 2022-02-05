package commands

import (
	"github.com/dan-lovelace/wink/common"
	"github.com/spf13/cobra"
)

type Response struct {
	Err error
}

func Execute(w *common.Wink, args []string) Response {
	var resp Response

	rootCmd := &cobra.Command{
		Use:   "wink",
		Short: "Wink is a simple time tracker",
		Long:  "An easy to use time tracking utility built with ♥ in Go",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(projectCommand(w))

	err := rootCmd.Execute()
	resp.Err = err

	return resp
}
