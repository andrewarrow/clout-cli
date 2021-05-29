package main

import (
	"clout/keys"
	"clout/session"
	"fmt"
)

func HandleBalances(argMap map[string]string) {
	m := session.ReadAccounts()
	for username, s := range m {
		fmt.Println("")
		fmt.Println("===========")
		fmt.Println(username)
		fmt.Println("===========")
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		_, balance := session.Pub58ToUsername(pub58)
		fmt.Printf("%.02f\n", float64(balance)/1000000.0)
	}
	fmt.Println("")
}
