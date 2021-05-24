package main

import (
	"bufio"
	"clout/display"
	"clout/files"
	"clout/keys"
	"clout/models"
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
	pub58 := LoggedInPub58()
	js := GetPostsStateless(pub58, follow)
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

func Pub58ToUsername(key string) string {
	js := GetUsersStateless(key)
	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	return us.UserList[0].ProfileEntryResponse.Username
}

var dir = "clout-cli-data"

func Whoami() {
	fmt.Println("Logged in as:")
	fmt.Println("")
	pub58, username := LoggedInAs()
	fmt.Println(pub58)
	fmt.Println(username)
	fmt.Println("")
}

func ReadLoggedInWords() string {
	home := files.UserHomeDir()
	path := home + "/" + dir + "/secret.txt"
	b, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("    --- not logged in yet, run clout login")
		return ""
	}
	mnemonic := strings.TrimSpace(string(b))
	return mnemonic
}
func SeedBytes() []byte {
	mnemonic := ReadLoggedInWords()
	//entropy, _ := bip39.NewEntropy(128)
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	//b, _ := bip39.MnemonicToByteArray(mnemonic)
	seedBytes, _ := bip39.NewSeedWithErrorChecking(mnemonic, "")
	//fmt.Printf("\n\nPRIVATE\n%x\n\n", seedBytes)
	return seedBytes
}

func LoggedInPub58() string {
	seedBytes := SeedBytes()
	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	return pub58
}
func LoggedInAs() (string, string) {

	seedBytes := SeedBytes()
	//fmt.Printf("%x\n", seedBytes)

	pub58, _ := keys.ComputeKeysFromSeed(seedBytes)
	return pub58, Pub58ToUsername(pub58)

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
func Post() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Say: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	pub58, priv := keys.ComputeKeysFromSeed(SeedBytes())
	bigString := SubmitPost(pub58, text)

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	tsUnix := tx.TstampNanos / 1000000000
	ts := time.Unix(tsUnix, 0)

	fmt.Println(ts)
	fmt.Println(tx.TransactionHex)

	jsonString := SubmitTx(tx.TransactionHex, priv)
	fmt.Println(jsonString)
}

func Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter mnenomic: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	_, e := bip39.MnemonicToByteArray(text)
	if e != nil {
		fmt.Println(e)
		return
	}
	//fmt.Printf("%x\n", b)

	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/secret.txt"
	ioutil.WriteFile(path, []byte(text), 0700)
	fmt.Println("Secret stored at:", path)
	fmt.Println("")
	Whoami()
}
func Logout() {
	home := files.UserHomeDir()
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
