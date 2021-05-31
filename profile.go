package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func HandleUpdateProfile(argMap map[string]string) {
	if len(argMap) == 0 {
		fmt.Println("")
		fmt.Println("--desc='any desc you want'")
		fmt.Println("--image=path_to_file.png")
		fmt.Println("--percent=18")
		fmt.Println("--username=cloutcli")
		fmt.Println("")
		return
	}

	image := argMap["image"]
	percent := argMap["percent"]
	username := argMap["username"]
	desc := argMap["desc"]
	imageData := ""
	percentData := "300"
	usernameData := ""
	descData := ""
	if image != "" {
		b, e := ioutil.ReadFile(image)
		if e != nil {
			fmt.Println(e)
			return
		}
		str := base64.StdEncoding.EncodeToString(b)
		imageData = "data:image/png;base64," + str
	}
	if percent != "" {
		p, _ := strconv.Atoi(percent)
		percentData = fmt.Sprintf("%d", p*100)
	}
	if username != "" {
		usernameData = username
	}
	if desc != "" {
		descData = desc
	}

	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))
	target := pub58
	jsonString := network.UpdateProfile(pub58, target, descData, usernameData, percentData, imageData)
	var tx models.TxReady
	json.Unmarshal([]byte(jsonString), &tx)
	jsonString = network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}
