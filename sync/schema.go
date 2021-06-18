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

	sqlStmt := `
create table notifications (to_user text, flavor text, from_user text, hash text, meta text, coin text, amount integer, created_at datetime);

CREATE UNIQUE INDEX notifications_hash_idx
  ON notifications (hash);
CREATE INDEX notifications_flavor_idx
  ON notifications (flavor);
CREATE INDEX notifications_to_idx
  ON notifications (to_user);
CREATE INDEX notifications_from_idx
  ON notifications (from_user);

create table posts (parent text, reclouts integer, hash text, body text, username text, created_at datetime);

CREATE UNIQUE INDEX posts_hash_idx
  ON posts (hash);

CREATE INDEX posts_username_idx
  ON posts (username);

CREATE INDEX posts_parent_idx
  ON posts (parent);

create table users (username text, pub58 text, created_at datetime, updated_at datetime);

CREATE UNIQUE INDEX users_idx
  ON users (pub58);
`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
