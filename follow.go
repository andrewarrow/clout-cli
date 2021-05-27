package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
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
	follower := LoggedInPub58()
	followed, _ := UsernameToPub58(username)
	jsonString := CreateFollow(follower, followed)
	var tx models.TxReady
	json.Unmarshal([]byte(jsonString), &tx)
	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	_, priv := keys.ComputeKeysFromSeed(SeedBytes(mnemonic))
	jsonString = SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
func HandleFollowing() {
	pub58 := LoggedInPub58()
	if len(os.Args) > 2 {
		pub58 = os.Args[2]
		LoopThruAllFollowing(pub58)
		return
	}
	js := GetFollowsStateless(pub58, "", "")

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
	pub58, username, _ := LoggedInAs()
	js := GetFollowsStateless(pub58, username, "")

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
	js := GetFollowsStateless(pub58, "", last)
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
		js := GetFollowsStateless(pub58, "", last)
		json.Unmarshal([]byte(js), &pktpe)
	}
}
