package main

import (
	"fmt"
	"os"
)

func HandleBackup() {
	words := os.Getenv("CLOUT_PHRASE")
	if len(words) < 36 {
		fmt.Println("")
		fmt.Println("Backup allows you to have just one list of words to unlock")
		fmt.Println("many other lists of words for N number of accounts.")
		fmt.Println("")
		fmt.Println("export CLOUT_PHRASE='these are some nice words and stuff.'")
		fmt.Println("")
		fmt.Println("Set an envionment variable called CLOUT_PHRASE with your words.")
		fmt.Println("The string must be >= 36.")
		fmt.Println("")
		return
	}
}
