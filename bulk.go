package main

import (
	"clout/session"
	"fmt"
	"os"
)

func HandleBulk() {
	query := argMap["query"]
	if query == "" {
		return
	}

	for _, username := range session.GetAccountsForTag(query) {
		fmt.Println(username)
		session.WriteSelected(username)
		os.Args = []string{"", "follow", "oblige"}
		HandleFollow()
	}

}
