package main

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"time"
)

func HandleMachine() {

	id := argMap["id"]
	if id == "" {
		return
	}

	LoopThruAllComments(id)
}

func CheckForValidEntry(username, body string) {
	fmt.Printf("%s commented with length %d\n", username, len(body))
}

func LoopThruAllComments(key string) {
	offset := int64(0)
	pub58 := session.LoggedInPub58()

	for {
		js := network.GetSinglePostWithOffset(offset, pub58, key)
		var ps models.PostStateless
		json.Unmarshal([]byte(js), &ps)

		if len(ps.PostFound.Comments) == 0 {
			break
		}

		for _, p := range ps.PostFound.Comments {
			CheckForValidEntry(p.ProfileEntryResponse.Username, p.Body)
		}
		offset += 20
		time.Sleep(time.Second * 1)
	}
}
