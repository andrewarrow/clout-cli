package main

import (
	"clout/display"
	"clout/models"
	"encoding/json"
	"fmt"
)

func ListNotifications() {
	pub58 := LoggedInPub58()
	//b, _ := ioutil.ReadFile("samples/get_notifications.list")
	js := GetNotifications(pub58)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	for i, n := range list.Notifications {
		fmt.Printf("%02d %s %s\n", i, display.LeftAligned(n.Metadata.TxnType, 30),
			n.Metadata.CreatorCoinTransferTxindexMetadata.CreatorUsername)
	}
}
