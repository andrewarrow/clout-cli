package main

import (
	"clout/display"
	"clout/models"
	"encoding/json"
	"fmt"
	"strings"
)

func ListFollowing() {
	js := GetFollowsStateless("")

	var pktpe models.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	fmt.Println("NumFollowers", pktpe.NumFollowers)
	fmt.Println("")
	for _, v := range pktpe.PublicKeyToProfileEntry {
		tokens := strings.Split(v.Description, "\n")
		fmt.Printf("%s %s\n", display.LeftAligned(v.Username, 30),
			display.LeftAligned(tokens[0], 30))
	}
}
func ListFollowers() {
	_, username := LoggedInAs()
	js := GetFollowsStateless(username)

	var pktpe models.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	fmt.Println("NumFollowers", pktpe.NumFollowers)
	fmt.Println("")
	for _, v := range pktpe.PublicKeyToProfileEntry {
		tokens := strings.Split(v.Description, "\n")
		fmt.Printf("%s %s\n", display.LeftAligned(v.Username, 30),
			display.LeftAligned(tokens[0], 30))
	}
}
