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

	m := session.ReadAccounts()
	for username, s := range m {
		fmt.Println("")
		fmt.Println("===========")
		fmt.Println(username)
		fmt.Println("===========")
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		session.Pub58ToBoards(pub58)
	}
	fmt.Println("")
}
