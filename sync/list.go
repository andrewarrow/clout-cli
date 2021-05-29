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
	rows, err := db.Query("select username, body, created_at from posts where body like '%" + s + "%'")
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
func FindUsers() {
	db := OpenTheDB()
	defer db.Close()
	rows, err := db.Query("select username from users order by username")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println("./clout follow", username)
		fmt.Println("sleep 1")
	}
}
func FindTopReclouted() {
	db := OpenTheDB()
	defer db.Close()
	sql := `SELECT sum(p.reclouts) as total, 
                 p.username, 
								 u.points,
								 u.market_cap,
								 u.num_hodl,
								 u.num_board
					FROM posts p, users u 
          WHERE p.username = u.username 
					GROUP BY p.username, u.points, u.market_cap, u.num_hodl, u.num_board
					ORDER BY total DESC limit 1000`

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	fields := []string{"reclouts", "username", "points", "cap", "holders", "board"}
	sizes := []int{9, 20, 7, 10, 10, 10}
	display.Header(sizes, fields...)
	for rows.Next() {
		var total string
		var username string
		var points string
		var marketCap string
		var numHodl string
		var numBoard string
		rows.Scan(&total, &username, &points, &marketCap, &numHodl, &numBoard)
		display.Row(sizes, total, username, points, marketCap, numHodl, numBoard)
	}
}
