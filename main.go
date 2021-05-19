package main

import (
	"clout-cli/network"
	"fmt"
)

func main() {
	fmt.Println("clout-cli")
	jsonString := network.DoGet("api/v0/get-exchange-rate")
	fmt.Println(jsonString)
}
