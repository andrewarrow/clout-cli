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
