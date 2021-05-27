package sync

import (
	"clout/files"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenTheDB() *sql.DB {
	db, err := sql.Open("sqlite3", files.UserHomeDir()+"/wolf.db")
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
create table posts (body text, username text, created_at datetime);
`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		//fmt.Printf("%q\n", err)
		return
	}
}
