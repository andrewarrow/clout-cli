package main

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
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
	js := network.GetNotifications(pub58)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	for _, n := range list.Notifications {
		//username := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check].Username
		if n.Metadata.TxnType == "SUBMIT_POST" {
		} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
		} else if n.Metadata.TxnType == "CREATOR_COIN" {
			cctm := n.Metadata.CreatorCoinTxindexMetadata
			fmt.Println(" ", cctm.OperationType)
		}
		fmt.Println(" ", n.Metadata.TxnType)
	}
}
