package main

import (
	"fmt"
	"os"

	"github.com/dan-lovelace/wink/api"
)

func main() {
	argsWithProg := os.Args
	fmt.Println(argsWithProg)
	api.StartNewTimer()
}
