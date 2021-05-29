package sync

import (
	"time"
)

func InsertPost(reclouts int64, ts time.Time, hash, body, username string) {
	db := OpenTheDB()
	defer db.Close()
	tx, _ := db.Begin()

	s := `insert into posts (reclouts, hash, body, username, created_at) values (?, ?, ?, ?, ?)`
	thing, _ := tx.Prepare(s)
	thing.Exec(reclouts, hash, body, username, ts)

	tx.Commit()
}
func InsertUser(hash, username string) {
	db := OpenTheDB()
	defer db.Close()
	tx, _ := db.Begin()

	s := `insert into users (hash, username) values (?, ?)`
	thing, _ := tx.Prepare(s)
	thing.Exec(hash, username)

	tx.Commit()
}
