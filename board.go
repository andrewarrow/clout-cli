package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
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
		Pub58ToBoards(pub58)
	}
	fmt.Println("")
}

func Pub58ToBoards(key string) {
	js := network.GetUsersStateless(key)
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	for _, thing := range us.UserList[0].UsersYouHODL {
		coins := float64(thing.BalanceNanos) / 1000000000.0
		if coins < 1 {
			continue
		}
		fmt.Printf("%s %0.2f\n",
			display.LeftAligned(thing.ProfileEntryResponse.Username, 30), coins)

		other := thing.ProfileEntryResponse.PublicKeyBase58Check
		js = network.GetUsersStateless(other)
		var us models.UsersStateless
		json.Unmarshal([]byte(js), &us)
		for _, friend := range us.UserList[0].UsersWhoHODLYou {
			if friend.ProfileEntryResponse.PublicKeyBase58Check == key {
				continue
			}
			coins := float64(friend.BalanceNanos) / 1000000000.0
			if coins < 1 {
				continue
			}
			username := friend.ProfileEntryResponse.Username
			if username == "" {
				username = "anonymous"
			}
			fmt.Printf("  %s %0.2f\n",
				display.LeftAligned(username, 30), coins)
		}
	}
}
