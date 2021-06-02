package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func HandleMachine() {

	id := argMap["id"]
	if id == "" {
		return
	}

	LoopThruAllComments(id)
}

func ThreeEmojiWorkForAmount(s string, amount int) bool {
	sum := 0
	for _, r := range []rune(s) {
		hex := fmt.Sprintf("%x", r)
		sum += int(SumIt(hex))
	}
	return sum == amount
}

func AwardMonies(username string, amount int) {
	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))

	creator := session.UsernameToPub58("TheClown")

	// TODO use passed in amount to make this the right amount to give
	amountInNanos := int64(100000000)

	bigString := network.SubmitTransferCoin(pub58, creator, username, amountInNanos)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("SubmitTx Success!")
	}
}

func CheckForValidEntry(username, body string) {
	tokens := strings.Split(body, "\n")
	if len(tokens) > 1 {
		fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "ERR has newlines")
	} else {
		tokens = strings.Split(tokens[0], "=")
		if len(tokens) == 2 {
			three := strings.TrimSpace(tokens[0])
			amount := strings.TrimSpace(tokens[1])
			if strings.HasPrefix(amount, "$") {
				intAmount, _ := strconv.Atoi(amount[1:])
				if intAmount > 0 {
					if ThreeEmojiWorkForAmount(three, intAmount) {
						fmt.Printf("%s %s %d\n", display.LeftAligned(username, 20), "SUCCESS", intAmount)
						AwardMonies(username, intAmount)
					} else {
						fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "emoji != amount")
					}
				} else {
					fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "bad amount")
				}
			} else {
				fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "missing $")
			}
		} else {
			fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "Bad =")
		}
	}
}

func LoopThruAllComments(key string) {
	offset := int64(0)
	pub58 := session.LoggedInPub58()

	for {
		js := network.GetSinglePostWithOffset(offset, pub58, key)
		var ps models.PostStateless
		json.Unmarshal([]byte(js), &ps)

		if len(ps.PostFound.Comments) == 0 {
			break
		}

		for _, p := range ps.PostFound.Comments {
			CheckForValidEntry(p.ProfileEntryResponse.Username, p.Body)
		}
		offset += 20
		time.Sleep(time.Second * 1)
	}
}
