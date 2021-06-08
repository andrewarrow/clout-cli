package main

import (
	"fmt"
	"os"
)

func HandleInspect() {
	if len(os.Args) < 3 {
		fmt.Println("username")
		return
	}

	username := os.Args[2]
	fmt.Println(username)
}
