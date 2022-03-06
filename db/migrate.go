package db

import (
	"log"
	"path"

	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/configs"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

const (
	UpCmd   string = "up"
	DownCmd string = "down"
)

func RunMigrations(w *common.Wink, direction string) error {
	db := GetDB(w)
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open(path.Join(viper.GetString(configs.InitDir), "migrate", "migrations"))
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, w.Config.DBDriver, instance)
	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case UpCmd:
		if err := m.Up(); err != nil {
			return err
		}
	case DownCmd:
		if err := m.Down(); err != nil {
			return err
		}
	}

	return nil
}
