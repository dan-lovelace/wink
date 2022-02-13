package commands

import (
	"errors"
	"log"
	"os"

	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Checks to see if Wink needs to be initialized
func CheckInit(w *common.Wink) CommandResponse {
	if _, err := os.Stat(w.Config.DB.Location); len(os.Args) > 1 && os.Args[1] != "init" && errors.Is(err, os.ErrNotExist) {
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
	// create env config file if not exists
	env := w.Config.Env
	if _, err := os.Stat(env.Path); os.IsNotExist(err) {
		if _, err := os.Create(env.Path); err != nil {
			log.Fatal(err)
		}
	}

	// setup viper
	viper.SetConfigFile(env.Path)
	viper.SetConfigType(env.Type)

	// get init directory
	wd, err := os.Getwd()
	if err != nil {
		return CommandResponse{Error: err}
	}

	// write init directory to env config
	viper.Set("INIT_PATH", wd)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.WriteConfig(); err != nil {
		log.Fatal(err)
	}

	// TODO: create a backup
	if err := db.RunMigrations(w, db.UpCmd); err != nil {
		return CommandResponse{Error: err}
	}

	return CommandResponse{}
}
