package sync

import (
	"clout/files"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenTheDB() *sql.DB {
	db, err := sql.Open("sqlite3", files.UserHomeDir()+"/clout-cli-data/sync.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func CreateSchema() {
	db := OpenTheDB()
	defer db.Close()

	/*
		CREATOR_COIN_TRANSFER     BeActive             [1] everyone? even m 5fb5823
		CREATOR_COIN_TRANSFER     BeActive             0.00 BeActive
		FOLLOW                    Lianna_jigger
		CREATOR_COIN              Clouterry            [BUY] 0.00
		CREATOR_COIN_TRANSFER     BeActive             [1]                  698712f
		CREATOR_COIN_TRANSFER     BeActive             0.00 BeActive
		LIKE                      I_LOVE_BITCLOUT      we are trying to hel 95cba16
		LIKE                      I_LOVE_BITCLOUT      we also like @derish 6d6f293
		SUBMIT_POST               DerishaViar          we also like @derish 6d6f293
	*/
	sqlStmt := `
create table notifications (to text, flavor text, from text, hash text, amount text, meta text, created_at datetime);

CREATE UNIQUE INDEX notifications_hash_idx
  ON notifications (hash);

create table posts (parent text, reclouts integer, hash text, body text, username text, created_at datetime);

CREATE UNIQUE INDEX posts_hash_idx
  ON posts (hash);

CREATE INDEX posts_username_idx
  ON posts (username);

CREATE INDEX posts_parent_idx
  ON posts (parent);

create table users (market_cap text, num_hodl integer, num_board integer, points integer, hash text, username text, created_at datetime);

CREATE UNIQUE INDEX users_idx
  ON users (username);
`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		//fmt.Printf("%q\n", err)
		return
	}
}
