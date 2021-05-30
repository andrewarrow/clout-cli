package main

import (
	"clout/network"
	"clout/session"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/tyler-smith/go-bip39/wordlists"
)

func HandleWords(argMap map[string]string) {
	filepath := argMap["file"]
	if filepath != "" {
		b, _ := ioutil.ReadFile(filepath)
		for _, word := range strings.Split(string(b), "\n") {
			if len(word) == 0 {
				continue
			}
			fmt.Println(word)
			m := map[string]string{}
			m["new"] = word
			session.HandleAccounts(m)
		}
	}
}

func SearchFor404s() {
	for _, word := range wordlists.English {
		fmt.Println(word)
		test := network.DoTest404(word)
		if test == "404" {
			fmt.Println("  " + test + " !!!!!!!!!!!!!")
		}
		time.Sleep(time.Second * 1)
	}

}
