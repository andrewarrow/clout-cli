package sync

import (
	"time"
)

func InsertPost(hash, body, username string) {
	ts := time.Now()
	db := OpenTheDB()
	defer db.Close()
	tx, _ := db.Begin()

	s := `insert into posts (hash, body, username, created_at) values (?, ?, ?, ?)`
	thing, _ := tx.Prepare(s)
	thing.Exec(hash, body, username, ts)

	tx.Commit()
}
