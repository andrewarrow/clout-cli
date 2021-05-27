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
		LoopThruAllFollowing(pub58)
		return
	}
	js := network.GetFollowsStateless(pub58, "", "")

	var pktpe models.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	fmt.Println("NumFollowers", pktpe.NumFollowers)
	fmt.Println("")
	for _, v := range pktpe.PublicKeyToProfileEntry {
		tokens := strings.Split(v.Description, "\n")
		fmt.Printf("%s %s\n", display.LeftAligned(v.Username, 30),
			display.LeftAligned(tokens[0], 30))
	}
}
func ListFollowers() {
	pub58, username, _ := session.LoggedInAs()
	js := network.GetFollowsStateless(pub58, username, "")

	var pktpe models.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	fmt.Println("NumFollowers", pktpe.NumFollowers)
	fmt.Println("")
	for _, v := range pktpe.PublicKeyToProfileEntry {
		tokens := strings.Split(v.Description, "\n")
		fmt.Printf("%s %s\n", display.LeftAligned(v.Username, 30),
			display.LeftAligned(tokens[0], 30))
	}
}
func LoopThruAllFollowing(pub58 string) {
	last := ""
	//f, _ := os.OpenFile("i.follow", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//defer f.Close()
	js := network.GetFollowsStateless(pub58, "", last)
	var pktpe models.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	fmt.Println("NumFollowers", pktpe.NumFollowers)
	NumFollowers := pktpe.NumFollowers
	total := map[string]bool{}
	for {
		for key, v := range pktpe.PublicKeyToProfileEntry {
			last = key
			//f.WriteString(v.Username + "\n")
			if total[v.Username] == false {
				fmt.Println(v.Username)
				total[v.Username] = true
			}
		}
		if len(total) >= int(NumFollowers) {
			break
		}
		time.Sleep(time.Second * 1)
		js := network.GetFollowsStateless(pub58, "", last)
		json.Unmarshal([]byte(js), &pktpe)
	}
}
