package main

import (
	"bufio"
	"clout/files"
	"clout/keys"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

var dir = "clout-cli-data"
var file = "secrets.txt"

func JustReadFile() string {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/" + file
	b, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(b))
}

func ListAccounts() {
}

func Whoami() string {
	fmt.Println("Logged in as:")
	fmt.Println("")
	pub58, username := LoggedInAs()
	fmt.Println(pub58)
	fmt.Println(username)
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
	asBytes := []byte(JustReadFile())
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
	home := files.UserHomeDir()
	path := home + "/" + dir + "/" + file
	os.Remove(path)
	fmt.Println("Secret removed.")
	fmt.Println("")
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

func LoggedInAs() (string, string) {

	mnemonic := ReadLoggedInWords()
	if mnemonic == "" {
		return "", ""
	}
	seedBytes := SeedBytes(mnemonic)
	//fmt.Printf("%x\n", seedBytes)

	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	return pub58, Pub58ToUsername(pub58)
}

func UsernameFromSecret(s string) string {
	seedBytes := SeedBytes(s)
	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	return Pub58ToUsername(pub58)
}
