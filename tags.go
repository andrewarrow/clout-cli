package main

import (
	"clout/session"
	"fmt"
)

func HandleTags() {

	query := argMap["query"]
	if query == "" {
		fmt.Println("--query=x")
		return
	}

	for _, username := range session.GetAccountsForTag(query) {
		fmt.Println(username)
	}
}
