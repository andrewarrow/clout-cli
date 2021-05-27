package session

import (
	"bufio"
	"clout/files"
	"clout/keys"
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

var dir = "clout-cli-data"
var file = "secrets.txt"
var selected = "selected.account"
var cache = "cache.usernames"

func JustReadFile(s string) string {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/" + s
	b, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(b))
}

func HandleAccounts() {
	if len(os.Args) > 2 {
		username := os.Args[2]
		home := files.UserHomeDir()
		path := home + "/" + dir + "/" + selected
		ioutil.WriteFile(path, []byte(username), 0700)
		return
	}
	m := ReadAccounts()
	fmt.Println("")
	for k, _ := range m {
		fmt.Printf("%s\n", k)
	}
	fmt.Println("")
	fmt.Println("To select account, run `clout account [username]`")
	fmt.Println("")
}

func Whoami() string {
	fmt.Println("Logged in as:")
	fmt.Println("")
	pub58, username, balance := LoggedInAs()
	fmt.Println(pub58)
	fmt.Println(username, balance)
	fmt.Println("")
	return username
}

func SecretFileExists() bool {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/" + file
	_, e := ioutil.ReadFile(path)
	if e != nil {
		return false
	}
	return true
}

func ReadLoggedInWords() string {
	m := ReadAccounts()
	if len(m) == 0 {
		fmt.Println("    --- not logged in yet, run clout login")
		return ""
	}
	username := JustReadFile(selected)
	for k, v := range m {
		if k == username {
			return v
		}
	}
	for _, v := range m {
		return v
	}
	return ""
}
func Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter mnenomic: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	_, e := bip39.MnemonicToByteArray(text)
	if e != nil {
		fmt.Println(e)
		return
	}
	//fmt.Printf("%x\n", b)

	username := UsernameFromSecret(text)
	usernames := ReadAccounts()
	usernames[username] = text
	WriteAccounts(usernames)

	//fmt.Println("")
	//Whoami()
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

func Logout() {
	m := ReadAccounts()
	if len(m) == 0 {
		return
	}
	if len(m) == 1 {
		home := files.UserHomeDir()
		path := home + "/" + dir + "/" + file
		os.Remove(path)
		fmt.Println("Secret removed.")
		fmt.Println("")
		return
	}
	username := JustReadFile(selected)
	if username == "" {
		fmt.Println("Please run `clout account [username]` to select account first.")
		return
	}
	delete(m, username)
	WriteAccounts(m)
}

func LoggedInPub58() string {
	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return ""
	}
	seedBytes := SeedBytes(mnemonic)
	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	return pub58
}
func SeedBytes(mnemonic string) []byte {
	//entropy, _ := bip39.NewEntropy(128)
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	//b, _ := bip39.MnemonicToByteArray(mnemonic)
	seedBytes, _ := bip39.NewSeedWithErrorChecking(mnemonic, "")
	//fmt.Printf("\n\nPRIVATE\n%x\n\n", seedBytes)
	return seedBytes
}

func Pub58ToUsername(key string) (string, int64) {
	js := network.GetUsersStateless(key)
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	return us.UserList[0].ProfileEntryResponse.Username, us.UserList[0].BalanceNanos
}
func UsernameToPub58(s string) string {
	js := network.GetSingleProfile(s)
	var sp models.SingleProfile
	json.Unmarshal([]byte(js), &sp)
	return sp.Profile.PublicKeyBase58Check
}

func LoggedInAs() (string, string, int64) {

	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return "", "", 0
	}
	seedBytes := SeedBytes(mnemonic)
	//fmt.Printf("%x\n", seedBytes)

	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	username, balance := Pub58ToUsername(pub58)
	return pub58, username, balance
}

func UsernameFromSecret(s string) string {
	seedBytes := SeedBytes(s)
	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	username, _ := Pub58ToUsername(pub58)
	return username
}
