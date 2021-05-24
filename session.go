package main

import (
	"bufio"
	"clout/files"
	"clout/keys"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

var dir = "clout-cli-data"

func Whoami() {
	fmt.Println("Logged in as:")
	fmt.Println("")
	pub58, username := LoggedInAs()
	fmt.Println(pub58)
	fmt.Println(username)
	fmt.Println("")
}

func ReadLoggedInWords() string {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/secret.txt"
	b, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("    --- not logged in yet, run clout login")
		return ""
	}
	mnemonic := strings.TrimSpace(string(b))
	return mnemonic
}
func SeedBytes() []byte {
	mnemonic := ReadLoggedInWords()
	//entropy, _ := bip39.NewEntropy(128)
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	//b, _ := bip39.MnemonicToByteArray(mnemonic)
	seedBytes, _ := bip39.NewSeedWithErrorChecking(mnemonic, "")
	//fmt.Printf("\n\nPRIVATE\n%x\n\n", seedBytes)
	return seedBytes
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

	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/secret.txt"
	ioutil.WriteFile(path, []byte(text), 0700)
	fmt.Println("Secret stored at:", path)
	fmt.Println("")
	Whoami()
}
func Logout() {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/secret.txt"
	os.Remove(path)
	fmt.Println("Secret removed.")
	fmt.Println("")
}

func LoggedInPub58() string {
	seedBytes := SeedBytes()
	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	return pub58
}
func LoggedInAs() (string, string) {

	seedBytes := SeedBytes()
	//fmt.Printf("%x\n", seedBytes)

	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	return pub58, Pub58ToUsername(pub58)
}
