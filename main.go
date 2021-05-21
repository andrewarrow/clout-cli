package main

import (
	"clout/args"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clout help                  # this menu")
	fmt.Println("  clout ls                    # list global posts")
	fmt.Println("  clout [username]            # that username")
	fmt.Println("  clout login                 # enter secret phrase")
	fmt.Println("  clout logout                # delete secret from drive")
	fmt.Println("")
}

var argMap map[string]string

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	argMap = args.ToMap()

	if argMap["username"] != "" {
		PostsForPublicKey("")
		return
	}

	if command == "ls" {
		ListPosts()
	} else if command == "seal" {
		Seal()
	} else if command == "login" {
		Login()
	} else if command == "logout" {
		Logout()
	} else if command == "gus" {
		GetUsersStateless()
	} else if command == "help" {
		PrintHelp()
	} else {
		PostsForPublicKey(command)
	}

}
