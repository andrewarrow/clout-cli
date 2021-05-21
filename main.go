package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	//"github.com/tyler-smith/go-bip32"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clout-cli help         # this menu")
	fmt.Println("  clout-cli ls           # list posts")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "ls" {
		ListPosts()
	} else if command == "" {
	} else if command == "help" {
		PrintHelp()
	}

}
