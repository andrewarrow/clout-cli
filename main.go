package main

import (
	"clout-cli/models"
	"clout-cli/network"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("clout-cli")
	jsonString := network.DoGet("api/v0/get-exchange-rate")
	var rate models.Rate
	json.Unmarshal([]byte(jsonString), &rate)
	fmt.Println(rate)
}
