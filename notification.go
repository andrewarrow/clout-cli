package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/justincampbell/timeago"
)

func HandleNotifications(argMap map[string]string) {
	if len(os.Args) > 2 {
		username := os.Args[2]

		if username == "sync" {
			FillUpLocalDatabaseWithNotifications()
			return
		}
		//m := session.ReadAccounts()
		//pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(m[username]))
		pub58 := session.UsernameToPub58(username)
		ListNotificationsForUser(pub58)
		return
	}
	ListNotifications(argMap)
}
func NotificationsForSyncUser(to, pub58 string) {
	js := network.GetNotifications(pub58)
	//ioutil.WriteFile("foo.json", []byte(js), 0755)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	for _, n := range list.Notifications {
		from := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check].Username
		flavor := n.Metadata.TxnType
		hash := ""
		if n.Metadata.TxnType == "SUBMIT_POST" {
			p := list.PostsByHash[n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex]
			if p.Body == "" {
				phh := p.RecloutedPostEntryResponse.PostHashHex
				hash = fmt.Sprintf("%s_%s_reclout_%s", from, to, phh)
			} else {
				phh := n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex
				hash = fmt.Sprintf("%s_%s_mention_%s", from, to, phh)
			}
		} else if n.Metadata.TxnType == "LIKE" {
			phh := n.Metadata.LikeTxindexMetadata.PostHashHex
			hash = fmt.Sprintf("%s_%s_like_%s", from, to, phh)
		} else if n.Metadata.TxnType == "FOLLOW" {
			hash = fmt.Sprintf("%s_%s", from, to)
		} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
			md := n.Metadata.CreatorCoinTransferTxindexMetadata
			if md.PostHashHex != "" {
				hash = fmt.Sprintf("%s_%s_%s_d_%d", from, to, md.PostHashHex, md.DiamondLevel)
			} else {
				hash = fmt.Sprintf("%s_%s_tx_%s_%d", from, to, md.CreatorUsername, md.CreatorCoinToTransferNanos)
			}
		} else if n.Metadata.TxnType == "CREATOR_COIN" {
			cctm := n.Metadata.CreatorCoinTxindexMetadata
			amount := int64(0)
			if cctm.OperationType == "buy" {
				amount = cctm.BitCloutToSellNanos
			} else if cctm.OperationType == "sell" {
				amount = cctm.CreatorCoinToSellNanos
			}
			hash = fmt.Sprintf("%s_%s_%s_%d", from, to, cctm.OperationType, amount)
		}
		fmt.Println(" ", "from", display.LeftAligned(from, 20),
			display.LeftAligned(flavor, 30), hash)
	}
}
func FillUpLocalDatabaseWithNotifications() {
	sorted := session.ReadAccountsSorted()
	m := session.ReadAccounts()
	for _, to := range sorted {
		fmt.Println("to", to)
		s := m[to]
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		NotificationsForSyncUser(to, pub58)
		time.Sleep(time.Second * 1)
	}
}

func ListNotifications(argMap map[string]string) {

	sorted := session.ReadAccountsSorted()
	m := session.ReadAccounts()
	fields := []string{"username", "follows", "likes", "posts", "coin", "coin_tx"}
	sizes := []int{20, 9, 9, 9, 9, 9}
	display.Header(sizes, fields...)

	baseline := session.ReadBaseline()
	save := map[string]map[string]int{}
	for _, username := range sorted {
		s := m[username]
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		m := NotificationForPub(pub58)
		save[username] = m
		deltaFollows := m["follows"] - baseline[username]["follows"]
		deltaLikes := m["likes"] - baseline[username]["likes"]
		deltaPosts := m["posts"] - baseline[username]["posts"]
		deltaCoin := m["coin"] - baseline[username]["coin"]
		deltaCoinTx := m["coin_tx"] - baseline[username]["coin_tx"]
		if deltaFollows == 0 && deltaLikes == 0 &&
			deltaPosts == 0 && deltaCoin == 0 &&
			deltaCoinTx == 0 {
			continue
		}
		display.Row(sizes, username, deltaFollows, deltaLikes,
			deltaPosts, deltaCoin, deltaCoinTx)
	}
	session.SaveBaselineNotifications(save)
}
func NotificationForPub(pub58 string) map[string]int {
	js := network.GetNotifications(pub58)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)

	m := map[string]int{}
	follows := 0
	likes := 0
	posts := 0
	coin := 0
	coinTx := 0
	for _, n := range list.Notifications {
		if n.Metadata.TxnType == "FOLLOW" {
			follows++
		} else if n.Metadata.TxnType == "LIKE" {
			likes++
		} else if n.Metadata.TxnType == "SUBMIT_POST" {
			posts++
		} else if n.Metadata.TxnType == "CREATOR_COIN" {
			coin++
		} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
			coinTx++
		} else if n.Metadata.TxnType == "BASIC_TRANSFER" {
		} else {
			fmt.Println(n.Metadata.TxnType)
		}
	}
	m["follows"] = follows
	m["likes"] = likes
	m["posts"] = posts
	m["coin"] = coin
	m["coin_tx"] = coinTx
	return m
}

