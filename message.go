package main

import (
	"clout/keys"
	"clout/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/justincampbell/timeago"
)

func ListMessages() {
	m := ReadAccounts()
	for username, s := range m {
		fmt.Println(username)
		pub58, _ := keys.ComputeKeysFromSeed(SeedBytes(s))
		ListMessagesForPub(pub58)
	}
}
func ListMessagesForPub(pub58 string) {
	js := GetMessagesStateless(pub58)
	var list models.MessageList
	json.Unmarshal([]byte(js), &list)
	for _, oc := range list.OrderedContactsWithMessages {
		for _, m := range oc.Messages {
			ts := time.Unix(m.TstampNanos/1000000000, 0)
			ago := timeago.FromDuration(time.Since(ts))
			fmt.Println("  ", ago)
		}
	}
}
