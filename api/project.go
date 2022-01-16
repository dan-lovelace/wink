package api

import (
	"context"
	"time"

	db "github.com/dan-lovelace/wink/db"
)

type Project struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool

	UserId int
	Name   string
}

func CreateProject(ctx context.Context, name string) {
	mDB := db.GetDB(ctx)
	defer mDB.Close()
}
