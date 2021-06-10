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

	"github.com/wcharczuk/go-chart"
)

var alreadyDone map[string]bool
var r models.Rate

func FindBuysSellsAndTransfers() {
	js := network.GetExchangeRate()
	json.Unmarshal([]byte(js), &r)

	alreadyDone = LoadEnrichMessages()
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
		js := network.GetNotifications(pub58)
		var list models.NotificationList
		json.Unmarshal([]byte(js), &list)
		m := map[string]int64{}

		for _, n := range list.Notifications {
			fromPub58 := n.Metadata.TransactorPublicKeyBase58Check
			if n.Metadata.TxnType == "!CREATOR_COIN" {
				cctm := n.Metadata.CreatorCoinTxindexMetadata
				if cctm.OperationType != "buy" {
					continue
				}
				m[fromPub58] += cctm.BitCloutToSellNanos
			} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
				md := n.Metadata.CreatorCoinTransferTxindexMetadata
				if md.PostHashHex != "" {
					continue
				}
				if PostAboutTransfer(&list, username, fromPub58, md) {
					os.Exit(0)
				}
			}
		}
		for fromPub58, sum := range m {
			if FindPercentAndPost(&list, username, pub58, fromPub58, sum) {
				break
			}
		}

	}
}
func PostAboutTransfer(list *models.NotificationList, username, fromPub58 string, md models.CreatorCoinTransferTxindexMetadata) bool {
	pub58 := session.UsernameToPub58(md.CreatorUsername)
	user := session.Pub58ToUser(pub58)
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
	sort.SliceStable(user.UsersWhoHODLYou, func(i, j int) bool {
		return user.UsersWhoHODLYou[i].BalanceNanos >
			user.UsersWhoHODLYou[j].BalanceNanos
	})
	friendMap := map[string]int{}
	for i, friend := range user.UsersWhoHODLYou {
		per := int((float64(friend.BalanceNanos) / float64(total)) * 100.0)
		friendMap[friend.ProfileEntryResponse.Username] = per
		if i >= 8 {
			break
		}
	}
	ChartIt(friendMap)
	for _, friend := range user.UsersWhoHODLYou {

		if friend.ProfileEntryResponse.Username == username {

			per := float64(friend.BalanceNanos) / float64(total)
			if per >= 0.01 {
				byUSD := ConvertToUSD(r, md.CreatorCoinToTransferNanos)

				perString := fmt.Sprintf("%d", int(per*100))
				text := fmt.Sprintf("anything you can tell us @%s on why you transfered %d ($%0.2f USD) of @%s to @%s and they now own %s%%? Enrich followers want to know. You could re-clout this and explain...", from, md.CreatorCoinToTransferNanos, byUSD, md.CreatorUsername, username, perString)
				fmt.Println(text)
				exec.Command("montage", "actor.webp", "from.webp", "chart.png", "coin.webp", "-tile", "4x1",
					"-geometry", "+0+0", "out.png").CombinedOutput()

				//m := map[string]string{"text": text, "image": "/Users/aa/clout-cli/out.png"}
				//Post(m)
				os.Exit(0)
				return true
			}
		}
	}

	return false
}

func FindPercentAndPost(list *models.NotificationList, username, pub58, fromPub58 string, sum int64) bool {
	user := session.Pub58ToUser(pub58)
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

	sort.SliceStable(user.UsersWhoHODLYou, func(i, j int) bool {
		return user.UsersWhoHODLYou[i].BalanceNanos >
			user.UsersWhoHODLYou[j].BalanceNanos
	})
	friendMap := map[string]int{}
	for i, friend := range user.UsersWhoHODLYou {
		per := int((float64(friend.BalanceNanos) / float64(total)) * 100.0)
		friendMap[friend.ProfileEntryResponse.Username] = per
		if i >= 8 {
			break
		}
	}
	ChartIt(friendMap)
	for _, friend := range user.UsersWhoHODLYou {

		if friend.ProfileEntryResponse.Username == from {

			per := float64(friend.BalanceNanos) / float64(total)
			if per >= 0.01 {
				byUSD := ConvertToUSD(r, sum)

				perString := fmt.Sprintf("%d", int(per*100))
				text := fmt.Sprintf("anything you can tell us @%s on why you spent %d ($%0.2f USD) to buy @%s and now own %s%%? Enrich followers want to know. You could re-clout this and explain...", from, sum, byUSD, username, perString)
				fmt.Println(text)
				exec.Command("montage", "from.webp", "chart.png", "actor.webp", "-tile", "3x1",
					"-geometry", "+0+0", "out.png").CombinedOutput()
				//exec.Command("convert", "out.png", "-gravity", "center",
				//"-background", "black", "-extent", "400x250", "out2.png").CombinedOutput()

				//m := map[string]string{"text": text, "image": "/Users/aa/clout-cli/out.png"}
				//Post(m)
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
				m[token[1:]] = true
			}
		}
	}
	return m
}
func ChartIt(m map[string]int) {

	items := []chart.Value{}
	for k, v := range m {
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
