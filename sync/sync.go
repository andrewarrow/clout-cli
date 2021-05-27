package sync

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func HandleSync(argMap map[string]string) {
	CreateSchema()
	if argMap["query"] != "" {
		FindPosts(argMap["query"])
		return
	}
	limit := argMap["limit"]
	fmt.Println("-=-=-= SYNC =-=-=-")
	fmt.Println("Run this in background to query nodes for blockchain")
	fmt.Println("data about the recent past, further and further back in time.")
	fmt.Println("")
	fmt.Println("You can never reach genesis, it's getting more")
	fmt.Println("difficult as time goes on; the blockchain gets larger and larger.")
	fmt.Println("")
	fmt.Println("1. Decide how much hard drive space you want to")
	fmt.Println("   allocate to this.")
	fmt.Println("")
	fmt.Println("2. Start `clout sync --limit=2000000000` with the number")
	fmt.Println("   of bytes limit you pick.")
	fmt.Println("")
	fmt.Println("Then as your drives fills with data, we index it, makes searching")
	fmt.Println("for old content better and better.")
	fmt.Println("")
	if limit == "" {
		return
	}
	limitBytes, _ := strconv.ParseInt(limit, 10, 64)
	k := limitBytes / 1000
	mb := k / 1000
	fmt.Printf("Syncing now... %d mb\n", mb)

	for {
		SyncLoop()
	}
}

func SyncLoop() {
	pub58 := session.LoggedInPub58()
	last := LastHash()
	fmt.Println(last)
	for {
		js := network.GetPostsStatelessWithOptions(last, pub58)
		var ps models.PostsStateless
		json.Unmarshal([]byte(js), &ps)

		for _, p := range ps.PostsFound {
			ts := time.Unix(p.TimestampNanos/1000000000, 0)
			InsertPost(ts, p.PostHashHex, p.Body, p.ProfileEntryResponse.Username)
			InsertUser(p.ProfileEntryResponse.PublicKeyBase58Check, p.ProfileEntryResponse.Username)
			last = p.PostHashHex
		}
		fmt.Println(len(ps.PostsFound))
		if len(ps.PostsFound) == 0 {
			break
		}
		time.Sleep(time.Second * 1)
		fmt.Println(last)
	}
}
