package commands

import (
	"errors"
	"os"

	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/configs"
	"github.com/dan-lovelace/wink/db"
	"github.com/spf13/cobra"
)

// Checks to see if Wink needs to be initialized
func CheckInit(w *common.Wink) CommandResponse {
	if _, err := configs.LoadConfig(); err != nil {
		return CommandResponse{Error: err}
	}

	if _, err := os.Stat(w.Config.DBLocation); len(os.Args) > 1 && os.Args[1] != "init" && errors.Is(err, os.ErrNotExist) {
		return CommandResponse{Error: errors.New("Database does not exist. Did you forget to run init?")}
	}

	return CommandResponse{}
}

func initializeCommand(w *common.Wink) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initializes Wink for first time use",
		Run: func(cmd *cobra.Command, args []string) {
			initialize(w)
		},
	}
}

// Initializes Wink resources such as the database and initial configuration files
func initialize(w *common.Wink) CommandResponse {
	// TODO: create a backup
	if err := db.RunMigrations(w, db.UpCmd); err != nil {
		return CommandResponse{Error: err}
	}

	return CommandResponse{}
}
