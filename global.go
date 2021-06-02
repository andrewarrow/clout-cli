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

	m := map[string]string{}
	for _, p := range ps.PostsFound {
		coin := p.ProfileEntryResponse.Username
		pub58 := p.ProfileEntryResponse.PublicKeyBase58Check
		m[coin] = pub58
	}
	for coin, pub58 := range m {
		fmt.Println(coin)
		buyers := GetNotificationsForEachGlobalPost(coin, pub58)
		for k, _ := range buyers {
			fmt.Println(" ", k)
		}
	}
}

func GetNotificationsForEachGlobalPost(coin, pub58 string) map[string]bool {
	offset := -1
	buyers := map[string]bool{}
	for {
		//fmt.Println("offset", offset)
		js := network.GetNotificationsWithOffset(offset, pub58)
		var list models.NotificationList
		json.Unmarshal([]byte(js), &list)
		if len(list.Notifications) == 0 || offset > 200 {
			break
		}
		for _, n := range list.Notifications {
			buyer := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check].Username
			if buyer == "" {
				continue
			}
			if n.Metadata.TxnType == "SUBMIT_POST" {
			} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
			} else if n.Metadata.TxnType == "CREATOR_COIN" {
				cctm := n.Metadata.CreatorCoinTxindexMetadata

				if display.OneE9Float(cctm.BitCloutToSellNanos) >= 10.0 {
					//display.Row(sizes, username, target, display.OneE9(cctm.BitCloutToSellNanos))
					buyers[buyer] = true
				}
			}
		}
		time.Sleep(time.Second * 1)
		offset += 50
	}
	return buyers
}
