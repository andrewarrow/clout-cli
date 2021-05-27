package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
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
	theirPub58 := session.UsernameToPub58(username)

	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	bigString := network.SubmitBuyOrSellCoin(pub58, theirPub58, 0, 0)
	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	amount, _ := strconv.ParseInt(amountString, 10, 64)

	bigString = network.SubmitBuyOrSellCoin(pub58, theirPub58, amount, tx.ExpectedCreatorCoinReturnedNanos)
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
