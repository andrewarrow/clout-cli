package main

import (
	"clout/display"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
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
	fields := []string{"username"}
	sizes := []int{20}
	display.Header(sizes, fields...)
	for _, thing := range us.UserList[0].UsersYouHODL {
		display.Row(sizes, thing.ProfileEntryResponse.Username)
	}
}
