package main

import (
	"clout/session"
	"fmt"
	"time"
)

func HandleBulk() {

	query := argMap["query"]
	if query == "" {
		return
	}

	for _, username := range session.GetAccountsForTag(query) {
		fmt.Println(username)
		session.WriteSelected(username)
		//os.Args = []string{"", "follow", "changeme"}
		//HandleFollow()
		//os.Args = []string{"", "reclout", "changeme"}
		//HandleReclout()
		//m := map[string]string{"text": "we also like @derishaviar", "reply": "changeme"}
		//Post(m)
		//m := map[string]string{"hash": "changeme"}
		//HandleLike(m)
		time.Sleep(time.Second * 1)
	}

}
