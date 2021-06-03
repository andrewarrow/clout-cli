package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
)

func HandleLike(argMap map[string]string) {
	hash := argMap["hash"]
	actor := session.LoggedInPub58()
	jsonString := network.CreateLike(actor, hash)
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
