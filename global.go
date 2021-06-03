package main

import (
	"clout/display"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"strings"
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
		for k, v := range buyers {
			fmt.Println(" ", k, v)
		}
	}
}

func GetNotificationsForEachGlobalPost(coin, pub58 string) map[string]string {
	offset := -1
	buyers := map[string]string{}
	for {
		//fmt.Println("offset", offset)
		js := network.GetNotificationsWithOffset(offset, pub58)
		var list models.NotificationList
		//ioutil.WriteFile(fmt.Sprintf("%s_%d.json", coin, offset), []byte(js), 0755)
		json.Unmarshal([]byte(js), &list)
		if len(list.Notifications) == 0 || offset > 200 {
			break
		}
		for _, n := range list.Notifications {
			buyer := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check].Username
			if buyer == "" || strings.HasPrefix(buyer, "B1") {
				continue
			}
			if n.Metadata.TxnType == "SUBMIT_POST" {
			} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
			} else if n.Metadata.TxnType == "CREATOR_COIN" {
				cctm := n.Metadata.CreatorCoinTxindexMetadata

				if display.OneE9Float(cctm.BitCloutToSellNanos) >= 1.0 {
					//display.Row(sizes, username, target, display.OneE9(cctm.BitCloutToSellNanos))
					buyers[buyer] = display.OneE9(cctm.BitCloutToSellNanos)
				}
			}
		}
		time.Sleep(time.Second * 1)
		offset += 50
	}
	return buyers
}
