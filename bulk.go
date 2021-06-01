package main

import (
	"clout/session"
	"fmt"
	"os"
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
		os.Args = []string{"", "follow", "changeme"}
		HandleFollow()
		//os.Args = []string{"", "reclout", "changeme"}
		//HandleReclout()
		time.Sleep(time.Second * 1)
	}

}
