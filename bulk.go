package main

import (
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func HandleBulk() {

	query := argMap["query"]
	if query == "" {
		return
	}

	for _, username := range session.GetAccountsForTag(query) {
		fmt.Println(username)
		session.WriteSelected(username)

		VideoFromVimeo(username)
		//os.Args = []string{"", "follow", "changeme"}
		//HandleFollow()
		//os.Args = []string{"", "reclout", "changeme"}
		//HandleReclout()
		//m := map[string]string{"text": "we also like @derishaviar", "reply": "changeme"}
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

func VideoFromVimeo(query string) string {
	pat := os.Getenv("VIMEO_PAT")
	url := "https://api.vimeo.com/videos?query=" + query
	js := network.DoGetWithPat(pat, url)
	var vr VimeoReply
	json.Unmarshal([]byte(js), &vr)

	fmt.Println(vr.Data[0].Uri)
	fmt.Println(vr.Data[0].Name)

	return ""
}
