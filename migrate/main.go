package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dan-lovelace/wink/configs"
	winkDB "github.com/dan-lovelace/wink/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

const (
	upCmd   string = "up"
	downCmd string = "down"
)

func runMigrate(ctx context.Context, direction string) error {
	db := winkDB.GetDB(ctx)
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./migrate/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, configs.DBConn.Driver, instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrating", direction)
	switch direction {
	case upCmd:
		if err := m.Up(); err != nil {
			log.Fatal(err)
			return err
		}
	case downCmd:
		if err := m.Down(); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func main() {
	ctx := context.Background()

	flag.NewFlagSet(upCmd, flag.ExitOnError)
	flag.NewFlagSet(downCmd, flag.ExitOnError)

	firstArg := os.Args[1]
	switch firstArg {
	case upCmd:
		fallthrough
	case downCmd:
		runMigrate(ctx, firstArg)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
