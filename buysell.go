package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func HandleBuy() {
	if len(os.Args) < 4 {
		fmt.Println("missing username or amount")
		return
	}
	username := os.Args[2]
	amountString := os.Args[3]
	theirPub58 := UsernameToPub58(username)

	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(SeedBytes(mnemonic))

	amount, _ := strconv.ParseInt(amountString, 10, 64)
	bigString := network.SubmitBuyOrSellCoin(pub58, theirPub58, amount)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
