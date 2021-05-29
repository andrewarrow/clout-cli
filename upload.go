package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcec"
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
	var image models.Image
	json.Unmarshal([]byte(js), &image)
	fmt.Println(image.ImageURL)
}

func UploadImage(path, pub58 string, priv *btcec.PrivateKey) string {
	jwt := keys.MakeJWT(priv)
	js := network.UploadImage(path, pub58, jwt)
	var image models.Image
	json.Unmarshal([]byte(js), &image)
	return image.ImageURL
}
