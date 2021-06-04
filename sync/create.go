package sync

import (
	"fmt"
	"time"
)

func InsertNotification(to, from, flavor, meta, hash string) {
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
func InsertUser(marketCap string, numUsersYouHODL int,
	numBoardMembers int, points int64, hash, username string) {
	db := OpenTheDB()
	defer db.Close()
	tx, _ := db.Begin()

	s := `insert into users (market_cap, num_hodl, num_board, points, hash, username) values (?, ?, ?, ?, ?, ?)`
	thing, _ := tx.Prepare(s)
	thing.Exec(marketCap, numUsersYouHODL, numBoardMembers, points, hash, username)

	tx.Commit()
}
