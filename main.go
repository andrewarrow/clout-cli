package main

import (
	"clout/args"
	"clout/draw"
	"clout/session"
	"clout/sync"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clout accounts               # list your various accounts")
	fmt.Println("  clout backup                 # encrypt and copy secrets")
	fmt.Println("  clout balances               # list available $bitclout")
	fmt.Println("  clout boards                 # list boards you are on")
	fmt.Println("  clout buy                    # buy creator coin")
	fmt.Println("  clout diamond [username]     # award 1 diamond to last post")
	fmt.Println("  clout follow [username]      # toggle follow")
	fmt.Println("  clout followers              # who follows you")
	fmt.Println("  clout following              # who you follow")
	fmt.Println("  clout help                   # this menu")
	fmt.Println("  clout like --hash=x          # like a post")
	fmt.Println("  clout ls                     # list global posts")
	fmt.Println("  clout ls --follow            # filter by follow")
	fmt.Println("  clout ls --hash=x            # show single post")
	fmt.Println("  clout login                  # enter secret phrase")
	fmt.Println("  clout logout                 # delete secret from drive")
	fmt.Println("  clout messages               # list messages")
	fmt.Println("  clout notifications          # list notifications")
	fmt.Println("  clout post --reply=x         # post or reply")
	fmt.Println("  clout reclout [username]     # reclout last post")
	fmt.Println("  clout sync                   # fill local hard drive with data")
	fmt.Println("  clout update                 # update profile description")
	fmt.Println("  clout wallet                 # list what you own")
	fmt.Println("  clout whoami                 # base58 pubkey logged in")
	fmt.Println("  clout [username]             # username's profile & posts")
	fmt.Println("")
	username := session.SelectedAccount()
	if username != "" {
		fmt.Println("SELECTED ACCOUNT:", username)
		fmt.Println("")
	}
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
		session.HandleAccounts(argMap)
	} else if command == "backup" || command == "backups" {
		HandleBackup(argMap)
	} else if command == "balance" || command == "balances" {
		HandleBalances(argMap)
	} else if command == "board" || command == "boards" {
		HandleBoards()
	} else if command == "bulk" {
		HandleBulk()
	} else if command == "buy" {
		HandleBuy()
	} else if command == "clown" {
		HandleClown()
	} else if command == "diamond" {
		HandleDiamond()
	} else if command == "draw" {
		draw.DrawDiamondImage()
	} else if command == "enrich" {
		FindBuysSellsAndTransfers()
	} else if command == "follow" {
		HandleFollow()
	} else if command == "followers" {
		ListFollowers()
	} else if command == "following" {
		HandleFollowing()
	} else if command == "global" {
		HandleGlobal()
	} else if command == "help" {
		PrintHelp()
	} else if command == "inspect" {
		HandleInspect()
	} else if command == "like" {
		HandleLike(argMap)
	} else if command == "login" {
		session.Login()
	} else if command == "logout" {
		session.Logout()
	} else if command == "long" {
		HandleLongThread()
	} else if command == "ls" {
		HandlePosts()
	} else if command == "post" {
		Post(argMap)
	} else if command == "machine" {
		HandleMachine()
	} else if command == "messages" || command == "message" {
		ListMessages()
	} else if command == "n" || command == "notifications" || command == "notification" {
		HandleNotifications(argMap)
	} else if command == "random" {
		RandomEmo(20)
	} else if command == "reclout" {
		HandleReclout()
	} else if command == "sell" {
		HandleSell()
	} else if command == "send" {
		HandleSend()
	} else if command == "sync" {
		sync.HandleSync(argMap)
	} else if command == "tags" {
		HandleTags()
	} else if command == "update" {
		HandleUpdateProfile(argMap)
	} else if command == "upload" {
		HandleUpload()
	} else if command == "wallet" {
		HandleWallet(argMap)
	} else if command == "words" {
		HandleWords(argMap)
	} else if command == "whoami" {
		session.Whoami(argMap)
	} else if command == "youtube" {
		HandleYoutube()
	} else {
		GuiViewUser(command)
	}

}
