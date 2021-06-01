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
	"sort"
	"time"
)

func HandleFollow() {
	if len(os.Args) < 3 {
		fmt.Println("missing username")
		return
	}
	username := os.Args[2]
	follower := session.LoggedInPub58()
	followed := session.UsernameToPub58(username)
	jsonString := network.CreateFollow(follower, followed)
	var tx models.TxReady
	json.Unmarshal([]byte(jsonString), &tx)
	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	_, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	jsonString = network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}

func HandleFollowing() {
	pub58 := session.LoggedInPub58()
	if len(os.Args) > 2 {
		pub58 = os.Args[2]
		LoopThruAllFollowing(pub58, "")
		return
	}
	RunFollowLogic(pub58, "")
}
func ListFollowers() {
	pub58, username, _, _ := session.LoggedInAs()
	RunFollowLogic(pub58, username)
}
func RunFollowLogic(pub58, username string) {

	items := LoopThruAllFollowing(pub58, username)
	//fmt.Println("NumFollowers", pktpe.NumFollowers)
	//fmt.Println("")
	fields := []string{"username", "cap", "price"}
	sizes := []int{20, 10, 10}
	display.Header(sizes, fields...)

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].CoinPriceBitCloutNanos < items[j].CoinPriceBitCloutNanos
	})
	for _, v := range items {
		display.Row(sizes, v.Username, display.Float(v.MarketCap()),
			display.OneE9(v.CoinPriceBitCloutNanos))
	}
}
func LoopThruAllFollowing(pub58, username string) []models.ProfileEntryResponse {
	last := ""
	js := network.GetFollowsStateless(pub58, username, last)
	var pktpe models.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	NumFollowers := pktpe.NumFollowers
	total := map[string]bool{}
	bigList := []models.ProfileEntryResponse{}
	fmt.Println("Getting all", pktpe.NumFollowers, "...")
	for {
		for key, v := range pktpe.PublicKeyToProfileEntry {
			last = key
			if total[v.Username] == false {
				total[v.Username] = true
				bigList = append(bigList, v)
			}
		}
		if len(total) >= int(NumFollowers) {
			break
		}
		fmt.Println("got", len(bigList), "out of", NumFollowers)
		time.Sleep(time.Second * 1)
		js := network.GetFollowsStateless(pub58, username, last)
		json.Unmarshal([]byte(js), &pktpe)
	}
	return bigList
}
