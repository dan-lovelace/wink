package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/dan-lovelace/wink/configs"
)

func GetDB(ctx context.Context) *sql.DB {
	db, err := sql.Open(configs.DBConn.Driver, configs.DBConn.Location)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
