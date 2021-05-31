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
		m := session.ReadAccounts()
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(m[username]))
		ListNotificationsForUser(pub58)
		return
	}
	ListNotifications(argMap)
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
func ListNotifications(argMap map[string]string) {

	sorted := session.ReadAccountsSorted()
	m := session.ReadAccounts()
	fields := []string{"username", "follows", "likes", "posts", "coin", "coin_tx"}
	sizes := []int{20, 9, 9, 9, 9, 9}
	display.Header(sizes, fields...)
	for _, username := range sorted {
		s := m[username]
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		m := NotificationForPub(pub58)
		display.Row(sizes, username, m["follows"], m["likes"],
			m["posts"], m["coin"], m["coin_tx"])
	}
}
func NotificationForPub(pub58 string) map[string]interface{} {
	js := network.GetNotifications(pub58)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)

	m := map[string]interface{}{}
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
