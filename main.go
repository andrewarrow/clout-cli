package main

import (
	"clout-cli/models"
	"encoding/json"
	"fmt"
	"github.com/justincampbell/timeago"
	"io/ioutil"
	"time"
	//"github.com/tyler-smith/go-bip32"
)

func main() {
	fmt.Println("clout-cli")
	fmt.Println("")

	b, _ := ioutil.ReadFile("samples/get_posts_stateless.list")
	var ps models.PostsStateless
	json.Unmarshal(b, &ps)

	for _, p := range ps.PostsFound {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		fmt.Println(ago, p.ProfileEntryResponse.Username)
	}

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
