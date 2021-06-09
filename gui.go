package main

import (
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

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
	list := []string{}
	for _, p := range ps.PostsFound {
		username := p.ProfileEntryResponse.Username
		list = append(list, username)
	}
	readyHTML := fmt.Sprintf(template, strings.Join(list, "<br/>"))
	url := "data:text/html," + url.PathEscape(readyHTML)

	debug := true
	w := webview.New(debug)
	w.SetTitle("cloutcli")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
