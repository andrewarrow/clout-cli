package session

import (
	"clout/files"
	"clout/keys"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func WriteSelected(username string) {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/" + selected
	ioutil.WriteFile(path, []byte(username), 0700)
}

func HandleAccounts(argMap map[string]string) {
	if argMap["new"] != "" {
		words := NewWords()
		_, _, btc := keys.ComputeKeysFromSeedWithAddress(SeedBytes(words))
		fmt.Println("SAVE THESE WORDS!")
		fmt.Println("")
		fmt.Println(words)
		fmt.Println("")
		fmt.Println("We use `add-oath` command in github.com/andrewarrow/wolfservers.")
		fmt.Println("")
		fmt.Println("Send >= 0.0007 BTC to", btc)
		fmt.Println("")
		fmt.Println("Wait a few minutes then run `clout balance`")
		username := argMap["new"]
		usernames := ReadAccounts()
		usernames[username] = words
		WriteAccounts(usernames)
		WriteSelected(username)
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
	username := SelectedAccount()
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
