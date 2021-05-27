package session

import (
	"clout/files"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func HandleAccounts() {
	if len(os.Args) > 2 {
		username := os.Args[2]
		i, _ := strconv.Atoi(username)
		if i > 0 {
			list := ReadAccountsSorted()
			username = list[i-1]
		}
		home := files.UserHomeDir()
		path := home + "/" + dir + "/" + selected
		ioutil.WriteFile(path, []byte(username), 0700)
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
