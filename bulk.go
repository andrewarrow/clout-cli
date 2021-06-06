package main

import (
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func HandleBulk() {

	file := argMap["file"]
	if file != "" {
		b, _ := ioutil.ReadFile(file)
		s := string(b)
		lines := strings.Split(s, "\n")
		for _, line := range lines {
			fmt.Println(line)
			//os.Args = []string{"", "diamond", line}
			//HandleDiamond()
		}
		return
	}

	query := argMap["query"]
	if query == "" {
		return
	}

	//changeme := ""
	for _, username := range session.GetAccountsForTag(query) {
		fmt.Println(username)
		session.WriteSelected(username)

		//name, url := VideoFromVimeo(username)
		//m := map[string]string{"text": name, "video": url}
		//Post(m)

		//m := map[string]string{"percent": "3333"}
		//HandleUpdateProfile(m)
		//os.Args = []string{"", "follow", ""}
		//HandleFollow()
		//os.Args = []string{"", "reclout", changeme}
		//HandleReclout()
		//m := map[string]string{"text": "everyone? even me?", "reply": changeme}
		//Post(m)
		//m := map[string]string{"hash": "changeme"}
		//HandleLike(m)
		time.Sleep(time.Second * 1)
	}

}

type VimeoReply struct {
	Data []VimeoRecord
}

type VimeoRecord struct {
	Uri  string
	Name string
}

func VideoFromVimeo(query string) (string, string) {
	pat := os.Getenv("VIMEO_PAT")
	url := "https://api.vimeo.com/videos?query=" + query
	js := network.DoGetWithPat(pat, url)
	var vr VimeoReply
	json.Unmarshal([]byte(js), &vr)

	fmt.Println(vr.Data[0].Uri)
	fmt.Println(vr.Data[0].Name)

	tokens := strings.Split(vr.Data[0].Uri, "/")
	id := tokens[len(tokens)-1]

	return vr.Data[0].Name, "https://player.vimeo.com/video/" + id
}
