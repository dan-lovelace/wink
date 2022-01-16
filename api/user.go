package api

import (
	"context"
	"log"
	"time"

	winkDB "github.com/dan-lovelace/wink/db"
)

type User struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool

	Username string
}

func CreateUser(ctx context.Context, username string) (int, error) {
	db := winkDB.GetDB(ctx)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(username) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(username)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(lastId), nil
}
