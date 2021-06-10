package main

import (
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

func FindBuysSellsAndTransfers() {
	pub58 := session.LoggedInPub58()
	js := network.GetPostsStateless(pub58, false)
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	for _, p := range ps.PostsFound {
		username := p.ProfileEntryResponse.Username
		coinPic := p.ProfileEntryResponse.ProfilePic
		savePic("coin", coinPic)
		fmt.Println("notifications for", username)
		pub58 := session.UsernameToPub58(username)
		js := network.GetNotifications(pub58)
		var list models.NotificationList
		json.Unmarshal([]byte(js), &list)
		m := map[string]int64{}
		for _, n := range list.Notifications {
			fromPub58 := n.Metadata.TransactorPublicKeyBase58Check
			if n.Metadata.TxnType == "CREATOR_COIN" {
				cctm := n.Metadata.CreatorCoinTxindexMetadata
				if cctm.OperationType != "buy" {
					continue
				}
				m[fromPub58] += cctm.BitCloutToSellNanos
			}
		}
		for fromPub58, sum := range m {
			user := session.Pub58ToUser(pub58)
			from := list.ProfilesByPublicKey[fromPub58].Username
			if from == "" {
				continue
			}
			fromPic := list.ProfilesByPublicKey[fromPub58].ProfilePic
			savePic("from", fromPic)
			total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos

			sort.SliceStable(user.UsersWhoHODLYou, func(i, j int) bool {
				return user.UsersWhoHODLYou[i].BalanceNanos >
					user.UsersWhoHODLYou[j].BalanceNanos
			})
			friendMap := map[string]int{}
			for _, friend := range user.UsersWhoHODLYou {
				per := int((float64(friend.BalanceNanos) / float64(total)) * 100.0)
				friendMap[friend.ProfileEntryResponse.Username] = per
			}
			ChartIt(friendMap)
			for _, friend := range user.UsersWhoHODLYou {

				if friend.ProfileEntryResponse.Username == from {

					per := float64(friend.BalanceNanos) / float64(total)
					if per >= 0.01 {
						perString := fmt.Sprintf("%0.2f", per*100)
						text := fmt.Sprintf("@%s spends %d to buy @%s and now owns %s%%", from, sum, username, perString)
						fmt.Println(text)
						exec.Command("montage", "from.webp", "chart.png", "coin.webp", "-tile", "3x1",
							"-geometry", "+0+0", "out.png").CombinedOutput()
						//exec.Command("convert", "out.png", "-gravity", "center",
						//"-background", "black", "-extent", "400x250", "out2.png").CombinedOutput()

						//m := map[string]string{"text": text, "image": "/Users/aa/clout-cli/out2.png"}
						//Post(m)
					}
				}
			}
		}

	}
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
