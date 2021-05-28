package main

import (
	"bufio"
	"clout/display"
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/justincampbell/timeago"
)

func HandlePosts() {
	if argMap["hash"] != "" {
		m := session.ReadShortMap()
		short := argMap["hash"]
		ShowSinglePost(m[short])
		return
	}
	ListPosts(argMap["follow"] == "true")
}
func ShowSinglePost(key string) {
	pub58 := session.LoggedInPub58()
	js := network.GetSinglePost(pub58, key)
	var ps models.PostStateless
	json.Unmarshal([]byte(js), &ps)

	fmt.Println("")
	fmt.Println(ps.PostFound.Body)
	fmt.Println("")

	fmt.Printf("%s %s %s\n", display.LeftAligned("username", 20),
		display.LeftAligned("ago", 15),
		display.LeftAligned("hash", 10))
	fmt.Printf("%s %s %s\n", display.LeftAligned("--------", 20),
		display.LeftAligned("---", 15),
		display.LeftAligned("--------", 10))
	for i, p := range ps.PostFound.Comments {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		//fmt.Println(display.LeftAligned(p.ProfileEntryResponse.Username, 30),
		//	display.LeftAligned(p.ProfileEntryResponse.CoinEntry.NumberOfHolders, 20),
		//	ago)
		//tokens := strings.Split(p.Body, "\n")
		//fmt.Println("        ", display.LeftAligned(tokens[0], 40))
		//fmt.Println("")
		username := p.ProfileEntryResponse.Username
		short := p.PostHashHex[0:7]
		fmt.Printf("%s %s %s\n", display.LeftAligned(username, 20),
			display.LeftAligned(ago, 15),
			display.LeftAligned(short, 10))

		if i > 20 {
			break
		}
	}
	fmt.Println("")
}

func PostsForPublicKey(key string) {
	js := network.GetSingleProfile(key)
	var sp models.SingleProfile
	json.Unmarshal([]byte(js), &sp)
	fmt.Println("---", sp.Profile.CoinEntry.CreatorBasisPoints)
	fmt.Println(sp.Profile.Description)
	fmt.Println("---")

	js = network.GetPostsForPublicKey(key)
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
	pub58 := session.LoggedInPub58()
	js := network.GetPostsStateless(pub58, follow)
	//b, _ := ioutil.ReadFile("samples/get_posts_stateless.list")
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	fmt.Printf("%s %s %s %s %s\n", display.LeftAligned("username", 20),
		display.LeftAligned("ago", 15),
		display.LeftAligned("replies", 10),
		display.LeftAligned("reclouts", 10),
		display.LeftAligned("hash", 10))
	fmt.Printf("%s %s %s %s %s\n", display.LeftAligned("--------", 20),
		display.LeftAligned("---", 15),
		display.LeftAligned("-------", 10),
		display.LeftAligned("-------", 10),
		display.LeftAligned("--------", 10))
	shortMap := map[string]string{}
	for i, p := range ps.PostsFound {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))

		//fmt.Println(display.LeftAligned(p.ProfileEntryResponse.Username, 30),
		//	display.LeftAligned(p.ProfileEntryResponse.CoinEntry.NumberOfHolders, 20),
		//	ago)
		//tokens := strings.Split(p.Body, "\n")
		//fmt.Println("        ", display.LeftAligned(tokens[0], 40))
		//fmt.Println("")

		username := p.ProfileEntryResponse.Username
		short := p.PostHashHex[0:7]
		shortMap[short] = p.PostHashHex
		fmt.Printf("%s %s %s %s %s\n", display.LeftAligned(username, 20),
			display.LeftAligned(ago, 15),
			display.LeftAligned(p.CommentCount, 10),
			display.LeftAligned(p.RecloutCount, 10),
			display.LeftAligned(short, 10))

		if i > 20 {
			break
		}
	}
	session.SaveShortMap(shortMap)
}

func Post(reply string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Say: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	bigString := network.SubmitPost(pub58, text, reply)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
