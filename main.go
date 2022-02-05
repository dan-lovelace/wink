package main

import (
	"context"
	"os"

	"github.com/dan-lovelace/wink/commands"
	"github.com/dan-lovelace/wink/common"
)

func main() {
	w := &common.Wink{Context: context.Background()}
	resp := commands.Execute(w, os.Args[1:])
	if resp.Err != nil {
		os.Exit(-1)
	}
}
