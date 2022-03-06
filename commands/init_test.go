package commands

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/configs"
	"github.com/dan-lovelace/wink/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"
)

func TestInitialize(t *testing.T) {

}

func init() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg.DBLocation = viper.GetString(configs.TestDBLocation)
	w = &common.Wink{
		Config:  cfg,
		Context: context.Background(),
		Out:     os.Stdout,
	}

	if err := db.RunMigrations(w, db.UpCmd); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
