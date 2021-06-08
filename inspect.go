package main

import (
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func HandleInspect() {
	if len(os.Args) < 3 {
		fmt.Println("username")
		return
	}

	username := os.Args[2]
	js := network.GetSingleProfile(username)
	var sp models.SingleProfile
	json.Unmarshal([]byte(js), &sp)
	tokens := strings.Split(sp.Profile.Description, "\n")
	for _, token := range tokens {
		tokens = strings.Split(token, " ")
		for _, token := range tokens {
			fmt.Println(token)
		}
	}
}
