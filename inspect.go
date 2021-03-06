package main

import (
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/webview/webview"
)

func HandleInspect() {
	if len(os.Args) < 3 {
		fmt.Println("username")
		return
	}

	username := os.Args[2]
	if argMap["scrape"] != "" {
		js := network.GetSingleProfile(username)
		var sp models.SingleProfile
		json.Unmarshal([]byte(js), &sp)
		tokens := strings.Split(sp.Profile.Description, "\n")
		for _, token := range tokens {
			tokens = strings.Split(token, " ")
			for _, token := range tokens {
				fmt.Println(token)

				if strings.Contains(token, "twitter.com/") {
					HandleTwitterGrab(token)
				}
			}
		}

		return
	}

	GuiShowNotifications(username)
}

var w webview.WebView

func HandleTwitterGrab(url string) {
	debug := true
	w = webview.New(debug)
	w.SetTitle("cloutcli")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Bind("sendBackBodyInnerHTML", twitterCallback)

	w.Dispatch(func() {
		go func() {
			time.Sleep(time.Second * 4)
			w.Eval("sendBackBodyInnerHTML(document.body.innerHTML);")
		}()
	})
	w.Run()
}

func twitterCallback(data string) {
	tokens := strings.Split(data, "followers")
	tokens = strings.Split(tokens[1], ">")
	tokens = strings.Split(tokens[3], "<")
	fmt.Println("followers:", tokens[0])
	w.Terminate()
	//w.Destroy()
}
