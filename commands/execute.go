package commands

import (
	"github.com/dan-lovelace/wink/common"
	"github.com/spf13/cobra"
)

type CommandResponse struct {
	Data  interface{}
	Error error
}

func Execute(w *common.Wink, args []string) CommandResponse {
	var resp CommandResponse

	rootCmd := &cobra.Command{
		Use:   "wink",
		Short: "Wink is a simple time tracker",
		Long:  "An easy to use time tracking utility built with ♥ in Go",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(initializeCommand(w))
	rootCmd.AddCommand(projectCommand(w))

	err := rootCmd.Execute()
	resp.Error = err

	return resp
}
