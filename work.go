package main

import (
	"bufio"
	"clout/display"
	"clout/files"
	"clout/models"
	"clout/network"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/justincampbell/timeago"
	"github.com/tyler-smith/go-bip32"
	"golang.org/x/crypto/nacl/secretbox"
)

func PostsForPublicKey(key string) {
	js := GetSingleProfile(key)
	var sp models.SingleProfile
	json.Unmarshal([]byte(js), &sp)
	fmt.Println("---", sp.Profile.CoinEntry.CreatorBasisPoints)
	fmt.Println(sp.Profile.Description)
	fmt.Println("---")

	js = GetPostsForPublicKey(key)
	//b, _ := ioutil.ReadFile("samples/get_posts_for_public_key.list")
	var ppk models.PostsPublicKey
	json.Unmarshal([]byte(js), &ppk)
	for _, p := range ppk.Posts {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		body := ""
		if p.Body != "" {
			body = p.Body
		} else {
			body = p.RecloutedPostEntryResponse.Body
		}
		tokens := strings.Split(body, "\n")
		for i, t := range tokens {
			if strings.TrimSpace(t) == "" {
				continue
			}
			fmt.Println("    ", i, t)
		}
		fmt.Println(ago)
		fmt.Println("")
	}
}

func ListPosts() {
	js := GetPostsStateless()
	//b, _ := ioutil.ReadFile("samples/get_posts_stateless.list")
	var ps models.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	for i, p := range ps.PostsFound {
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		fmt.Println(display.LeftAligned(p.ProfileEntryResponse.Username, 30),
			display.LeftAligned(p.ProfileEntryResponse.CoinEntry.NumberOfHolders, 20),
			ago)
		tokens := strings.Split(p.Body, "\n")
		fmt.Println("        ", display.LeftAligned(tokens[0], 40))
		fmt.Println("")
		if i > 6 {
			break
		}
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
func GetPostsForPublicKey(key string) string {
	jsonString := `{"PublicKeyBase58Check":"","Username":"%s","ReaderPublicKeyBase58Check":"BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v","LastPostHashHex":"","NumToFetch":10}`
	jsonString = network.DoPost("api/v0/get-posts-for-public-key",
		[]byte(fmt.Sprintf(jsonString, key)))
	return jsonString
}
func GetSingleProfile(key string) string {
	jsonString := `{"PublicKeyBase58Check":"","Username":"%s"}`
	jsonString = network.DoPost("api/v0/get-single-profile",
		[]byte(fmt.Sprintf(jsonString, key)))
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

func Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	home := files.UserHomeDir()
	dir := "clout-cli-data"
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/secret.txt"
	ioutil.WriteFile(path, []byte(text), 0700)
	fmt.Println("Secret stored at:", path)
	fmt.Println("")
}
func Logout() {
	home := files.UserHomeDir()
	dir := "clout-cli-data"
	path := home + "/" + dir + "/secret.txt"
	os.Remove(path)
	fmt.Println("Secret removed.")
	fmt.Println("")
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
