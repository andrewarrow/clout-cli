package main

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/webview/webview"
)

func ListPostsWithGui(follow bool) {
	template := `
  <html>
    <head><title>Hello</title></head>
		<body style="font-family: courier; background-color: black; color: green;">%s</body>
  </html>`

	pub58 := session.LoggedInPub58()
	js := network.GetPostsStateless(pub58, follow)
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)
	html := "<table>"
	for _, p := range ps.PostsFound {
		html += "<tr>"
		username := p.ProfileEntryResponse.Username
		html += "<td>" + username + "</td>"
		html += "<td>"
		for _, image := range p.ImageURLs {
			html += fmt.Sprintf("<img height='100' src='%s'/>", image)
		}
		html += "</td>"
		html += "<td><div style='width: 400px;'>" + p.Body
		html += "</div></td>"
		html += "</tr>"
	}
	html += "</table>"
	readyHTML := fmt.Sprintf(template, html)
	url := "data:text/html," + url.PathEscape(readyHTML)

	debug := false
	w := webview.New(debug)
	w.SetTitle("cloutcli")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
