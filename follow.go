package main

import (
	"clout/display"
	"clout/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func UsernameToPub58(s string) string {
	return ""
}

func HandleFollow() {
	if len(os.Args) < 3 {
		fmt.Println("missing username")
		return
	}
	username := os.Args[2]
	follower := LoggedInPub58()
	followed := UsernameToPub58(username)
	jsonString := CreateFollow(follower, followed)
	fmt.Println(jsonString)
}
func ListFollowing() {
	pub58 := LoggedInPub58()
	js := GetFollowsStateless(pub58, "")

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
	pub58, username := LoggedInAs()
	js := GetFollowsStateless(pub58, username)

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
