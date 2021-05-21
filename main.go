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
	fmt.Println("  clout ls                    # list posts")
	fmt.Println("  clout --username=x          # ")
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
		PostsForPublicKey()
		return
	}

	if command == "ls" {
		ListPosts()
	} else if command == "seal" {
		Seal()
	} else if command == "gus" {
		GetUsersStateless()
	} else if command == "help" {
		PrintHelp()
	}

}
