package main

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// 106579af0baa776bcda80ffea77b3549be64326dfe2bfef9978f28b7211ce9fc
// 0c86890b5955cf28660c91797dc93cfdc12b2e2d4cd45cda4e26635c7cc75a78
func HandleLongThread() {
	id := argMap["id"]
	if id == "" {
		return
	}

	m := map[string]map[string]bool{}
	LoopThruLongThread(id, m)

	for k, v := range m {
		fmt.Printf("%s was mentioned by %d other users:\n", k, len(v))
		for kk, _ := range v {
			fmt.Println(" ", kk)
		}
	}

}
func CheckForContent(username, body string, m map[string]map[string]bool) {
	fmt.Println(username)
	tokens := strings.Split(body, "\n")
	for _, line := range tokens {
		tokens = strings.Split(line, " ")
		for _, word := range tokens {
			if strings.HasPrefix(word, "@") {
				tagged := word[1:]
				if m[tagged] == nil {
					m[tagged] = map[string]bool{}
				}
				m[tagged][username] = true
			}
		}
	}
}

func LoopThruLongThread(key string, m map[string]map[string]bool) {
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
			CheckForContent(p.ProfileEntryResponse.Username, p.Body, m)
		}
		offset += 20
		time.Sleep(time.Second * 1)
	}
}
