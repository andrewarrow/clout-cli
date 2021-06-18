package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"os"
)

func HandleDiamond() {
	if len(os.Args) < 3 {
		fmt.Println("missing username")
		return
	}
	level := "1"
	if argMap["level"] != "" {
		level = argMap["level"]
	}
	username := os.Args[2]
	theirPub58 := session.UsernameToPub58(username)
	js := network.GetPostsForPublicKey(username)
	var ppk models.PostsPublicKey
	json.Unmarshal([]byte(js), &ppk)
	if len(ppk.Posts) == 0 {
		return
	}
	lastPost := ppk.Posts[0].PostHashHex

	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	bigString := network.SubmitDiamond(level, pub58, theirPub58, lastPost)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
