package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func HandleUnFollow() {
	if argMap["mass"] != "" {
		UnfollowInMass()
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("missing username")
		return
	}
	username := os.Args[2]
	UnfollowOne(username)
}

func UnfollowOne(username string) {
	follower := session.LoggedInPub58()
	followed := session.UsernameToPub58(username)
	jsonString := network.CreateUnFollow(follower, followed)
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

func UnfollowInMass() {
	pub58, _, _, _ := session.LoggedInAs()
	items := LoopThruAllFollowing(pub58, "", 100)
	for _, v := range items {
		UnfollowOne(v.Username)
		time.Sleep(time.Second * 1)
	}
}
