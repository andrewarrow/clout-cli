package main

import (
	"clout/keys"
	"clout/session"
	"fmt"
)

func HandleBoards() {

	fmt.Println("")
	fmt.Println("Boards are coins where you are one of the few significant")
	fmt.Println("owners and have some responsibilities.")
	fmt.Println("")

	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	seedBytes := session.SeedBytes(mnemonic)
	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	session.Pub58ToBoards(pub58)
	fmt.Println("")
	fmt.Println("")
}
