package main

import (
	"clout/keys"
	"clout/session"
	"fmt"
)

func HandleUpload() {
	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	seedBytes := session.SeedBytes(mnemonic)

	_, priv, _ := keys.ComputeKeysFromSeedWithAddress(seedBytes)
	jwt := keys.MakeJWT(priv)
	fmt.Println(jwt)
	//network.UploadImage(os.Args[2])
}
