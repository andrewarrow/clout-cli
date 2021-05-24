package main

import (
	"clout/keys"
	"clout/network"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

func SubmitTx(hexString string, priv *btcec.PrivateKey) string {
	jsonString := `{"TransactionHex": "%s"}`
	transactionBytes, _ := hex.DecodeString(hexString)
	fmt.Println("transactionBytes", transactionBytes)
	first := sha256.Sum256(transactionBytes)
	transactionHash := sha256.Sum256(first[:])
	fmt.Println("transactionHash", transactionHash[:])

	sig, _ := priv.Sign(transactionHash[:])
	signatureBytes := keys.SerializeToDer(sig)

	fmt.Println("signatureBytes", signatureBytes)

	signatureLength := make([]byte, 8)
	binary.LittleEndian.PutUint64(signatureLength, uint64(len(signatureBytes)))
	fmt.Println("signatureLength", signatureLength)

	buff := []byte{}
	buff = append(buff, transactionBytes[0:len(transactionBytes)-1]...)
	buff = append(buff, signatureLength[0])
	buff = append(buff, signatureBytes...)

	fmt.Println("buff", buff)

	signedHex := fmt.Sprintf("%x", buff)

	fmt.Println("signedHex", signedHex)
	send := fmt.Sprintf(jsonString, signedHex)
	jsonString = network.DoPost("api/v0/submit-transaction",
		[]byte(send))
	return jsonString
}
func SubmitPost(pub58, body string) string {
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","PostHashHexToModify":"","ParentStakeID":"","Title":"","BodyObj":{"Body":"%s","ImageURLs":[]},"RecloutedPostHashHex":"","PostExtraData":{},"Sub":"","IsHidden":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, pub58, body)
	jsonString = network.DoPost("api/v0/submit-post",
		[]byte(send))
	return jsonString
}

func GetUsersStateless(key string) string {
	jsonString := `{"PublicKeysBase58Check":["%s"],"SkipHodlings":false}`
	send := fmt.Sprintf(jsonString, key)
	jsonString = network.DoPost("api/v0/get-users-stateless",
		[]byte(send))
	return jsonString
}

func GetPostsStateless(pub58 string, follow bool) string {
	jsonString := `{"GetPostsForGlobalWhitelist":%s,"GetPostsForFollowFeed":%s, "OrderBy":"newest", "ReaderPublicKeyBase58Check": "%s"}`

	withFollow := fmt.Sprintf(jsonString, "true", "false", pub58)
	if follow {
		withFollow = fmt.Sprintf(jsonString, "false", "true", pub58)
	}
	jsonString = network.DoPost("api/v0/get-posts-stateless",
		[]byte(withFollow))
	return jsonString
}
func GetFollowsStateless(pub58, username string) string {
	jsonString := `{"Username":"%s","PublicKeyBase58Check":"%s","GetEntriesFollowingUsername":%s,"LastPublicKeyBase58Check":"","NumToFetch":50}`

	withDirection := fmt.Sprintf(jsonString, username, pub58, "false")
	if username != "" {
		withDirection = fmt.Sprintf(jsonString, username, pub58, "true")
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
func GetNotifications(pub58 string) string {
	jsonString := `{"PublicKeyBase58Check":"%s","FetchStartIndex":-1,"NumToFetch":50}`
	sendString := fmt.Sprintf(jsonString, pub58)
	jsonString = network.DoPost("api/v0/get-notifications", []byte(sendString))
	return jsonString
}
