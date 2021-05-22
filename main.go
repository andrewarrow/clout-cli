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
	fmt.Println("  clout help                   # this menu")
	fmt.Println("  clout ls                     # list global posts")
	fmt.Println("  clout ls --follow            # filter by follow")
	fmt.Println("  clout [username]             # username's profile & posts")
	fmt.Println("  clout login                  # enter secret phrase")
	fmt.Println("  clout logout                 # delete secret from drive")
	fmt.Println("  clout like [postHash]        # like/unlike a post")
	fmt.Println("  clout diamond [postHash]     # send 1 diamond")
	fmt.Println("  clout post --reply=postHash  # post or reply")
	fmt.Println("  clout reclout [postHash]     # reclout specific post")
	fmt.Println("  clout follow [username]      # toggle follow")
	fmt.Println("  clout notifications          # list notifications")
	fmt.Println("  clout following              # list of users")
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
		ListPosts(argMap["follow"] == "true")
	} else if command == "following" {
		ListFollowing()
	} else if command == "gus" {
		GetUsersStateless()
	} else if command == "help" {
		PrintHelp()
	} else if command == "login" {
		Login()
	} else if command == "logout" {
		Logout()
	} else if command == "notifications" || command == "notification" {
		ListNotifications()
	} else if command == "seal" {
		Seal()
	} else {
		PostsForPublicKey(command)
	}

}
