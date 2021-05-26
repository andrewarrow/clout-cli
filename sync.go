package main

import "fmt"

func HandleSync() {
	fmt.Println("")
	fmt.Println("-=-=-= SYNC =-=-=-")
	fmt.Println("")
	fmt.Println("Run this in background to query nodes for blockchain")
	fmt.Println("data about the recent past, further and further back in time.")
	fmt.Println("")
	fmt.Println("You can never reach genesis, it's getting more")
	fmt.Println("difficult as time goes on; the blockchain gets larger and larger.")
	fmt.Println("")
	fmt.Println("1. Decide how much hard drive space you want to")
	fmt.Println("   allocate to this.")
	fmt.Println("")
	fmt.Println("2. Start `clout sync --limit=2000000000` with the number")
	fmt.Println("   of bytes limit you pick.")
	fmt.Println("")
	fmt.Println("Then as your drives fills with data, we index it, makes searching")
	fmt.Println("for old content better and better.")
	fmt.Println("")
	fmt.Println("Syncing now...")
}
