package sync

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/justincampbell/timeago"
)

func HandleSync(limit string) {
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
	last := ""
	for {
		js := network.GetPostsStatelessWithOptions(last, pub58)
		var ps models.PostsStateless
		json.Unmarshal([]byte(js), &ps)

		for _, p := range ps.PostsFound {
			ts := time.Unix(p.TimestampNanos/1000000000, 0)
			ago := timeago.FromDuration(time.Since(ts))
			fmt.Println(ago)
			last = p.PostHashHex
		}
		time.Sleep(time.Second * 1)
	}
}
