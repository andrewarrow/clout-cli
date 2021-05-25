package main

import (
	"clout/display"
	"clout/models"
	"encoding/json"
	"fmt"
)

func ParseUserList(js string) map[string]string {
	m := map[string]string{}
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	for _, u := range us.UserList {
		m[u.ProfileEntryResponse.PublicKeyBase58Check] = u.ProfileEntryResponse.Username
	}
	return m
}
func ListNotifications() {
	pub58 := LoggedInPub58()
	//b, _ := ioutil.ReadFile("samples/get_notifications.list")
	js := GetNotifications(pub58)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	mapOfUsers := map[string]bool{}
	pub58ToUsernames := map[string]string{}
	buff := []string{}
	for _, n := range list.Notifications {
		mapOfUsers[n.Metadata.TransactorPublicKeyBase58Check] = true
	}
	fmt.Println("mapOfUsers length", len(mapOfUsers))
	for k, _ := range mapOfUsers {
		buff = append(buff, k)
		if len(buff) == 10 {
			fmt.Println("fetching 10...")
			js = GetManyUsersStateless(buff)
			m := ParseUserList(js)
			for k, v := range m {
				pub58ToUsernames[k] = v
			}
			buff = []string{}
		}
	}
	if len(buff) > 0 {
		fmt.Println("fetching ", len(buff))
		js = GetManyUsersStateless(buff)
		m := ParseUserList(js)
		for k, v := range m {
			pub58ToUsernames[k] = v
		}
	}

	for i, n := range list.Notifications {
		fmt.Printf("%02d %s %s %s\n", i, display.LeftAligned(n.Metadata.TxnType, 30),
			pub58ToUsernames[n.Metadata.TransactorPublicKeyBase58Check],
			n.Metadata.CreatorCoinTransferTxindexMetadata.CreatorUsername)
	}
}
