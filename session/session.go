package session

import (
	"clout/display"
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

func JustReadFile(s string) string {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/" + s
	b, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(b))
}

func Whoami() string {
	fmt.Println("Logged in as:")
	fmt.Println("")
	pub58, username, balance, btc := LoggedInAs()
	fmt.Println(display.LeftAligned("clout pub58", 20), pub58)
	fmt.Println(display.LeftAligned("btc address", 20), btc)
	fmt.Println(display.LeftAligned("clout username", 20), username)
	fmt.Println(display.LeftAligned("clout balance", 20), balance)
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

func Pub58ToBoards(key string) {
	js := network.GetUsersStateless(key)
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	for _, thing := range us.UserList[0].UsersYouHODL {
		coins := float64(thing.BalanceNanos) / 1000000000.0
		if coins < 1 {
			continue
		}
		fmt.Printf("%s %0.2f\n",
			display.LeftAligned(thing.ProfileEntryResponse.Username, 30), coins)

		other := thing.ProfileEntryResponse.PublicKeyBase58Check
		js = network.GetUsersStateless(other)
		var us models.UsersStateless
		json.Unmarshal([]byte(js), &us)
		for _, friend := range us.UserList[0].UsersWhoHODLYou {
			if friend.ProfileEntryResponse.PublicKeyBase58Check == key {
				continue
			}
			coins := float64(friend.BalanceNanos) / 1000000000.0
			if coins < 1 {
				continue
			}
			username := friend.ProfileEntryResponse.Username
			if username == "" {
				username = "anonymous"
			}
			fmt.Printf("  %s %0.2f\n",
				display.LeftAligned(username, 30), coins)
		}
	}
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
