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
	b := bytes.NewBufferString("")
	w := &common.Wink{
		Config: configs.Config{
			DB: &configs.DBConn{
				Driver:   "sqlite3",
				Location: "../testing.db",
			},
		},
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

	w := &common.Wink{
		Config: configs.Config{
			DB: &configs.DBConn{
				Driver:   "sqlite3",
				Location: "../testing.db",
			},
		},
		Context: context.Background(),
		// Out:     b,
	}

	initialize(w)
}
