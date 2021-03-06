package session

import (
	"clout/files"
	"clout/keys"
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

var dir = "clout-cli-data"
var file = "secrets.txt"
var selected = "selected.account"
var cache = "cache.usernames"
var short = "short.map"
var backup = "clout.enc"
var baseline = "baseline.json"
var tags = "tags.json"

func JustReadFile(s string) string {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/" + s
	b, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(b))
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

func SelectedAccount() string {
	return JustReadFile(selected)
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
func NewWords() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}

func Pub58ToUser(key string) models.User {
	js := network.GetUsersStateless(key)
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	return us.UserList[0]
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

func LoggedInAs() (string, string, int64, string) {

	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return "", "", 0, ""
	}
	seedBytes := SeedBytes(mnemonic)
	//fmt.Printf("%x\n", seedBytes)

	pub58, _, btc := keys.ComputeKeysFromSeedWithAddress(seedBytes)
	username, balance := Pub58ToUsername(pub58)
	return pub58, username, balance, btc
}

func UsernameFromSecret(s string) string {
	seedBytes := SeedBytes(s)
	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	username, _ := Pub58ToUsername(pub58)
	return username
}
