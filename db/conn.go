package db

import (
	"database/sql"
	"log"

	"github.com/dan-lovelace/wink/common"
	_ "github.com/mattn/go-sqlite3"
)

func GetDB(w *common.Wink) *sql.DB {
	db, err := sql.Open(w.Config.DBDriver, w.Config.DBLocation)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
