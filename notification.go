package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/justincampbell/timeago"
)

func ParseUserList(js string, buff []string) map[string]string {
	m := map[string]string{}
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	for _, u := range us.UserList {
		m[u.ProfileEntryResponse.PublicKeyBase58Check] = u.ProfileEntryResponse.Username
	}
	for _, b := range buff {
		if m[b] == "" {
			m[b] = "anonymous"
		}
	}
	return m
}
func ListNotifications() {
	m := ReadAccounts()
	for username, s := range m {
		fmt.Println(username)
		pub58, _ := keys.ComputeKeysFromSeed(SeedBytes(s))
		ListNotificationForPub(pub58)
	}
}
func ListNotificationForPub(pub58 string) {
	//b, _ := ioutil.ReadFile("samples/get_notifications.list")
	js := GetNotifications(pub58)
	//ioutil.WriteFile("samples/get_notifications.list", []byte(js), 0755)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	mapOfUsers := map[string]bool{}
	buff := []string{}
	for _, n := range list.Notifications {
		mapOfUsers[n.Metadata.TransactorPublicKeyBase58Check] = true
	}
	cache := ReadCache()
	for pub58, _ := range mapOfUsers {
		if cache[pub58] != "" {
			delete(mapOfUsers, pub58)
		}
	}
	for k, _ := range mapOfUsers {
		buff = append(buff, k)
		if len(buff) == 10 {
			fmt.Println("fetching 10...")
			js = GetManyUsersStateless(buff)
			m := ParseUserList(js, buff)
			for k, v := range m {
				cache[k] = v
			}
			buff = []string{}
		}
	}
	if len(buff) > 0 {
		fmt.Println("fetching ", len(buff))
		js = GetManyUsersStateless(buff)
		m := ParseUserList(js, buff)
		for k, v := range m {
			cache[k] = v
		}
	}
	WriteCache(cache)

	for i, n := range list.Notifications {
		fmt.Printf("  %02d %s %s %s\n", i, display.LeftAligned(n.Metadata.TxnType, 30),
			cache[n.Metadata.TransactorPublicKeyBase58Check],
			n.Metadata.CreatorCoinTransferTxindexMetadata.CreatorUsername)

		if n.Metadata.TxnType == "SUBMIT_POST" {
			p := list.PostsByHash[n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex]
			ts := time.Unix(p.TimestampNanos/1000000000, 0)
			ago := timeago.FromDuration(time.Since(ts))
			//fmt.Println(n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex)
			tokens := strings.Split(p.Body, "\n")
			fmt.Println(display.LeftAligned(tokens[0], 60), "     ", ago)
		} else if n.Metadata.TxnType == "CREATOR_COIN" {
			cctm := n.Metadata.CreatorCoinTxindexMetadata
			fmt.Printf("%s %0.2f %d %d\n", cctm.OperationType,
				float64(cctm.BitCloutToSellNanos)/1000000.0,
				cctm.CreatorCoinToSellNanos, cctm.BitCloutToAddNanos)
		}
		if i > 5 {
			break
		}
	}
}
