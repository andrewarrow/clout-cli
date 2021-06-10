package main

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"sort"
)

func FindBuysSellsAndTransfers() {
	pub58 := session.LoggedInPub58()
	js := network.GetPostsStateless(pub58, false)
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	for _, p := range ps.PostsFound {
		username := p.ProfileEntryResponse.Username
		fmt.Println("notifications for", username)
		pub58 := session.UsernameToPub58(username)
		js := network.GetNotifications(pub58)
		var list models.NotificationList
		json.Unmarshal([]byte(js), &list)
		m := map[string]int64{}
		for _, n := range list.Notifications {
			fromPub58 := n.Metadata.TransactorPublicKeyBase58Check
			if n.Metadata.TxnType == "CREATOR_COIN" {
				cctm := n.Metadata.CreatorCoinTxindexMetadata
				if cctm.OperationType != "buy" {
					continue
				}
				m[fromPub58] += cctm.BitCloutToSellNanos
			}
		}
		for fromPub58, sum := range m {
			user := session.Pub58ToUser(pub58)
			from := list.ProfilesByPublicKey[fromPub58].Username
			if from == "" {
				continue
			}
			total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos

			sort.SliceStable(user.UsersWhoHODLYou, func(i, j int) bool {
				return user.UsersWhoHODLYou[i].BalanceNanos >
					user.UsersWhoHODLYou[j].BalanceNanos
			})
			for _, friend := range user.UsersWhoHODLYou {

				if friend.ProfileEntryResponse.Username == from {

					per := float64(friend.BalanceNanos) / float64(total)
					if per >= 0.01 {
						perString := fmt.Sprintf("%0.2f", per*100)
						text := fmt.Sprintf("@%s spends %d to buy @%s and now owns %s%%", from, sum, username, perString)
						fmt.Println(text)
						//m := map[string]string{"text": text}
						//Post(m)
					}
				}
			}
		}

	}
}
