package main

import (
	"clout/display"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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
	tokens := strings.Split(body, "\n")
	if len(tokens) > 1 {
		fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "ERR has newlines")
	} else {
		tokens = strings.Split(tokens[0], "=")
		if len(tokens) == 2 {
			amount := strings.TrimSpace(tokens[1])
			if strings.HasPrefix(amount, "$") {
				intAmount, _ := strconv.Atoi(amount[1:])
				if intAmount > 0 {
					fmt.Printf("%s %d\n", display.LeftAligned(username, 20), intAmount)
				} else {
					fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "bad amount")
				}
			} else {
				fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "missing $")
			}
		} else {
			fmt.Printf("%s %s\n", display.LeftAligned(username, 20), "Bad =")
		}
	}
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
