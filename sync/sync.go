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
	if argMap["top"] != "" {
		FindTopReclouted()
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

	SyncLoop()
}

func SyncLoop() {
	pub58 := session.LoggedInPub58()
	last := ""
	for {
		js := network.GetPostsStatelessWithOptions(last, pub58)
		var ps models.PostsStateless
		json.Unmarshal([]byte(js), &ps)

		fmt.Printf("PostsFound %d\n", len(ps.PostsFound))
		for _, p := range ps.PostsFound {
			ts := time.Unix(p.TimestampNanos/1000000000, 0)
			InsertPost("", p.RecloutCount, ts, p.PostHashHex, p.Body, p.ProfileEntryResponse.Username)
			if p.CommentCount > 0 {
				LoopThruAllCommentsToInsert("", p.PostHashHex)
			}
		}
	}
}
func LoopThruAllCommentsToInsert(tabs, key string) {
	offset := int64(0)
	pub58 := session.LoggedInPub58()

	for {
		js := network.GetSinglePostWithOffset(offset, pub58, key)
		var ps models.PostStateless
		json.Unmarshal([]byte(js), &ps)

		if len(ps.PostFound.Comments) == 0 {
			break
		}

		fmt.Printf("PostFound.Comments %d\n", len(ps.PostFound.Comments))
		for _, p := range ps.PostFound.Comments {
			ts := time.Unix(p.TimestampNanos/1000000000, 0)
			InsertPost(key, p.RecloutCount, ts, p.PostHashHex, p.Body, p.ProfileEntryResponse.Username)
			if p.CommentCount > 0 {
				LoopThruAllCommentsToInsert(tabs+"  ", p.PostHashHex)
			}
		}
		offset += 20
		time.Sleep(time.Second * 1)
	}
}

func OldSyncLoop() {
	pub58 := session.LoggedInPub58()
	last := LastHash()
	last = ""
	fmt.Println(last)
	for {
		js := network.GetPostsStatelessWithOptions(last, pub58)
		var ps models.PostsStateless
		json.Unmarshal([]byte(js), &ps)

		userMap := map[string]bool{}
		for _, p := range ps.PostsFound {
			ts := time.Unix(p.TimestampNanos/1000000000, 0)
			InsertPost("", p.RecloutCount, ts, p.PostHashHex, p.Body, p.ProfileEntryResponse.Username)
			userMap[p.ProfileEntryResponse.PublicKeyBase58Check] = true
			last = p.PostHashHex
		}

		users := []string{}
		for k, _ := range userMap {
			users = append(users, k)
		}

		js = network.GetManyUsersStateless(users)

		var us models.UsersStateless
		json.Unmarshal([]byte(js), &us)
		for _, u := range us.UserList {
			marketCap := u.ProfileEntryResponse.MarketCap()
			numUsersYouHODL := len(u.UsersYouHODL)
			numBoardMembers := 0
			for _, thing := range u.UsersYouHODL {
				coins := float64(thing.BalanceNanos) / 1000000000.0
				if coins < 1 {
					continue
				}
				numBoardMembers++
			}
			InsertUser(fmt.Sprintf("%.02f", marketCap), numUsersYouHODL, numBoardMembers,
				u.ProfileEntryResponse.CoinEntry.CreatorBasisPoints/100,
				u.ProfileEntryResponse.PublicKeyBase58Check,
				u.ProfileEntryResponse.Username)
		}

		fmt.Println(len(ps.PostsFound))
		if len(ps.PostsFound) == 0 {
			return
		}
		time.Sleep(time.Second * 1)
		fmt.Println(last)
	}
}
