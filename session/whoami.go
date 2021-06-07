package session

import (
	"clout/display"
	"fmt"
)

func Whoami(argMap map[string]string) string {
	privHex := ""

	if argMap["private"] != "" {
		mnemonic := ReadLoggedInWords()
		if mnemonic == "" {
			return ""
		}
		seedBytes := SeedBytes(mnemonic)
		privHex = fmt.Sprintf("%x", seedBytes)
	}
	fmt.Println("Logged in as:")
	fmt.Println("")
	pub58, username, balance, btc := LoggedInAs()
	fmt.Println(display.LeftAligned("clout pub58", 20), pub58)
	if privHex != "" {
		fmt.Println(display.LeftAligned("clout priv", 20), privHex)
	}
	fmt.Println(display.LeftAligned("btc address", 20), btc)
	fmt.Println(display.LeftAligned("clout username", 20), username)
	fmt.Println(display.LeftAligned("clout balance", 20), balance)
	fmt.Println("")
	return username
}