func Old(limit int, pub58 string) {
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

func BodyParse(body string) string {
	tokens := strings.Split(body, "\n")
	return tokens[0]
}
func DisplayPostForNotification(flavor string, p models.Post) {
	ts := time.Unix(p.TimestampNanos/1000000000, 0)
	ago := timeago.FromDuration(time.Since(ts))
	tokens := strings.Split(p.Body, "\n")
	fmt.Println(flavor, display.LeftAligned(tokens[0], 60), "     ", ago)
}
func ListNotificationsForUser(pub58 string) {

	js := network.GetNotifications(pub58)
	//ioutil.WriteFile("foo.json", []byte(js), 0755)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	fields := []string{"flavor", "username", "meta", "hash"}
	sizes := []int{25, 20, 20, 10}
	display.Header(sizes, fields...)
	shortMap := map[string]string{}
	for _, n := range list.Notifications {
		username := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check].Username
		meta := ""
		short := ""
		if n.Metadata.TxnType == "SUBMIT_POST" {
			//parent := list.PostsByHash[n.Metadata.SubmitPostTxindexMetadata.ParentPostHashHex]
			p := list.PostsByHash[n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex]
			short = p.PostHashHex[0:7]
			shortMap[short] = p.PostHashHex
			meta = BodyParse(p.Body)
			if meta == "" {
				meta = BodyParse(p.RecloutedPostEntryResponse.Body)
				short = p.RecloutedPostEntryResponse.PostHashHex[0:7]
				shortMap[short] = p.PostHashHex
			}
		} else if n.Metadata.TxnType == "LIKE" {
			p := list.PostsByHash[n.Metadata.LikeTxindexMetadata.PostHashHex]
			short = p.PostHashHex[0:7]
			shortMap[short] = p.PostHashHex
			meta = BodyParse(p.Body)
		} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
			md := n.Metadata.CreatorCoinTransferTxindexMetadata
			if md.PostHashHex != "" {
				p := list.PostsByHash[md.PostHashHex]
				meta = fmt.Sprintf("[%d] %s", md.DiamondLevel, BodyParse(p.Body))
				short = p.PostHashHex[0:7]
				shortMap[short] = p.PostHashHex
			} else {
				meta = display.OneE9(md.CreatorCoinToTransferNanos) + " " +
					md.CreatorUsername
			}
		} else if n.Metadata.TxnType == "CREATOR_COIN" {
			cctm := n.Metadata.CreatorCoinTxindexMetadata
			if cctm.OperationType == "buy" {
				meta = fmt.Sprintf("[BUY] %s", display.OneE9(cctm.BitCloutToSellNanos))
			} else if cctm.OperationType == "sell" {
				meta = fmt.Sprintf("[SELL] %s", display.OneE9(cctm.CreatorCoinToSellNanos))
			}
			//cctm.BitCloutToAddNanos
		}

		display.Row(sizes, n.Metadata.TxnType, username, meta, short)
	}
	session.SaveShortMap(shortMap)
}
