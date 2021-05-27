package sync

import (
	"time"
)

func InsertPost(ts time.Time, hash, body, username string) {
	db := OpenTheDB()
	defer db.Close()
	tx, _ := db.Begin()

	s := `insert into posts (hash, body, username, created_at) values (?, ?, ?, ?)`
	thing, _ := tx.Prepare(s)
	thing.Exec(hash, body, username, ts)

	tx.Commit()
}
