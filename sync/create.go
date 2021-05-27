package sync

import (
	"time"
)

func InsertPost(body, username string) {
	ts := time.Now()
	db := OpenTheDB()
	defer db.Close()
	tx, _ := db.Begin()

	s := `insert into posts (body, username, created_at) values (?, ?, ?)`
	thing, _ := tx.Prepare(s)
	thing.Exec(body, username, ts)

	tx.Commit()
}
