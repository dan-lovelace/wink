package main

import (
	"os"

	"github.com/dan-lovelace/wink/commands"
)

func main() {
	resp := commands.Execute(os.Args[1:])
	if resp.Err != nil {
		os.Exit(-1)
	}
}
