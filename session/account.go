package session

import (
	"clout/files"
	"clout/keys"
	"clout/network"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func WriteSelected(username string) {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/" + selected
	ioutil.WriteFile(path, []byte(username), 0700)
}

func GetAccountsForTag(tag string) []string {
	m := ReadTags()
	tagMap := map[string][]string{}
	for k, v := range m {
		tokens := strings.Split(v, ",")
		for _, token := range tokens {
			tagMap[token] = append(tagMap[token], k)
		}
	}
	return tagMap[tag]
}

func HandleAccounts(argMap map[string]string) {
	username := argMap["new"]
	if username != "" {
		if network.DoTest404(username) == "200" {
			fmt.Printf("username `%s` is already in use.\n", username)
			return
		}
		words := NewWords()
		pub58, _, _ := keys.ComputeKeysFromSeedWithAddress(SeedBytes(words))
		fmt.Println("New words added to your secrets.txt file, back it up.")
		fmt.Println("")
		fmt.Println("New BITCLOUT address is")
		fmt.Println(pub58)
		fmt.Println("")
		fmt.Println("Secure this username, send clout to that new address then run")
		fmt.Println("")
		fmt.Printf("./clout update --username=%s\n\n\n", username)
		usernames := ReadAccounts()
		usernames[username] = words
		WriteAccounts(usernames)
		WriteSelected(username)
		return
	}
	if argMap["tag"] != "" {
		m := ReadTags()

		username := SelectedAccount()
		list := []string{}
		if m[username] != "" {
			list = strings.Split(m[username], ",")
		}
		list = append(list, argMap["tag"])
		m[username] = strings.Join(list, ",")
		WriteTags(m)
		return
	}
	if argMap["query"] != "" {
		for _, username := range GetAccountsForTag(argMap["query"]) {
			fmt.Println(username)
		}
		return
	}
	if len(os.Args) > 2 {
		username := os.Args[2]
		i, _ := strconv.Atoi(username)
		if i > 0 {
			list := ReadAccountsSorted()
			username = list[i-1]
		}
		WriteSelected(username)
	}
	fmt.Println("")
	for i, k := range ReadAccountsSorted() {
		fmt.Printf("%02d. %s\n", i+1, k)
	}
	fmt.Println("")
	fmt.Println("To select account, run `clout account [username or i]`")
	fmt.Println("")
	username = SelectedAccount()
	if username != "" {
		fmt.Println("SELECTED ACCOUNT:", username)
		fmt.Println("")
	}
}

func ReadAccountsSorted() []string {
	m := ReadAccounts()
	buff := []string{}
	for k, _ := range m {
		buff = append(buff, k)
	}
	sort.Strings(buff)
	return buff
}
func ReadAccounts() map[string]string {
	m := map[string]string{}
	asBytes := []byte(JustReadFile(file))
	if len(asBytes) == 0 {
		return m
	}

	json.Unmarshal(asBytes, &m)

	return m
}

func WriteAccounts(m map[string]string) {
	b, _ := json.Marshal(m)
	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/" + file
	ioutil.WriteFile(path, b, 0700)
	fmt.Println("Secret stored at:", path)
}
func ReadTags() map[string]string {
	m := map[string]string{}
	asBytes := []byte(JustReadFile(tags))
	if len(asBytes) == 0 {
		return m
	}

	json.Unmarshal(asBytes, &m)

	return m
}

func WriteTags(m map[string]string) {
	b, _ := json.Marshal(m)
	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/" + tags
	ioutil.WriteFile(path, b, 0700)
	fmt.Println("Stored at:", path)
}
