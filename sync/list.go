package sync

import (
	"clout/display"
	"fmt"
	"strings"
	"time"

	"github.com/justincampbell/timeago"
)

func LastHash() string {
	db := OpenTheDB()
	defer db.Close()
	rows, err := db.Query("select hash from posts order by created_at limit 1")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		var s1 string
		rows.Scan(&s1)
		return s1
	}
	return ""
}
func FindPosts(s string) {
	db := OpenTheDB()
	defer db.Close()
	rows, err := db.Query("select username, body, created_at from posts where body like '%" + s + "%")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		var body string
		var ts time.Time
		rows.Scan(&username, &body, &ts)
		ago := timeago.FromDuration(time.Since(ts))
		fmt.Println(display.LeftAligned(username, 30), ago)
		tokens := strings.Split(body, "\n")
		for _, b := range tokens {
			fmt.Println("  " + display.LeftAligned(b, 60))
		}
	}
}
