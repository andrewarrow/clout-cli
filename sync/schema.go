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

	//percent
	//cap
	//number of holders
	//number of board members
	sqlStmt := `
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
