package main

import (
	"bufio"
	"clout/display"
	"clout/files"
	"clout/keys"
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/justincampbell/timeago"
	"github.com/tyler-smith/go-bip39"
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

func ListPosts(follow bool) {
	js := GetPostsStateless(follow)
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

func ListFollowing(args []string) {
	js := ""
	if len(args) == 0 {
		js = GetFollowsStateless("")
	} else if len(args) == 2 {
		fmt.Println("command missing username")
		return
	} else {
		js = GetFollowsStateless(args[2])
	}
	var pktpe models.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	fmt.Println("NumFollowers", pktpe.NumFollowers)
	fmt.Println("")
	for _, v := range pktpe.PublicKeyToProfileEntry {
		tokens := strings.Split(v.Description, "\n")
		fmt.Printf("%s %s\n", display.LeftAligned(v.Username, 30),
			display.LeftAligned(tokens[0], 30))
	}
}

func ListNotifications() {
	//b, _ := ioutil.ReadFile("samples/get_notifications.list")
	js := GetNotifications()
	var list models.NotificationList
	json.Unmarshal([]byte(js), &list)
	for i, n := range list.Notifications {
		fmt.Printf("%02d %s %s\n", i, display.LeftAligned(n.Metadata.TxnType, 30),
			n.Metadata.CreatorCoinTransferTxindexMetadata.CreatorUsername)
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

func GetPostsStateless(follow bool) string {
	jsonString := `{"GetPostsForGlobalWhitelist":%s,"GetPostsForFollowFeed":%s, "OrderBy":"newest", "ReaderPublicKeyBase58Check": "BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v"}`

	withFollow := fmt.Sprintf(jsonString, "true", "false")
	if follow {
		withFollow = fmt.Sprintf(jsonString, "false", "true")
	}
	jsonString = network.DoPost("api/v0/get-posts-stateless",
		[]byte(withFollow))
	return jsonString
}
func GetFollowsStateless(username string) string {
	jsonString := `{"Username":"%s","PublicKeyBase58Check":"BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v","GetEntriesFollowingUsername":%s,"LastPublicKeyBase58Check":"","NumToFetch":50}`

	withDirection := fmt.Sprintf(jsonString, username, "false")
	if username != "" {
		withDirection = fmt.Sprintf(jsonString, username, "true")
	}

	jsonString = network.DoPost("api/v0/get-follows-stateless",
		[]byte(withDirection))
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
func GetNotifications() string {
	jsonString := `{"PublicKeyBase58Check":"BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v","FetchStartIndex":-1,"NumToFetch":50}`
	jsonString = network.DoPost("api/v0/get-notifications", []byte(jsonString))
	return jsonString
}

func Seal() {

	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println(mnemonic)
	b, _ := bip39.MnemonicToByteArray(mnemonic)
	fmt.Printf("%x\n", b)
	seedBytes, _ := bip39.NewSeedWithErrorChecking(mnemonic, "")
	fmt.Printf("%x\n", seedBytes)

	keys.ComputeKeysFromSeed(seedBytes)
	//params := &lib.BitCloutMainnetParams
	//fmt.Println("Network type set:", params.NetworkType.String())

	//pubKey, privKey, _, _ = lib.ComputeKeysFromSeed(seedBytes, 0, &params)
	//fmt.Printf("%v\n", pubKey)

	//const privateKey = this.cryptoService.seedHexToPrivateKey(seedHex);
	//const signature = privateKey.sign(transactionHash);

	//seedBytes, _ := bip39.NewSeedWithErrorChecking(mnemonic, "")
	//pubKey, privKey, _, _ = lib.ComputeKeysFromSeed(seedBytes, 0, params)

	//seed, _ := bip32.NewSeed()
	//rootPrivateKey, _ := bip32.NewMasterKey(seed)
	//rootPublicKey := rootPrivateKey.PublicKey()
	//key, _ := rootPrivateKey.NewChildKey(0)
	//var nonce [24]byte
	//io.ReadFull(rand.Reader, nonce[:])
	//var a [32]byte
	//copy(a[:], key.Key)
	//encrypted := secretbox.Seal(nonce[:], []byte(tx), &nonce, &a)
	// ec.sign(msg, privateKey, {canonical: true}

	//log.Println(rootPublicKey, fmt.Sprintf("%x", encrypted))
}

func Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter mnenomic: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	b, e := bip39.MnemonicToByteArray(text)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("%x\n", b)

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
