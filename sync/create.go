package sync

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func InsertNotification(tx *sql.Tx, to, from, flavor, meta, hash, coin string, amount int64) bool {

	s := `insert into notifications (to_user, flavor, from_user, hash, meta, coin, amount, created_at) values (?, ?, ?, ?, ?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println("1", e)
	}
	ts := time.Now()
	_, e = thing.Exec(to, flavor, from, hash, meta, coin, amount, ts)
	if e != nil {
		if strings.HasPrefix(e.Error(), "UNIQUE constraint failed") {
			return false
		}
		fmt.Println("2", e)
	}

	return true
}
func InsertPost(parent string, reclouts int64, ts time.Time, hash, body, username string) {
	db := OpenTheDB()
	defer db.Close()
	tx, _ := db.Begin()

	s := `insert into posts (parent, reclouts, hash, body, username, created_at) values (?, ?, ?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(parent, reclouts, hash, body, username, ts)
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
func InsertUser(username, pub58 string) bool {
	db := OpenTheDB()
	defer db.Close()
	tx, _ := db.Begin()

	s := `insert into users (username, pub58, created_at, updated_at) values (?, ?, ?, ?)`
	thing, _ := tx.Prepare(s)
	_, e := thing.Exec(username, pub58, time.Now(), time.Now())
	if e != nil {
		if strings.HasPrefix(e.Error(), "UNIQUE constraint failed") {
			return false
		}
	}

	tx.Commit()
	return true
}
