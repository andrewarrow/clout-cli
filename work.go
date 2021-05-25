package main

import (
	"bufio"
	"clout/display"
	"clout/keys"
	"clout/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/justincampbell/timeago"
)

func PostsForPublicKey(key string) {
	js := GetSingleProfile(key)
	var sp models.SingleProfile
	json.Unmarshal([]byte(js), &sp)
	fmt.Println("---", sp.Profile.CoinEntry.CreatorBasisPoints)
	fmt.Println(sp.Profile.Description)
	fmt.Println("---")

	js = GetPostsForPublicKey(key)
	//b, _ := ioutil.ReadFile("samples/get_posts_for_public_key.list")
	var ppk models.PostsPublicKey
	json.Unmarshal([]byte(js), &ppk)
	for _, p := range ppk.Posts {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		body := ""
		if p.Body != "" {
			body = p.Body
		} else {
			body = p.RecloutedPostEntryResponse.Body
		}
		tokens := strings.Split(body, "\n")
		for i, t := range tokens {
			if strings.TrimSpace(t) == "" {
				continue
			}
			fmt.Println("    ", i, t)
		}
		fmt.Println(ago)
		fmt.Println("")
	}
}

func ListPosts(follow bool) {
	pub58 := LoggedInPub58()
	js := GetPostsStateless(pub58, follow)
	//b, _ := ioutil.ReadFile("samples/get_posts_stateless.list")
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	for i, p := range ps.PostsFound {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		fmt.Println(display.LeftAligned(p.ProfileEntryResponse.Username, 30),
			display.LeftAligned(p.ProfileEntryResponse.CoinEntry.NumberOfHolders, 20),
			ago)
		tokens := strings.Split(p.Body, "\n")
		fmt.Println("        ", display.LeftAligned(tokens[0], 40))
		fmt.Println("")
		if i > 6 {
			break
		}
	}
}

func Pub58ToUsername(key string) string {
	js := GetUsersStateless(key)
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	return us.UserList[0].ProfileEntryResponse.Username
}
func Post(reply string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Say: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(SeedBytes(mnemonic))
	bigString := SubmitPost(pub58, text, reply)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := SubmitTx(tx.TransactionHex, priv)
	fmt.Println(len(jsonString))
}
