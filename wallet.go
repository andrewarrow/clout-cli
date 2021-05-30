package main

import (
	"clout/display"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"sort"
)

func HandleWallet(argMap map[string]string) {
	pub58, username, balance, _ := session.LoggedInAs()
	fmt.Println(username, balance)
	fmt.Println("")
	ListAssets(pub58)
}
func ListAssets(key string) {
	js := network.GetUsersStateless(key)
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	fields := []string{"username", "balance", "price", "worth"}
	sizes := []int{20, 11, 10, 10}

	YouHODL := us.UserList[0].UsersYouHODL
	sort.SliceStable(YouHODL, func(i, j int) bool {
		return YouHODL[i].BalanceNanos > YouHODL[j].BalanceNanos
	})

	total := 0.0
	for _, thing := range YouHODL {
		value := display.OneE9Float(thing.ProfileEntryResponse.CoinPriceBitCloutNanos) *
			display.OneE9Float(thing.BalanceNanos)
		total += value
	}
	display.Row(sizes,
		"",
		"",
		"",
		display.Float(total))

	display.Header(sizes, fields...)
	for _, thing := range YouHODL {
		value := display.OneE9Float(thing.ProfileEntryResponse.CoinPriceBitCloutNanos) *
			display.OneE9Float(thing.BalanceNanos)
		display.Row(sizes,
			thing.ProfileEntryResponse.Username,
			display.OneE9(thing.BalanceNanos),
			display.OneE9(thing.ProfileEntryResponse.CoinPriceBitCloutNanos),
			display.Float(value))
	}

}
