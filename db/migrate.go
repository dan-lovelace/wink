package db

import (
	"fmt"
	"log"

	"github.com/dan-lovelace/wink/common"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
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

	fSrc, err := (&file.File{}).Open("./migrate/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, w.Config.DB.Driver, instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrating", direction)
	switch direction {
	case UpCmd:
		if err := m.Up(); err != nil {
			log.Fatal(err)
			return err
		}
	case DownCmd:
		if err := m.Down(); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
