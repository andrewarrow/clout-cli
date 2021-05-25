package main

import (
	"clout/display"
	"clout/models"
	"encoding/json"
	"fmt"
)

func ParseUserList(js string) {
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	for _, u := range us.UserList {
		fmt.Println(u.ProfileEntryResponse.Username)
	}
}
func ListNotifications() {
	pub58 := LoggedInPub58()
	//b, _ := ioutil.ReadFile("samples/get_notifications.list")
	js := GetNotifications(pub58)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	mapOfUsers := map[string]bool{}
	buff := []string{}
	for _, n := range list.Notifications {
		mapOfUsers[n.Metadata.TransactorPublicKeyBase58Check] = true
	}
	fmt.Println("mapOfUsers length", len(mapOfUsers))
	for k, _ := range mapOfUsers {
		buff = append(buff, k)
		if len(buff) == 10 {
			js = GetManyUsersStateless(buff)
			ParseUserList(js)
			buff = []string{}
		}
	}
	if len(buff) > 0 {
		js = GetManyUsersStateless(buff)
		ParseUserList(js)
	}

	for i, n := range list.Notifications {
		fmt.Printf("%02d %s %s\n", i, display.LeftAligned(n.Metadata.TxnType, 30),
			n.Metadata.CreatorCoinTransferTxindexMetadata.CreatorUsername)
	}
}
