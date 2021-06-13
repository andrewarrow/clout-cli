package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"sort"
)

func HandleBalances(argMap map[string]string) {
	list := session.ReadAccountsSorted()
	m := session.ReadAccounts()
	fields := []string{"username", "% owned", "co-owner1", "co-owner2", "balance", "price"}
	sizes := []int{15, 10, 15, 15, 10, 10}
	display.Header(sizes, fields...)
	for _, username := range list {
		s := m[username]
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		js := network.GetHodlers(username)
		var hw models.HodlersWrap
		json.Unmarshal([]byte(js), &hw)
		user := session.Pub58ToUser(pub58)
		//points := user.ProfileEntryResponse.CoinEntry.CreatorBasisPoints
		total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos
		price := user.ProfileEntryResponse.CoinPriceBitCloutNanos

		holdMap := map[string]string{}
		sort.SliceStable(hw.Hodlers, func(i, j int) bool {
			return hw.Hodlers[i].BalanceNanos >
				hw.Hodlers[j].BalanceNanos
		})
		topThree := []string{}
		for _, friend := range hw.Hodlers {
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
			display.OneE9(user.BalanceNanos), display.OneE9(price))
	}
}
