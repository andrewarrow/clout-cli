package main

import (
	"clout/display"
	"clout/models"
	"clout/network"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/justincampbell/timeago"
	"github.com/tyler-smith/go-bip32"
	"golang.org/x/crypto/nacl/secretbox"
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
	js := GetPostsStateless()
	//b, _ := ioutil.ReadFile("samples/get_posts_stateless.list")
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

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

func GetPostsStateless() string {
	jsonString := `{"ReaderPublicKeyBase58Check": "BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v"}`
	jsonString = network.DoPost("api/v0/get-posts-stateless", []byte(jsonString))
	return jsonString
}

func Seal() {
	seed, _ := bip32.NewSeed()
	rootPrivateKey, _ := bip32.NewMasterKey(seed)
	rootPublicKey := rootPrivateKey.PublicKey()
	key, _ := rootPrivateKey.NewChildKey(0)
	var nonce [24]byte
	io.ReadFull(rand.Reader, nonce[:])
	var a [32]byte
	copy(a[:], key.Key)
	tx := "0161d2cb8074354650c8c34e607737d0ef140bfe871f8f4274c430b7818ce8e1e1000102a6240fb64100b38a3adb749e9ecda8ef0bdcc716a14157b816bf536b9a6e2095cc9805059501000083017b22426f6479223a225468697320736f6e67206279205061756c20576173746572626572672077617320696e2074686520313939322066696c6d205c2253696e676c65735c222068747470733a2f2f7777772e796f75747562652e636f6d2f77617463683f763d4d56684245745453456345222c22496d61676555524c73223a5b5d7de807d461a4b1b982b9f0bcc016002102a6240fb64100b38a3adb749e9ecda8ef0bdcc716a14157b816bf536b9a6e20950000"
	encrypted := secretbox.Seal(nonce[:], []byte(tx), &nonce, &a)

	log.Println(rootPublicKey, fmt.Sprintf("%x", encrypted))
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
