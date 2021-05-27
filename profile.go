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

func HandleUpdateProfile() {
	if len(os.Args) < 3 {
		fmt.Println("missing desc")
		return
	}
	desc := os.Args[2]
	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	jsonString := network.UpdateProfile(pub58, desc)
	var tx models.TxReady
	json.Unmarshal([]byte(jsonString), &tx)
	jsonString = network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
