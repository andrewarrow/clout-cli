package main

import (
	"clout/display"
	"clout/keys"
	"clout/session"
	"fmt"
)

func HandleBalances(argMap map[string]string) {
	m := session.ReadAccounts()
	for username, s := range m {
		fmt.Println("")
		fmt.Println(username)
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		user := session.Pub58ToUser(pub58)
		fmt.Printf("  %s %.02f\n", display.LeftAligned("BalanceNano", 20), float64(user.BalanceNanos)/1000000.0)
		fmt.Printf("  %s %.02f\n", display.LeftAligned("MarketCap", 20), user.ProfileEntryResponse.MarketCap())
	}
	fmt.Println("")
}
