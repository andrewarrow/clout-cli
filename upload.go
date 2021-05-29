package main

import (
	"clout/keys"
	"clout/network"
	"clout/session"
	"fmt"
	"os"
)

func HandleUpload() {
	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	seedBytes := session.SeedBytes(mnemonic)

	pub58, priv, _ := keys.ComputeKeysFromSeedWithAddress(seedBytes)
	jwt := keys.MakeJWT(priv)
	js := network.UploadImage(os.Args[2], pub58, jwt)
	fmt.Println(js)
}
