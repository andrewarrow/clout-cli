package main

import (
	"clout/display"
	"clout/keys"
	"clout/session"
	"fmt"
	"sort"
)

func HandleBalances(argMap map[string]string) {
	list := session.ReadAccountsSorted()
	m := session.ReadAccounts()
	fields := []string{"username", "% owned", "co-owner1", "co-owner2", "balance"}
	sizes := []int{15, 10, 20, 20, 10}
	display.Header(sizes, fields...)
	for _, username := range list {
		s := m[username]
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		user := session.Pub58ToUser(pub58)
		//points := user.ProfileEntryResponse.CoinEntry.CreatorBasisPoints
		total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos

		holdMap := map[string]string{}
		sort.SliceStable(user.UsersWhoHODLYou, func(i, j int) bool {
			return user.UsersWhoHODLYou[i].BalanceNanos >
				user.UsersWhoHODLYou[j].BalanceNanos
		})
		topThree := []string{}
		for _, friend := range user.UsersWhoHODLYou {
			username := friend.ProfileEntryResponse.Username
			if username == "" {
				username = "anonymous" //friend.HODLerPublicKeyBase58Check
			}

			perString := fmt.Sprintf("%d",
				int(100*(float64(friend.BalanceNanos)/float64(total))))
			holdMap[username] = perString

			if len(topThree) < 3 {
				topThree = append(topThree, username)
			}
		}

		others := []string{}
		for _, u := range topThree {
			if u == username {
				continue
			}
			others = append(others, u)
		}

		co1 := ""
		co2 := ""

		if len(others) == 2 {
			co1 = others[0] + " " + holdMap[others[0]]
			co2 = others[1] + " " + holdMap[others[1]]
		} else if len(others) == 1 {
			co1 = others[0] + " " + holdMap[others[0]]
		}

		display.Row(sizes, username, holdMap[username], co1, co2,
			display.OneE9(user.BalanceNanos))
	}
}
