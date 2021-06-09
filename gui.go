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
	html += "</tr>"

	return html
}
