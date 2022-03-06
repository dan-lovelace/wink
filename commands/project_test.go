package commands

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/configs"
	"github.com/dan-lovelace/wink/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"
)

var now = time.Now()
var pid = strconv.FormatInt(now.Unix(), 10)
var pName = fmt.Sprintf("test-%s", pid)
var b = bytes.NewBufferString("")
var w *common.Wink

func TestCreateProjectCommand(t *testing.T) {
	cmd := projectCommand(w)
	cmd.SetOut(b)

	cmd.SetArgs([]string{"create", pName})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	expect := fmt.Sprintf("Created %s", pName)
	if !strings.Contains(string(out), expect) {
		t.Fatalf("expected \"%s\" to include \"%s\"", string(out), expect)
	}
}

func TestGetProjectsCommand(t *testing.T) {
	cmd := projectCommand(w)
	cmd.SetOut(b)

	cmd.SetArgs([]string{"ls"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	expect := pName
	if !strings.Contains(string(out), pName) {
		t.Fatalf("expected \"%s\" to include \"%s\"", string(out), expect)
	}
}

func init() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg.DBLocation = viper.GetString(configs.TestDBLocation)
	w = &common.Wink{
		Config:  cfg,
		Context: context.Background(),
		Out:     os.Stdout,
	}

	if err := db.RunMigrations(w, db.UpCmd); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
