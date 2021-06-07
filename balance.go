package main

import (
	"clout/display"
	"clout/keys"
	"clout/session"
	"fmt"
	"sort"
)

func HandleBalances(argMap map[string]string) {
	m := session.ReadAccounts()
	for username, s := range m {
		fmt.Println("")
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		fmt.Println(username, pub58)
		user := session.Pub58ToUser(pub58)
		points := user.ProfileEntryResponse.CoinEntry.CreatorBasisPoints
		total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos

		fmt.Printf("  %s %.02f\n", display.LeftAligned("BalanceNano", 20), float64(user.BalanceNanos)/1000000.0)
		fmt.Printf("  %s %.02f\n", display.LeftAligned("MarketCap", 20), user.ProfileEntryResponse.MarketCap())
		fmt.Printf("  %s %d\n", display.LeftAligned("Points", 20), points)
		fmt.Printf("  %s %s\n", display.LeftAligned("Price", 20), display.OneE9(user.ProfileEntryResponse.CoinPriceBitCloutNanos))

		sort.SliceStable(user.UsersWhoHODLYou, func(i, j int) bool {
			return user.UsersWhoHODLYou[i].BalanceNanos >
				user.UsersWhoHODLYou[j].BalanceNanos
		})
		for i, friend := range user.UsersWhoHODLYou {
			coins := float64(friend.BalanceNanos) / 1000000000.0
			username := friend.ProfileEntryResponse.Username
			if username == "" {
				username = "anonymous"
			}
			fmt.Printf("  %s %0.6f %0.6f\n",
				display.LeftAligned(username, 30), coins,
				float64(friend.BalanceNanos)/float64(total))
			if username == "anonymous" {
				fmt.Println(friend.HODLerPublicKeyBase58Check)
			}

			if i > 9 {
				break
			}
		}
	}
	fmt.Println("")
}
