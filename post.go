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
	if len(os.Args) == 3 && !strings.HasPrefix(os.Args[2], "--") {
		m := session.ReadShortMap()
		ShowSinglePost(m[os.Args[2]])
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
	if ps.PostFound.Body == "" {
		long := ps.PostFound.RecloutedPostEntryResponse.PostHashHex
		short := long[0:7]
		session.SaveShortMap(map[string]string{short: long})
		fmt.Println("RECLOUT", short)
		fmt.Println(ps.PostFound.RecloutedPostEntryResponse.Body)
	} else {
		fmt.Println(ps.PostFound.Body)
	}
	fmt.Println("")
	for _, image := range ps.PostFound.ImageURLs {
		fmt.Println(image)
	}
	fmt.Println("")

	LsHeader()
	shortMap := map[string]string{}
	for i, p := range ps.PostFound.Comments {
		LsPost(p, shortMap)
		if i > 20 {
			break
		}
	}
	session.SaveShortMap(shortMap)
	fmt.Println("")
}

func LsHeader() {
	fmt.Printf("%s %s %s %s %s %s %s\n", display.LeftAligned("username", 20),
		display.LeftAligned("ago", 15),
		display.LeftAligned("likes", 6),
		display.LeftAligned("replies", 8),
		display.LeftAligned("reclouts", 9),
		display.LeftAligned("cap", 10),
		display.LeftAligned("hash", 10))
	fmt.Printf("%s %s %s %s %s %s %s\n", display.LeftAligned("--------", 20),
		display.LeftAligned("---", 15),
		display.LeftAligned("-----", 6),
		display.LeftAligned("-------", 8),
		display.LeftAligned("-------", 9),
		display.LeftAligned("-------", 10),
		display.LeftAligned("--------", 10))
}

func LsPost(p models.Post, shortMap map[string]string) {
	ts := time.Unix(p.TimestampNanos/1000000000, 0)
	ago := timeago.FromDuration(time.Since(ts))

	marketCap := p.ProfileEntryResponse.MarketCap()

	username := p.ProfileEntryResponse.Username
	short := p.PostHashHex[0:7]
	shortMap[short] = p.PostHashHex
	fmt.Printf("%s %s %s %s %s %s %s\n", display.LeftAligned(username, 20),
		display.LeftAligned(ago, 15),
		display.LeftAligned(p.LikeCount, 6),
		display.LeftAligned(p.CommentCount, 8),
		display.LeftAligned(p.RecloutCount, 9),
		display.LeftAligned(fmt.Sprintf("%.02f", marketCap), 10),
		display.LeftAligned(short, 10))
}

func ListPosts(follow bool) {
	rateString := network.GetExchangeRate()
	var rate models.Rate
	json.Unmarshal([]byte(rateString), &rate)
	//TODO use rate

	pub58 := session.LoggedInPub58()
	js := network.GetPostsStateless(pub58, follow)
	//b, _ := ioutil.ReadFile("samples/get_posts_stateless.list")
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	shortMap := map[string]string{}
	LsHeader()
	for i, p := range ps.PostsFound {

		LsPost(p, shortMap)
		if i > 20 {
			break
		}
	}
	session.SaveShortMap(shortMap)
}

func Post(argMap map[string]string) {
	reply := argMap["reply"]
	imagePath := argMap["image"]

	shortMap := session.ReadShortMap()

	longHash := ""

	if len(reply) == 7 {
		longHash = shortMap[reply]
	} else {
		longHash = reply
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Say: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	imageUrl := ""
	if imagePath != "" {
		imageUrl = UploadImage(imagePath, pub58, priv)
		if imageUrl == "" {
			return
		}
		imageUrl = "\"" + imageUrl + "\""
	}

	bigString := network.SubmitPost(pub58, text, longHash, imageUrl)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
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
