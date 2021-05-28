package session

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

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
	WriteSelected(username)
}
