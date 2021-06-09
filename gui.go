package main

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/webview/webview"
)

var template = `
  <html>
    <head><title>Hello</title></head>
		<body style="font-family: courier; background-color: black; color: green;">%s</body>
  </html>`

func ListPostsWithGui(follow bool) {

	pub58 := session.LoggedInPub58()
	js := network.GetPostsStateless(pub58, follow)
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)
	html := "<table>"
	for _, p := range ps.PostsFound {
		html += GuiMakeRow("", &p)
		if p.RecloutedPostEntryResponse != nil {
			html += GuiMakeRow("RECLOUT", p.RecloutedPostEntryResponse)
		}
	}

	html += "</table>"
	readyHTML := fmt.Sprintf(template, html)
	url := "data:text/html," + url.PathEscape(readyHTML)

	debug := false
	w := webview.New(debug)
	w.SetTitle("cloutcli")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Dispatch(func() {
		go func() {
			time.Sleep(time.Second * 1)
			for _, p := range ps.PostsFound {
				js := fmt.Sprintf("document.getElementById('p%s').innerHTML='%s';",
					p.PostHashHex, p.Body)
				w.Eval(js)
				if p.RecloutedPostEntryResponse != nil {
					js := fmt.Sprintf("document.getElementById('p%s').innerHTML='%s';",
						p.RecloutedPostEntryResponse.PostHashHex, p.RecloutedPostEntryResponse.Body)
					w.Eval(js)
				}
			}
		}()
	})
	w.Run()
}

func GuiViewUser(username string) {
	debug := false
	w := webview.New(debug)
	w.SetTitle("cloutcli")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://bitclout.com/u/" + username)
	w.Run()
}

func GuiShowNotifications(username string) {
	pub58 := session.UsernameToPub58(username)
	js := network.GetNotifications(pub58)
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	html := "<table>"
	for i, n := range list.Notifications {
		if n.Metadata.TxnType == "BASIC_TRANSFER" {
			continue
		}
		html += GuiMakeRowNotification(i, &list, n)
	}

	html += "</table>"
	readyHTML := fmt.Sprintf(template, html)
	url := "data:text/html," + url.PathEscape(readyHTML)

	debug := false
	w := webview.New(debug)
	w.SetTitle("cloutcli")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Dispatch(func() {
		go func() {
			time.Sleep(time.Second * 1)
			for i, n := range list.Notifications {
				if n.Metadata.TxnType == "BASIC_TRANSFER" {
					continue
				}
				p := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check]
				js := fmt.Sprintf("document.getElementById('i%d').src='%s';", i,
					p.ProfilePic)
				w.Eval(js)
			}
		}()
	})
	w.Run()
}

func GuiMakeRow(flavor string, p *models.Post) string {
	html := "<tr>"
	if flavor != "" {
		html = "<tr style='background-color: red;'>"
	}
	username := p.ProfileEntryResponse.Username
	html += "<td>" + username + "</td>"
	html += "<td>"

	for _, image := range p.ImageURLs {
		html += fmt.Sprintf("<img width='200' src='%s'/>", image)
	}
	if p.PostExtraData.EmbedVideoURL != "" {
		src := p.PostExtraData.EmbedVideoURL
		html += fmt.Sprintf("<a style='color: white;' href='%s'>%s</a>", src, "video")
	}
	html += "</td>"
	html += fmt.Sprintf("<td><div id='p%s' style='width: 400px;'></div>", p.PostHashHex)
	html += "</td>"
	html += "</tr>"

	return html
}

func GuiMakeRowNotification(index int, list *models.NotificationList, n models.Notification) string {
	html := "<tr>"
	p := list.ProfilesByPublicKey[n.Metadata.TransactorPublicKeyBase58Check]
	from := p.Username
	html += fmt.Sprintf("<td><img id='i%d' src=''/></td>", index)
	html += "<td>" + from + "</td>"
	flavor := ""
	meta := ""
	amount := int64(0)
	coin := ""
	if n.Metadata.TxnType == "SUBMIT_POST" {
		p := list.PostsByHash[n.Metadata.SubmitPostTxindexMetadata.PostHashBeingModifiedHex]
		if p.Body == "" {
			flavor = "reclout"
			//meta = BodyParse(p.RecloutedPostEntryResponse.Body)
		} else {
			flavor = "mention"
			//meta = BodyParse(p.Body)
		}
	} else if n.Metadata.TxnType == "LIKE" {
		//p := list.PostsByHash[n.Metadata.LikeTxindexMetadata.PostHashHex]
		//meta = BodyParse(p.Body)
		flavor = "like"
	} else if n.Metadata.TxnType == "FOLLOW" {
		meta = from
		flavor = "follow"
	} else if n.Metadata.TxnType == "CREATOR_COIN_TRANSFER" {
		md := n.Metadata.CreatorCoinTransferTxindexMetadata
		if md.PostHashHex != "" {
			//p := list.PostsByHash[md.PostHashHex]
			//meta = fmt.Sprintf("%d ", md.DiamondLevel) + BodyParse(p.Body)
			amount = md.DiamondLevel
			flavor = "diamond"
		} else {
			meta = fmt.Sprintf("%s %d", md.CreatorUsername, md.CreatorCoinToTransferNanos)
			amount = md.CreatorCoinToTransferNanos
			coin = md.CreatorUsername
			flavor = "coin"
		}
	} else if n.Metadata.TxnType == "CREATOR_COIN" {
		cctm := n.Metadata.CreatorCoinTxindexMetadata
		if cctm.OperationType == "buy" {
			amount = cctm.BitCloutToSellNanos
		} else if cctm.OperationType == "sell" {
			amount = cctm.CreatorCoinToSellNanos
		}
		meta = fmt.Sprintf("%s %d", from, amount)
		flavor = cctm.OperationType
	}
	html += "<td>" + flavor + "</td>"
	html += fmt.Sprintf("<td>%d</td>", amount)
	html += "<td>" + coin + "</td>"
	html += "<td>" + meta + "</td>"
	html += "</tr>"

	return html
}
