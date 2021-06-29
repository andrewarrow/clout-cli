package main

import (
	"clout/display"
	"clout/draw"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/fogleman/gg"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

var alreadyDone map[string]bool
var r models.Rate
var actorPub58 string

func FindBuysSellsAndTransfers() {
	session.WriteSelected("enrich")
	js := network.GetExchangeRate()
	json.Unmarshal([]byte(js), &r)

	if argMap["test"] != "" {
		TestBigImage()
		return
	}

	alreadyDone = LoadEnrichMessages()
	alreadyDone["douglasss"] = true
	alreadyDone["enrich"] = true
	fmt.Println(alreadyDone)
	pub58 := session.LoggedInPub58()
	last := ""
	for {
		js := network.GetPostsStatelessWithOptions(last, pub58)
		var ps models.PostsStateless
		json.Unmarshal([]byte(js), &ps)

		fmt.Printf("PostsFound %d\n", len(ps.PostsFound))
		FindBuysSellsAndTransfersFromPosts(ps.PostsFound)
		time.Sleep(time.Second * 1)
		for _, p := range ps.PostsFound {
			last = p.PostHashHex
		}
	}
}

func SaveImagesToDisk(urls []string) {
	for i, url := range urls {
		fmt.Println(url)
		filename := fmt.Sprintf("%03d", i+2)
		if url == "" {
			exec.Command("cp", "001.png", filename+".png").Output()
			continue
		}
		jsonString := network.DoGetWithPat("", url)
		ioutil.WriteFile(filename+".webp", []byte(jsonString), 0755)
		exec.Command("convert", filename+".webp", filename+".png").Output()
		os.Remove(filename + ".webp")
	}
}
func ImagesFromPosts(username string) {
	js := network.GetPostsForPublicKey(username)
	var ppk models.PostsPublicKey
	json.Unmarshal([]byte(js), &ppk)
	urls := []string{}
	for _, p := range ppk.Posts {
		urls = append(urls, p.ImageURLs...)
		urls = append(urls, "")
		if p.RecloutedPostEntryResponse != nil {
			urls = append(urls, p.RecloutedPostEntryResponse.ImageURLs...)
			urls = append(urls, "")
			if p.RecloutedPostEntryResponse.RecloutedPostEntryResponse != nil {
				urls = append(urls, p.RecloutedPostEntryResponse.RecloutedPostEntryResponse.ImageURLs...)
				urls = append(urls, "")
			}
		}
	}
	SaveImagesToDisk(urls)
}

func TestBigImage() {
	if true {
		draw.BuyPoster("spektr", "brockchain", 244.1)
		return
	}
	friendMap := map[string]int{}
	friendMap["username"] = 15
	friendMap["testing"] = 15
	friendMap["ouch"] = 15
	friendMap["other"] = 55
	ChartIt(friendMap)
	username := "guttercatsofbitclout"
	from := "purchaser"
	pub58 := "BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v"
	actorBytes := network.GetSingleProfilePicture(pub58)
	draw.SavePic("actor", actorBytes)

	fromPub58 := "BC1YLj2V95AZ3kuNKC59BJ1Mj99jQiZBJy1Dz7gPG1AcLjNfZgMa2nt"
	fromBytes := network.GetSingleProfilePicture(fromPub58)
	draw.SavePic("from", fromBytes)

	coinPub58 := "BC1YLgdpjLNf96dCvBpa9X9eTTdMDxreTs6Z5sWC2b4vQ1L1SAsmeEP"
	coinBytes := network.GetSingleProfilePicture(coinPub58)
	draw.SavePic("coin", coinBytes)

	sum := int64(10000000000)
	byUSD := ConvertToUSD(r, sum)
	//usdPerFollower := byUSD / float64(numFollowers)
	per := 0.20
	perString := fmt.Sprintf("%d", int(per*100))

	//BigImageBuy(fmt.Sprintf("$%0.2f", byUSD), username, 36, perString+"%", from)
	BigImageTransfer(fmt.Sprintf("$%0.2f", byUSD), username, 36, perString+"%", from, from)
	//ImagesFromPosts("artsyminal")
	//MakeVideoFromImages()
}

func FindBuysSellsAndTransfersFromPosts(found []models.Post) {
	for _, p := range found {
		username := p.ProfileEntryResponse.Username
		if alreadyDone[username] {
			continue
		}

		fmt.Println("notifications for", username)
		pub58 := session.UsernameToPub58(username)
		actorPub58 = pub58

		numFollowers := GetNumFollowers(pub58, username)
		fmt.Println(numFollowers)

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

				if byUSD < 40.0 {
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
			if byUSD < 40.0 {
				fmt.Println("price was only", byUSD)
				continue
			}
			if FindPercentAndPost(&list, username, pub58, numFollowers, fromPub58, sum) {
				break
			}
		}

	}
}

func FindTopHodlers(total int64, hw *models.HodlersWrap, filter []string) string {
	filterMap := map[string]bool{}
	for _, f := range filter {
		filterMap[f] = true
	}

	top := []string{}
	sort.SliceStable(hw.Hodlers, func(i, j int) bool {
		return hw.Hodlers[i].BalanceNanos >
			hw.Hodlers[j].BalanceNanos
	})
	friendMap := map[string]int{}
	for i, friend := range hw.Hodlers {
		per := int((float64(friend.BalanceNanos) / float64(total)) * 100.0)
		username := friend.ProfileEntryResponse.Username
		if username == "" {
			username = "anonymous"
		}
		friendMap[username] = per
		if filterMap[username] == false && username != "anonymous" {
			top = append(top, username)
		}
		if i >= 8 {
			break
		}
	}
	ChartIt(friendMap)

	blessed := []string{}
	for _, t := range top {
		if alreadyDone[t] {
			continue
		}
		blessed = append(blessed, t)
	}
	if len(blessed) == 1 {
		return "@" + blessed[0]
	} else if len(blessed) == 2 {
		return "@" + blessed[0] + " @" + blessed[1]
	} else if len(blessed) > 2 {
		return "@" + blessed[0] + " @" + blessed[1] + " @" + blessed[2]
	}
	return ""
}

func PostAboutTransfer(list *models.NotificationList, username, fromPub58 string, md models.CreatorCoinTransferTxindexMetadata) bool {
	pub58 := session.UsernameToPub58(md.CreatorUsername)
	t1 := time.Now().Unix()
	fmt.Println("PostAboutTransfer", username, fromPub58, md.CreatorUsername)
	js := network.GetHodlers(md.CreatorUsername)
	var hw models.HodlersWrap
	json.Unmarshal([]byte(js), &hw)
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
	total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos
	topMention := FindTopHodlers(total, &hw, []string{from, username, md.CreatorUsername})
	for _, friend := range hw.Hodlers {

		if friend.ProfileEntryResponse.Username == username {

			per := float64(friend.BalanceNanos) / float64(total)
			if per >= 0.01 {
				numFollowers := GetNumFollowers(pub58, md.CreatorUsername)
				actorBytes := network.GetSingleProfilePicture(actorPub58)
				draw.SavePic("actor", actorBytes)
				fromBytes := network.GetSingleProfilePicture(fromPub58)
				draw.SavePic("from", fromBytes)
				coinBytes := network.GetSingleProfilePicture(user.PublicKeyBase58Check)
				draw.SavePic("coin", coinBytes)
				byUSD := ConvertToUSD(r, md.CreatorCoinToTransferNanos)

				perString := fmt.Sprintf("%d", int(per*100))
				text := fmt.Sprintf("TRANSFER! @%s transfered %d ($%0.2f USD) of @%s to @%s\\n\\ncc %s you have a new co-holder.", from, md.CreatorCoinToTransferNanos, byUSD, md.CreatorUsername, username, topMention)
				fmt.Println(text)
				//exec.Command("montage", "actor.webp", "from.webp", "chart.png", "coin.webp", "-tile", "4x1", "-geometry", "+0+0", "out.png").CombinedOutput()
				BigImageTransfer(fmt.Sprintf("$%0.2f", byUSD), md.CreatorUsername, numFollowers, perString+"%", from, username)
				//ImagesFromPosts(md.CreatorUsername)
				//MakeVideoFromImages()

				if argMap["live"] != "" {
					m := map[string]string{"text": text, "image": "001.png"}
					Post(m)
				}
				os.Exit(0)
				return true
			}
		}
	}

	return false
}

func FindPercentAndPost(list *models.NotificationList, username, pub58 string,
	numFollowers int64, fromPub58 string, sum int64) bool {
	t1 := time.Now().Unix()
	fmt.Println("FindPercentAndPost", username, pub58, numFollowers, fromPub58)
	user := session.Pub58ToUser(pub58)
	js := network.GetHodlers(username)
	var hw models.HodlersWrap
	json.Unmarshal([]byte(js), &hw)

	t2 := time.Now().Unix()
	fmt.Println("FindPercentAndPost", t2-t1)
	from := list.ProfilesByPublicKey[fromPub58].Username
	if from == "" {
		fmt.Println("no username")
		return false
	}
	if alreadyDone[from] {
		return false
	}
	total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos

	topMention := FindTopHodlers(total, &hw, []string{from, username})
	for _, friend := range hw.Hodlers {

		if friend.ProfileEntryResponse.Username == from {

			per := float64(friend.BalanceNanos) / float64(total)
			if per >= 0.01 {
				actorBytes := network.GetSingleProfilePicture(pub58)
				draw.SavePic("actor", actorBytes)

				fromBytes := network.GetSingleProfilePicture(fromPub58)
				draw.SavePic("from", fromBytes)
				byUSD := ConvertToUSD(r, sum)
				//usdPerFollower := byUSD / float64(numFollowers)
				perString := fmt.Sprintf("%d", int(per*100))

				text := fmt.Sprintf("BUY! @%s spent %d ($%0.2f USD) to BUY @%s\\n\\ncc %s your %% may have changed.", from, sum, byUSD, username, topMention)
				fmt.Println(text)

				//draw.BuyPoster(username, from, byUSD)
				BigImageBuy(fmt.Sprintf("$%0.2f", byUSD), username, numFollowers, perString+"%", from)
				//ImagesFromPosts(username)
				//MakeVideoFromImages()

				//exec.Command("montage", "from.webp", "chart.png", "actor.webp", "-tile", "3x1",
				//"-geometry", "+0+0", "out.png").CombinedOutput()
				//exec.Command("convert", "out.png", "-gravity", "center",
				//"-background", "black", "-extent", "400x250", "out2.png").CombinedOutput()

				if argMap["live"] != "" {
					m := map[string]string{"text": text, "image": "001.png"}
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
		tokens := strings.Split(p.Body, "\n")
		for _, token := range tokens {
			tokens = strings.Split(token, " ")
			for _, token := range tokens {
				if strings.HasPrefix(token, "@") {
					if strings.HasSuffix(token, ".") {
						thing := token[1:]
						m[thing[:len(thing)-1]] = true
					} else {
						m[token[1:]] = true
					}
				}
			}
		}
	}
	return m
}

var (
	colorWhite          = drawing.Color{R: 241, G: 241, B: 241, A: 255}
	colorMariner        = drawing.Color{R: 60, G: 100, B: 148, A: 255}
	colorLightSteelBlue = drawing.Color{R: 182, G: 195, B: 220, A: 255}
)

func ChartIt(m map[string]int) {

	items := []chart.Value{}
	bars := []chart.StackedBar{}
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
	barWidth := 80
	style1 := chart.Style{}
	style1.StrokeWidth = 0.01
	style1.FillColor = colorMariner
	style1.FontColor = colorWhite
	style2 := chart.Style{}
	style2.StrokeWidth = 0.01
	style2.FillColor = colorLightSteelBlue
	style2.FontColor = colorWhite
	for k, v := range fixed {
		items = append(items, chart.Value{Value: float64(v), Label: k})

		sb := chart.StackedBar{}
		sb.Width = barWidth

		values := []chart.Value{}
		value := chart.Value{}
		value.Label = k
		value.Value = 100.0 - float64(v)
		value.Style = style1
		values = append(values, value)
		value = chart.Value{}
		value.Label = k
		value.Value = float64(v)
		value.Style = style2
		values = append(values, value)

		sb.Values = values
	}

	stackedBarChart := chart.StackedBarChart{
		Width:      800,
		Height:     600,
		BarSpacing: 40,
		Bars:       bars}

	stackedBarChart.IsHorizontal = true

	f, _ := os.Create("chart.png")
	defer f.Close()
	stackedBarChart.Render(chart.PNG, f)
}

func DrawUser(dc *gg.Context, file string, x, y float64, top, middle, bottom string) {
	font := "arial.ttf"
	dc.LoadFontFace(font, 48)
	im, _ := gg.LoadImage(file)
	dc.DrawImage(im, int(x), int(y))
	dc.SetLineWidth(2)
	dc.DrawRectangle(x, y, 100, 100)
	dc.Stroke()

	dc.LoadFontFace(font, 24)
	if len(top) > 9 {
		dc.LoadFontFace(font, 18)
	}
	dc.DrawStringAnchored(top, x+50, y+140-25, 0.5, 0.5)
	dc.LoadFontFace(font, 18)
	dc.DrawStringAnchored(middle, x+50, y+165-25, 0.5, 0.5)
	dc.DrawStringAnchored(bottom, x+50, y+165, 0.5, 0.5)
}
func BigImageBuy(price, coin string, numFollowers int64, percent, from string) {
	dc := gg.NewContext(600, 600)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	font := "arial.ttf"
	dc.LoadFontFace(font, 48)
	dc.DrawStringAnchored(price, 275+25, 45+50, 0.5, 0.5)
	dc.LoadFontFace(font, 48)
	dc.DrawStringAnchored("BUY", 275+25, 100+50, 0.5, 0.5)

	im, _ := gg.LoadImage("chart.png")
	dc.DrawImage(im, 30, 175)
	im, _ = gg.LoadImage("logo.png")
	dc.DrawImage(im, -40, -10+50)
	DrawUser(dc, "actor.png", 400+50, 25+50, coin, fmt.Sprintf("%d followers", numFollowers), "")

	DrawUser(dc, "from.png", 460, 250, "purchaser", from, fmt.Sprintf("owns %s", percent))
	dc.SavePNG("001.png")
}
func BigImageTransfer(price, coin string, numFollowers int64, percent, from, actor string) {
	dc := gg.NewContext(600, 600)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	font := "arial.ttf"
	dc.LoadFontFace(font, 48)
	dc.DrawStringAnchored(price, 275+25, 45+50, 0.5, 0.5)
	dc.LoadFontFace(font, 48)
	dc.DrawStringAnchored("TRANSFER", 275+25, 100+50, 0.5, 0.5)

	im, _ := gg.LoadImage("chart.png")
	dc.DrawImage(im, 30, 175)
	im, _ = gg.LoadImage("logo.png")
	dc.DrawImage(im, -40, -10+50)
	DrawUser(dc, "coin.png", 400+50, 50, coin, fmt.Sprintf("%d followers", numFollowers), "")

	DrawUser(dc, "actor.png", 460, 225, "receiver", actor, fmt.Sprintf("owns %s", percent))
	DrawUser(dc, "from.png", 460, 250+160, "giver", from, "")
	dc.SavePNG("001.png")
}
