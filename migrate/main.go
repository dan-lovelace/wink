package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/configs"
	"github.com/dan-lovelace/wink/db"
)

func main() {
	w := &common.Wink{
		Config:  configs.GetAppConfig(),
		Context: context.Background(),
	}

	flag.NewFlagSet(db.UpCmd, flag.ExitOnError)
	flag.NewFlagSet(db.DownCmd, flag.ExitOnError)

	firstArg := os.Args[1]
	switch firstArg {
	case db.UpCmd:
		fallthrough
	case db.DownCmd:
		db.RunMigrations(w, firstArg)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
