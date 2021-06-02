package main

import (
	"clout/display"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"time"
)

func HandleGlobal() {
	pub58 := session.LoggedInPub58()
	js := network.GetPostsStateless(pub58, false)
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	for _, p := range ps.PostsFound {
		username := p.ProfileEntryResponse.Username
		fmt.Println(username)
		pub58 := p.ProfileEntryResponse.PublicKeyBase58Check
		GetNotificationsForEachGlobalPost(pub58)
	}
}

func GetNotificationsForEachGlobalPost(pub58 string) {
	offset := -1
	for {
		fmt.Println("offset", offset)
		js := network.GetNotificationsWithOffset(offset, pub58)
		var list models.NotificationList
		json.Unmarshal([]byte(js), &list)
		if len(list.Notifications) == 0 || offset > 200 {
			break
		}
		for _, n := range list.Notifications {
			username := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check].Username
			if n.Metadata.TxnType == "SUBMIT_POST" {
			} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
			} else if n.Metadata.TxnType == "CREATOR_COIN" {
				cctm := n.Metadata.CreatorCoinTxindexMetadata

				if display.OneE9Float(cctm.BitCloutToSellNanos) >= 1.0 {
					fmt.Println(" ", username, cctm.OperationType, display.OneE9(cctm.BitCloutToSellNanos))
				}
			}
		}
		time.Sleep(time.Second * 1)
		offset += 50
	}
}
