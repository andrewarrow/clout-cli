package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"strconv"
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
func ListNotifications(argMap map[string]string) {

	limit, _ := strconv.Atoi(argMap["limit"])
	if limit == 0 {
		limit = 5
	}
	m := session.ReadAccounts()
	for username, s := range m {
		fmt.Println("")
		fmt.Println("===========")
		fmt.Println(username)
		fmt.Println("===========")
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		ListNotificationForPub(limit, pub58)
	}
}
func ListNotificationForPub(limit int, pub58 string) {
	//b, _ := ioutil.ReadFile("samples/get_notifications.list")
	js := network.GetNotifications(pub58)
	//ioutil.WriteFile("samples/get_notifications.list", []byte(js), 0755)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)

	for i, n := range list.Notifications {
		username := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check].Username
		fmt.Printf("  %02d %s %s\n", i, display.LeftAligned(n.Metadata.TxnType, 30),
			username)

		if n.Metadata.TxnType == "SUBMIT_POST" {
			parent := list.PostsByHash[n.Metadata.SubmitPostTxindexMetadata.ParentPostHashHex]
			p := list.PostsByHash[n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex]
			ts := time.Unix(p.TimestampNanos/1000000000, 0)
			ago := timeago.FromDuration(time.Since(ts))
			//fmt.Println(n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex)
			tokens := strings.Split(parent.Body, "\n")
			fmt.Println(display.LeftAligned(tokens[0], 60))
			tokens = strings.Split(p.Body, "\n")
			fmt.Println(display.LeftAligned("  "+tokens[0], 60), "     ", ago)
		} else if n.Metadata.TxnType == "LIKE" {
			p := list.PostsByHash[n.Metadata.LikeTxindexMetadata.PostHashHex]
			DisplayPostForNotification("LIKE", p)
		} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
			md := n.Metadata.CreatorCoinTransferTxindexMetadata
			//username := md.CreatorUsername
			//level := md.DiamondLevel
			if md.PostHashHex != "" {
				p := list.PostsByHash[md.PostHashHex]
				DisplayPostForNotification("DIAMOND", p)
			} else {
				fmt.Println("CREATOR_COIN_TRANSFER TODO")
			}
		} else if n.Metadata.TxnType == "CREATOR_COIN" {
			cctm := n.Metadata.CreatorCoinTxindexMetadata
			fmt.Printf("%s %0.2f %d %d\n", cctm.OperationType,
				float64(cctm.BitCloutToSellNanos)/1000000.0,
				cctm.CreatorCoinToSellNanos, cctm.BitCloutToAddNanos)
		}
		if i > limit {
			break
		}
	}
}

func DisplayPostForNotification(flavor string, p models.Post) {
	ts := time.Unix(p.TimestampNanos/1000000000, 0)
	ago := timeago.FromDuration(time.Since(ts))
	tokens := strings.Split(p.Body, "\n")
	fmt.Println(flavor, display.LeftAligned(tokens[0], 60), "     ", ago)
}
