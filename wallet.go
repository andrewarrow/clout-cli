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
	display.Header(sizes, fields...)

	YouHODL := us.UserList[0].UsersYouHODL
	sort.SliceStable(YouHODL, func(i, j int) bool {
		return YouHODL[i].BalanceNanos > YouHODL[j].BalanceNanos
	})

	for _, thing := range YouHODL {
		display.Row(sizes,
			thing.ProfileEntryResponse.Username,
			display.OneE9(thing.BalanceNanos),
			display.OneE9(thing.ProfileEntryResponse.CoinPriceBitCloutNanos),

			display.Float(display.OneE9Float(thing.ProfileEntryResponse.CoinPriceBitCloutNanos)*
				display.OneE9Float(thing.BalanceNanos)))
	}
}
