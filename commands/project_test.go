package commands

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/configs"
)

func TestGetProjectsCommand(t *testing.T) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// TODO
	cfg.DBLocation = ""

	b := bytes.NewBufferString("")
	w := &common.Wink{
		Config:  cfg,
		Context: context.Background(),
		Out:     b,
	}

	cmd := projectCommand(w)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"ls"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != "test" {
		t.Fatalf("expected \"%s\" got \"%s\"", "test", string(out))
	}
}

func init() {
	fmt.Println("initting...")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("init path", path)
	dir := filepath.Dir(path)
	fmt.Println("init dir", dir)
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	w := &common.Wink{
		Config:  cfg,
		Context: context.Background(),
		Out:     os.Stdout,
	}

	initialize(w)
}
