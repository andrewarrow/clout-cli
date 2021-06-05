package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
)

func HandleSend() {
	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))

	dest := argMap["dest"]
	if dest == "" {
		return
	}

	amountInNanos := int64(1815567 - 1200)

	bigString := network.SendBitclout(pub58, dest, amountInNanos)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("SubmitTx Success!")
	}
}
