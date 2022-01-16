package api

import (
	"context"
	"log"

	winkDB "github.com/dan-lovelace/wink/db"
)

func StartNewTimer(ctx context.Context) {
	db := winkDB.GetDB(ctx)
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
