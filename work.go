package main

import (
	"clout/display"
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/justincampbell/timeago"
)

func PostsForPublicKey() {
	b, _ := ioutil.ReadFile("samples/get_posts_for_public_key.list")
	var ppk models.PostsPublicKey
	json.Unmarshal(b, &ppk)
	for _, p := range ppk.Posts {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		if p.Body != "" {
			fmt.Println(display.LeftAligned(p.Body, 60), ago)
		} else {
			fmt.Println(display.LeftAligned(p.RecloutedPostEntryResponse.Body, 60), ago)
		}
	}
}

func ListPosts() {
	b, _ := ioutil.ReadFile("samples/get_posts_stateless.list")
	var ps models.PostsStateless
	json.Unmarshal(b, &ps)

	for _, p := range ps.PostsFound {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		fmt.Println(display.LeftAligned(p.ProfileEntryResponse.Username, 30),
			display.LeftAligned(p.ProfileEntryResponse.CoinEntry.NumberOfHolders, 20),
			ago)
	}
}

func GetUsersStateless() {
	jsonString := `{"PublicKeyBase58Check": "BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v"}`
	jsonString = network.DoPost("api/v0/get-users-stateless", []byte(jsonString))
	fmt.Println(jsonString)

	jsonString = `{"PublicKeyBase58Check": "", "Username":"katramdeen"}`
	jsonString = network.DoPost("api/v0/get-single-profile", []byte(jsonString))
	fmt.Println(jsonString)
}

/*
		jsonString := network.DoGet("api/v0/get-exchange-rate")
		var rate models.Rate
		json.Unmarshal([]byte(jsonString), &rate)
		fmt.Println(rate)
		jsonString = network.DoGet("api/v0/health-check")
		fmt.Println(jsonString)
		jsonString = `{"PublicKeyBase58Check": "hi"}`
		jsonString = network.DoPost("api/v0/get-app-state", []byte(jsonString))
		fmt.Println(jsonString)
	seed, _ := bip32.NewSeed()
	s, _ := bip32.NewMasterKey(seed)
	fmt.Println(s, s.PublicKey())

	    const network = this.globalVars.network;
	const keychain = this.cryptoService.mnemonicToKeychain(this.mnemonicCheck, this.extraTextCheck);
	const seedHex = this.cryptoService.keychainToSeedHex(keychain);
	const btcDepositAddress = this.cryptoService.keychainToBtcAddress(keychain, network);

	this.publicKeyAdded = this.accountService.addUser({
	  seedHex,
	  mnemonic: this.mnemonicCheck,
	  extraText: this.extraTextCheck,
	  btcDepositAddress,
	  network,
	});
*/
