package main

import (
	"context"
	"log"
	"os"

	"github.com/dan-lovelace/wink/commands"
	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/configs"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	w := &common.Wink{
		Config:  cfg,
		Context: context.Background(),
		Out:     os.Stdout,
	}

	if initResp := commands.CheckInit(w); initResp.Error != nil {
		log.Fatal(initResp.Error)
	}

	resp := commands.Execute(w, os.Args[1:])
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
}
