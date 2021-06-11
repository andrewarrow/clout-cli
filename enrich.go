package main

import (
	"clout/display"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/wcharczuk/go-chart"
)

var alreadyDone map[string]bool
var r models.Rate

func FindBuysSellsAndTransfers() {
	session.WriteSelected("enrich")
	js := network.GetExchangeRate()
	json.Unmarshal([]byte(js), &r)

	alreadyDone = LoadEnrichMessages()
	alreadyDone["douglasss"] = true
	fmt.Println(alreadyDone)
	pub58 := session.LoggedInPub58()
	js = network.GetPostsStateless(pub58, false)
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	for _, p := range ps.PostsFound {
		username := p.ProfileEntryResponse.Username
		if alreadyDone[username] {
			continue
		}
		actorPic := p.ProfileEntryResponse.ProfilePic
		savePic("actor", actorPic)
		fmt.Println("notifications for", username)
		pub58 := session.UsernameToPub58(username)
		t1 := time.Now().Unix()
		js := network.GetNotifications(pub58)
		t2 := time.Now().Unix()
		fmt.Println(t2 - t1)
		var list models.NotificationList
		json.Unmarshal([]byte(js), &list)
		m := map[string]int64{}

		for _, n := range list.Notifications {
			fromPub58 := n.Metadata.TransactorPublicKeyBase58Check
			if n.Metadata.TxnType == "CREATOR_COIN" {
				cctm := n.Metadata.CreatorCoinTxindexMetadata
				if cctm.OperationType == "buy" {
					fmt.Println("BUY", cctm.OperationType, cctm)
					m[fromPub58] += cctm.BitCloutToSellNanos
				} else {
					fmt.Println("SELL", cctm.OperationType, cctm)
				}
			} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
				md := n.Metadata.CreatorCoinTransferTxindexMetadata
				if md.PostHashHex != "" {
					continue
				}
				byUSD := ConvertToUSD(r, md.CreatorCoinToTransferNanos)

				if byUSD < 9.0 {
					fmt.Println("price was only", byUSD)
					continue
				}
				if PostAboutTransfer(&list, username, fromPub58, md) {
					os.Exit(0)
				}
			}
		}
		for fromPub58, sum := range m {
			byUSD := ConvertToUSD(r, sum)
			if byUSD < 9.0 {
				fmt.Println("price was only", byUSD)
				continue
			}
			if FindPercentAndPost(&list, username, pub58, fromPub58, sum) {
				break
			}
		}

	}
}

func FindTopHolders(total int64, user *models.User, filter []string) string {
	filterMap := map[string]bool{}
	for _, f := range filter {
		filterMap[f] = true
	}

	top := []string{}
	sort.SliceStable(user.UsersWhoHODLYou, func(i, j int) bool {
		return user.UsersWhoHODLYou[i].BalanceNanos >
			user.UsersWhoHODLYou[j].BalanceNanos
	})
	friendMap := map[string]int{}
	for i, friend := range user.UsersWhoHODLYou {
		per := int((float64(friend.BalanceNanos) / float64(total)) * 100.0)
		username := friend.ProfileEntryResponse.Username
		friendMap[username] = per
		if filterMap[username] == false {
			top = append(top, username)
		}
		if i >= 8 {
			break
		}
	}
	ChartIt(friendMap)
	if len(top) == 1 {
		return "@" + top[0]
	} else if len(top) == 2 {
		return "@" + top[0] + " @" + top[1]
	} else if len(top) > 2 {
		return "@" + top[0] + " @" + top[1] + " @" + top[2]
	}
	return ""
}

