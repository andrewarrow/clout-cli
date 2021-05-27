package main

import (
	"clout/keys"
	"clout/models"
	"encoding/json"
	"fmt"
	"os"
)

func HandleDiamond() {
	if len(os.Args) < 3 {
		fmt.Println("missing username")
		return
	}
	username := os.Args[2]
	theirPub58, _ := UsernameToPub58(username)
	js := GetPostsForPublicKey(username)
	var ppk models.PostsPublicKey
	json.Unmarshal([]byte(js), &ppk)
	if len(ppk.Posts) == 0 {
		return
	}
	lastPost := ppk.Posts[0].PostHashHex

	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(SeedBytes(mnemonic))
	bigString := SubmitDiamond(pub58, theirPub58, lastPost)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
