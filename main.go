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
	fmt.Println("  clout accounts               # list your various accounts")
	fmt.Println("  clout diamond [postHash]     # send 1 diamond")
	fmt.Println("  clout follow [username]      # toggle follow")
	fmt.Println("  clout followers              # who follows you")
	fmt.Println("  clout following              # who you follow")
	fmt.Println("  clout help                   # this menu")
	fmt.Println("  clout like [postHash]        # like/unlike a post")
	fmt.Println("  clout ls                     # list global posts")
	fmt.Println("  clout ls --follow            # filter by follow")
	fmt.Println("  clout ls --post=id           # show single post")
	fmt.Println("  clout login                  # enter secret phrase")
	fmt.Println("  clout logout                 # delete secret from drive")
	fmt.Println("  clout notifications          # list notifications")
	fmt.Println("  clout post --reply=id        # post or reply")
	fmt.Println("  clout reclout [username]     # reclout last thing username said")
	fmt.Println("  clout update [desc]          # update profile description")
	fmt.Println("  clout whoami                 # base58 pubkey logged in")
	fmt.Println("  clout [username]             # username's profile & posts")
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

	if command == "account" || command == "accounts" {
		HandleAccounts()
	} else if command == "follow" {
		HandleFollow()
	} else if command == "followers" {
		ListFollowers()
	} else if command == "following" {
		ListFollowing()
	} else if command == "help" {
		PrintHelp()
	} else if command == "login" {
		Login()
	} else if command == "logout" {
		Logout()
	} else if command == "ls" {
		HandlePosts()
	} else if command == "post" {
		Post(argMap["reply"])
	} else if command == "notifications" || command == "notification" {
		ListNotifications()
	} else if command == "reclout" {
		HandleReclout()
	} else if command == "update" {
		HandleUpdateProfile()
	} else if command == "v8" {
		//RunV8()
	} else if command == "whoami" {
		Whoami()
	} else {
		PostsForPublicKey(command)
	}

}