func PostAboutTransfer(list *models.NotificationList, username, fromPub58 string, md models.CreatorCoinTransferTxindexMetadata) bool {
	pub58 := session.UsernameToPub58(md.CreatorUsername)
	t1 := time.Now().Unix()
	fmt.Println("PostAboutTransfer", username, fromPub58, md.CreatorUsername)
	user := session.Pub58ToUser(pub58)
	t2 := time.Now().Unix()
	fmt.Println("PostAboutTransfer", t2-t1)
	from := list.ProfilesByPublicKey[fromPub58].Username
	if from == "" {
		return false
	}
	if from == md.CreatorUsername {
		return false
	}
	if alreadyDone[from] {
		return false
	}
	if alreadyDone[md.CreatorUsername] {
		return false
	}
	fromPic := list.ProfilesByPublicKey[fromPub58].ProfilePic
	savePic("from", fromPic)
	coinPic := user.ProfileEntryResponse.ProfilePic
	savePic("coin", coinPic)
	total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos
	topMention := FindTopHolders(total, &user, []string{from, username, md.CreatorUsername})
	for _, friend := range user.UsersWhoHODLYou {

		if friend.ProfileEntryResponse.Username == username {

			per := float64(friend.BalanceNanos) / float64(total)
			if per >= 0.01 {
				byUSD := ConvertToUSD(r, md.CreatorCoinToTransferNanos)

				perString := fmt.Sprintf("%d", int(per*100))
				text := fmt.Sprintf("This just in, there was a TRANSFER! @%s transfered %d ($%0.2f USD) of @%s to @%s and, according to our research, they now own %s%%. Enrich followers would love to hear the back story. You could re-clout this and explain...\\n\\ncc %s you have a new co-holder.", from, md.CreatorCoinToTransferNanos, byUSD, md.CreatorUsername, username, perString, topMention)
				fmt.Println(text)
				exec.Command("montage", "actor.webp", "from.webp", "chart.png", "coin.webp", "-tile", "4x1",
					"-geometry", "+0+0", "out.png").CombinedOutput()

				if argMap["live"] != "" {
					m := map[string]string{"text": text, "image": "out.png"}
					Post(m)
				}
				os.Exit(0)
				return true
			}
		}
	}

	return false
}

func FindPercentAndPost(list *models.NotificationList, username, pub58, fromPub58 string, sum int64) bool {
	t1 := time.Now().Unix()
	fmt.Println("FindPercentAndPost", username, pub58, fromPub58)
	user := session.Pub58ToUser(pub58)
	t2 := time.Now().Unix()
	fmt.Println("FindPercentAndPost", t2-t1)
	from := list.ProfilesByPublicKey[fromPub58].Username
	if from == "" {
		return false
	}
	if alreadyDone[from] {
		return false
	}
	fromPic := list.ProfilesByPublicKey[fromPub58].ProfilePic
	savePic("from", fromPic)
	total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos

	topMention := FindTopHolders(total, &user, []string{from, username})
	for _, friend := range user.UsersWhoHODLYou {

		if friend.ProfileEntryResponse.Username == from {

			per := float64(friend.BalanceNanos) / float64(total)
			if per >= 0.01 {

				byUSD := ConvertToUSD(r, sum)
				perString := fmt.Sprintf("%d", int(per*100))
				text := fmt.Sprintf("This just in, there was a BUY! @%s spent %d ($%0.2f USD) to BUY @%s and, according to our research, they now own %s%%. Enrich followers would love to hear the back story. You could re-clout this and explain...\\n\\ncc %s your %% may have changed.", from, sum, byUSD, username, perString, topMention)
				fmt.Println(text)
				exec.Command("montage", "from.webp", "chart.png", "actor.webp", "-tile", "3x1",
					"-geometry", "+0+0", "out.png").CombinedOutput()
				//exec.Command("convert", "out.png", "-gravity", "center",
				//"-background", "black", "-extent", "400x250", "out2.png").CombinedOutput()

				if argMap["live"] != "" {
					m := map[string]string{"text": text, "image": "out.png"}
					Post(m)
				}
				os.Exit(0)
				return true
			}
		}
	}

	return false
}

func ConvertToUSD(r models.Rate, sum int64) float64 {
	bySatoshi := float64(r.SatoshisPerBitCloutExchangeRate) * display.OneE9Float(sum)
	byUSD := float64(r.USDCentsPerBitcoinExchangeRate) * bySatoshi
	return byUSD / 10000000000.0
}

func LoadEnrichMessages() map[string]bool {
	m := map[string]bool{}
	js := network.GetPostsForPublicKey("enrich")
	var ppk models.PostsPublicKey
	json.Unmarshal([]byte(js), &ppk)
	for _, p := range ppk.Posts {
		tokens := strings.Split(p.Body, " ")
		for _, token := range tokens {
			if strings.HasPrefix(token, "@") {
				if strings.HasSuffix(token, "?") {
					thing := token[1:]
					m[thing[:len(thing)-1]] = true
				} else {
					m[token[1:]] = true
				}
			}
		}
	}
	return m
}
func ChartIt(m map[string]int) {

	items := []chart.Value{}
	fixed := map[string]int{}
	sum := 0
	for k, v := range m {
		if v < 5 {
			sum += v
		} else {
			fixed[k] = v
		}
	}
	fixed["other"] = sum
	for k, v := range fixed {
		items = append(items, chart.Value{Value: float64(v), Label: k})
	}

	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: items,
	}

	f, _ := os.Create("chart.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}

func savePic(flavor, data string) {
	tokens := strings.Split(data, ",")
	decoded, _ := base64.StdEncoding.DecodeString(tokens[1])
	ioutil.WriteFile(flavor+".webp", decoded, 0755)
}
