package sync

import (
	"fmt"
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
