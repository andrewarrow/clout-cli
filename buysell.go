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

func HandleBuy() {
	if len(os.Args) < 4 {
		fmt.Println("missing username or amount")
		return
	}
	username := os.Args[2]
	//amountString := os.Args[3]
	theirPub58 := session.UsernameToPub58(username)

	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	user := session.Pub58ToUser(pub58)
	amount := user.BalanceNanos - 100318
	// 604713
	// 504395
	//amount, _ := strconv.ParseInt(amountString, 10, 64)
	bigString := network.SubmitBuyCoin(pub58, theirPub58, amount, 0)
	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)
	fmt.Println(user.BalanceNanos, amount, tx.ExpectedCreatorCoinReturnedNanos)

	bigString = network.SubmitBuyCoin(pub58, theirPub58, amount, tx.ExpectedCreatorCoinReturnedNanos)
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
func HandleSell() {
	if len(os.Args) < 4 {
		fmt.Println("missing username or amount")
		return
	}
	username := os.Args[2]
	//amountString := os.Args[3]
	theirPub58 := session.UsernameToPub58(username)

	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	user := session.Pub58ToUser(pub58)
	amount := int64(0)
	for _, item := range user.UsersYouHODL {
		if username != item.ProfileEntryResponse.Username {
			continue
		}
		fmt.Println(item.BalanceNanos)
		amount = item.BalanceNanos
		break
	}
	bigString := network.SubmitSellCoin(pub58, theirPub58, amount, 0)
	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)
	fmt.Println(user.BalanceNanos, amount, tx.ExpectedBitCloutReturnedNanos)

	bigString = network.SubmitSellCoin(pub58, theirPub58, amount, tx.ExpectedBitCloutReturnedNanos)
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
