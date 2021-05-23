package main

import (
	"clout/network"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

func SubmitTx(hex string, priv *btcec.PrivateKey) string {
	jsonString := `{"TransactionHex": "%s"}`
	sig, _ := priv.Sign([]byte(hex))
	signedHex := fmt.Sprintf("%x", sig.Serialize())
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
