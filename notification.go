package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"clout/sync"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var GlobalListNeverAutoFollow = map[string]bool{"andrewarrow": true,
	"cloutcli":     true,
	"yournamehere": true}

func HandleNotifications(argMap map[string]string) {
	query := argMap["query"]
	autofollow := argMap["autofollow"]

	sync.CreateSchema()
	sorted := session.ReadAccountsSorted()
	m := session.ReadAccounts()
	db := sync.OpenTheDB()
	defer db.Close()

	if query != "" {
		sorted = session.GetAccountsForTag(query)
	}
	for _, to := range sorted {
		fmt.Println("to", to)
		session.WriteSelected(to)
		s := m[to]
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		NotificationsForSyncUser(autofollow, db, to, pub58)
		time.Sleep(time.Second * 2)
	}
}
func NotificationsForSyncUser(autofollow string, db *sql.DB, to, pub58 string) {
	js := network.GetNotifications(pub58)
	//ioutil.WriteFile("foo.json", []byte(js), 0755)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	tx, _ := db.Begin()
	followed := map[string]bool{}
	for _, n := range list.Notifications {
		if n.Metadata.TxnType == "BASIC_TRANSFER" {
			continue
		}
		from := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check].Username
		if from == "" {
			from = "anonymous"
		}
		hash := ""
		phh := ""
		meta := ""
		flavor := ""
		coin := ""
		amount := int64(0)
		if n.Metadata.TxnType == "SUBMIT_POST" {
			p := list.PostsByHash[n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex]
			if p.Body == "" {
				phh = p.RecloutedPostEntryResponse.PostHashHex
				hash = fmt.Sprintf("%s_%s_reclout_%s", from, to, phh)
				flavor = "reclout"
				meta = BodyParse(p.RecloutedPostEntryResponse.Body)
			} else {
				phh = p.PostHashHex
				hash = fmt.Sprintf("%s_%s_mention_%s", from, to, phh)
				flavor = "mention"
				meta = BodyParse(p.Body)
			}
		} else if n.Metadata.TxnType == "LIKE" {
			p := list.PostsByHash[n.Metadata.LikeTxindexMetadata.PostHashHex]
			phh = p.PostHashHex
			hash = fmt.Sprintf("%s_%s_like_%s", from, to, phh)
			meta = BodyParse(p.Body)
			flavor = "like"
		} else if n.Metadata.TxnType == "FOLLOW" {
			hash = fmt.Sprintf("%s_%s", from, to)
			meta = from
			flavor = "follow"
		} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
			md := n.Metadata.CreatorCoinTransferTxindexMetadata
			if md.PostHashHex != "" {
				p := list.PostsByHash[md.PostHashHex]
				phh = p.PostHashHex
				hash = fmt.Sprintf("%s_%s_%s_d_%d", from, to, phh, md.DiamondLevel)
				meta = fmt.Sprintf("%d ", md.DiamondLevel) + BodyParse(p.Body)
				amount = md.DiamondLevel
				flavor = "diamond"
			} else {
				hash = fmt.Sprintf("%s_%s_tx_%s_%d_%s", from, to, md.CreatorUsername, md.CreatorCoinToTransferNanos, n.Metadata.BlockHashHex)
				meta = fmt.Sprintf("%s %d", md.CreatorUsername, md.CreatorCoinToTransferNanos)
				amount = md.CreatorCoinToTransferNanos
				coin = md.CreatorUsername
				flavor = "coin"
			}
		} else if n.Metadata.TxnType == "CREATOR_COIN" {
			cctm := n.Metadata.CreatorCoinTxindexMetadata
			if cctm.OperationType == "buy" {
				amount = cctm.BitCloutToSellNanos
			} else if cctm.OperationType == "sell" {
				amount = cctm.CreatorCoinToSellNanos
			}
			hash = fmt.Sprintf("%s_%s_%s_%d_%s", from, to, cctm.OperationType, amount, n.Metadata.BlockHashHex)
			meta = fmt.Sprintf("%s %d", from, amount)
			flavor = cctm.OperationType
		}
		ok := sync.InsertNotification(tx, to, from, flavor, meta, hash, coin, amount)
		if ok {
			fmt.Printf("%s %s %s %s\n", display.LeftAligned(flavor, 10),
				display.LeftAligned(from, 20),
				display.LeftAligned(coin, 20),
				display.LeftAligned(amount, 10))
			fmt.Printf("\n%s\n%s\n", meta, phh)
			if autofollow == "true" && followed[from] == false && GlobalListNeverAutoFollow[to] == false {
				os.Args = []string{"", "follow", from}
				followed[from] = true
				fmt.Println("--------", to, from)
				HandleFollow()
			}
		}
	}
	e := tx.Commit()
	if e != nil {
		fmt.Println("3", e)
	}
}

func BodyParse(body string) string {
	tokens := strings.Split(body, "\n")
	return tokens[0]
}
