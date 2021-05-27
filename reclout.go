package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"os"
)

func HandleReclout() {
	if len(os.Args) < 3 {
		fmt.Println("missing username")
		return
	}
	username := os.Args[2]
	lastPost := username // hack to allow passing in specific msg
	if len(username) < 64 {
		js := network.GetPostsForPublicKey(username)
		var ppk models.PostsPublicKey
		json.Unmarshal([]byte(js), &ppk)
		if len(ppk.Posts) == 0 {
			return
		}
		lastPost = ppk.Posts[0].PostHashHex
	}

	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(SeedBytes(mnemonic))
	bigString := network.SubmitReclout(pub58, lastPost)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
