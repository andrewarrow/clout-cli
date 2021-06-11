package network

import (
	"bytes"
	"clout/keys"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/textproto"
	"strings"

	"github.com/btcsuite/btcd/btcec"
)

func SubmitTx(hexString string, priv *btcec.PrivateKey) string {
	jsonString := `{"TransactionHex": "%s"}`
	transactionBytes, _ := hex.DecodeString(hexString)
	//fmt.Println("transactionBytes", transactionBytes)
	first := sha256.Sum256(transactionBytes)
	transactionHash := sha256.Sum256(first[:])
	//fmt.Println("transactionHash", transactionHash[:])

	sig, _ := priv.Sign(transactionHash[:])
	signatureBytes := keys.SerializeToDer(sig)

	//fmt.Println("signatureBytes", signatureBytes)

	signatureLength := make([]byte, 8)
	binary.LittleEndian.PutUint64(signatureLength, uint64(len(signatureBytes)))
	//fmt.Println("signatureLength", signatureLength)

	if len(transactionBytes) == 0 {
		return ""
	}

	buff := []byte{}
	buff = append(buff, transactionBytes[0:len(transactionBytes)-1]...)
	buff = append(buff, signatureLength[0])
	buff = append(buff, signatureBytes...)

	//fmt.Println("buff", buff)

	signedHex := fmt.Sprintf("%x", buff)

	//fmt.Println("signedHex", signedHex)
	send := fmt.Sprintf(jsonString, signedHex)
	jsonString = DoPost("api/v0/submit-transaction",
		[]byte(send))
	return jsonString
}
func UpdateProfile(pub58, target, desc, username, percent, image string) string {
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","ProfilePublicKeyBase58Check":"%s","NewUsername":"%s","NewDescription":"%s","NewProfilePic":"%s","NewCreatorBasisPoints":%s,"NewStakeMultipleBasisPoints":12500,"IsHidden":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, pub58, target, username, desc, image, percent)
	jsonString = DoPost("api/v0/update-profile",
		[]byte(send))
	return jsonString
}
func CreateLike(actor, hash string) string {
	jsonString := `{"ReaderPublicKeyBase58Check":"%s","LikedPostHashHex":"%s","IsUnlike":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, actor, hash)
	jsonString = DoPost("api/v0/create-like-stateless",
		[]byte(send))
	return jsonString
}
func CreateFollow(follower, followed string) string {
	jsonString := `{"FollowerPublicKeyBase58Check":"%s","FollowedPublicKeyBase58Check":"%s","IsUnfollow":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, follower, followed)
	jsonString = DoPost("api/v0/create-follow-txn-stateless",
		[]byte(send))
	return jsonString
}
func SendBitclout(sender, receiver string, amount int64) string {
	jsonString := `{"SenderPublicKeyBase58Check":"%s","RecipientPublicKeyOrUsername":"%s","AmountNanos":%d,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, sender, receiver, amount)
	jsonString = DoPost("api/v0/send-bitclout",
		[]byte(send))
	return jsonString
}
func SubmitTransferCoin(sender, creator, receiver string, amount int64) string {
	jsonString := `{"SenderPublicKeyBase58Check":"%s","CreatorPublicKeyBase58Check":"%s","ReceiverUsernameOrPublicKeyBase58Check":"%s","CreatorCoinToTransferNanos":%d,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, sender, creator, receiver, amount)
	jsonString = DoPost("api/v0/transfer-creator-coin",
		[]byte(send))
	return jsonString
}
func SubmitBuyCoin(updater, creator string, sell, expected int64) string {
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","CreatorPublicKeyBase58Check":"%s","OperationType":"buy","BitCloutToSellNanos":%d,"CreatorCoinToSellNanos":0,"BitCloutToAddNanos":0,"MinBitCloutExpectedNanos":0,"MinCreatorCoinExpectedNanos":%d,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, updater, creator, sell, expected)
	jsonString = DoPost("api/v0/buy-or-sell-creator-coin",
		[]byte(send))
	return jsonString
}
func SubmitSellCoin(updater, creator string, sell, expected int64) string {
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","CreatorPublicKeyBase58Check":"%s","OperationType":"sell","BitCloutToSellNanos":0,"CreatorCoinToSellNanos":%d,"BitCloutToAddNanos":0,"MinBitCloutExpectedNanos":%d,"MinCreatorCoinExpectedNanos":0,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, updater, creator, sell, expected)
	jsonString = DoPost("api/v0/buy-or-sell-creator-coin",
		[]byte(send))
	return jsonString
}
func UploadImage(filepath, pub58, jwt string) string {
	imageBytes, _ := ioutil.ReadFile(filepath)
	tokens := strings.Split(filepath, "/")
	filename := tokens[len(tokens)-1]
	tokens = strings.Split(filename, ".")
	ext := tokens[len(tokens)-1]
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	var fw io.Writer
	r := bytes.NewReader(imageBytes)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	imageType := "image/png"
	if ext == "jpg" {
		imageType = "image/jpeg"
	} else if ext == "webp" {
		imageType = "image/webp"
	} else if ext == "gif" {
		imageType = "image/gif"
	}
	h.Set("Content-Type", imageType)

	fw, _ = w.CreatePart(h)
	io.Copy(fw, r)

	rs := strings.NewReader(pub58)
	h = make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s";`, "UserPublicKeyBase58Check"))
	fw, _ = w.CreatePart(h)
	io.Copy(fw, rs)

	rs = strings.NewReader(jwt)
	h = make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s";`, "JWT"))
	fw, _ = w.CreatePart(h)
	io.Copy(fw, rs)

	w.Close()
	jsonString := DoPostMultipart("api/v0/upload-image", w.FormDataContentType(), b.Bytes())
	return jsonString
}
func SubmitDiamond(sender, receiver, post string) string {
	jsonString := `{"SenderPublicKeyBase58Check":"%s","ReceiverPublicKeyBase58Check":"%s","DiamondPostHashHex":"%s","DiamondLevel":1,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, sender, receiver, post)
	jsonString = DoPost("api/v0/send-diamonds",
		[]byte(send))
	return jsonString
}
func SubmitReclout(pub58, RecloutedPostHashHex string) string {
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","PostHashHexToModify":"","ParentStakeID":"","Title":"","BodyObj":{},"RecloutedPostHashHex":"%s","PostExtraData":{},"Sub":"","IsHidden":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, pub58, RecloutedPostHashHex)
	jsonString = DoPost("api/v0/submit-post",
		[]byte(send))
	return jsonString
}
func SubmitPostWithVideo(pub58, body, reply, videoURL string) string {
	if strings.HasPrefix(reply, "https") {
		reply = reply[27:]
	}
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","PostHashHexToModify":"","ParentStakeID":"%s","Title":"","BodyObj":{"Body":"%s","ImageURLs":[]},"RecloutedPostHashHex":"","PostExtraData":{"EmbedVideoURL": "%s"},"Sub":"","IsHidden":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, pub58, reply, body, videoURL)
	jsonString = DoPost("api/v0/submit-post",
		[]byte(send))
	return jsonString
}
func SubmitPostReclout(pub58, body, reclout, imageURL string) string {
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","PostHashHexToModify":"","ParentStakeID":"","Title":"","BodyObj":{"Body":"%s","ImageURLs":[%s]},"RecloutedPostHashHex":"%s","PostExtraData":{},"Sub":"","IsHidden":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, pub58, body, imageURL, reclout)
	jsonString = DoPost("api/v0/submit-post",
		[]byte(send))
	return jsonString
}
func SubmitPost(pub58, body, reply, imageURL string) string {
	if strings.HasPrefix(reply, "https") {
		reply = reply[27:]
	}
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","PostHashHexToModify":"","ParentStakeID":"%s","Title":"","BodyObj":{"Body":"%s","ImageURLs":[%s]},"RecloutedPostHashHex":"","PostExtraData":{},"Sub":"","IsHidden":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, pub58, reply, body, imageURL)
	jsonString = DoPost("api/v0/submit-post",
		[]byte(send))
	return jsonString
}

func GetManyUsersStateless(keys []string) string {
	keyBuff := []string{}
	for _, k := range keys {
		keyBuff = append(keyBuff, fmt.Sprintf("\"%s\"", k))
	}
	jsonString := `{"PublicKeysBase58Check":[%s],"SkipHodlings":false}`
	send := fmt.Sprintf(jsonString, strings.Join(keyBuff, ","))
	jsonString = DoPost("api/v0/get-users-stateless",
		[]byte(send))
	return jsonString
}
func GetUsersStateless(key string) string {
	jsonString := `{"PublicKeysBase58Check":["%s"],"SkipHodlings":false}`
	send := fmt.Sprintf(jsonString, key)
	jsonString = DoPost("api/v0/get-users-stateless",
		[]byte(send))
	return jsonString
}

func GetPostsStatelessWithOptions(last, pub58 string) string {
	jsonString := `{"PostHashHex":"%s","ReaderPublicKeyBase58Check":"%s","OrderBy":"","StartTstampSecs":null,"PostContent":"","NumToFetch":50,"FetchSubcomments":false,"GetPostsForFollowFeed":false,"GetPostsForGlobalWhitelist":true,"GetPostsByClout":false,"PostsByCloutMinutesLookback":0,"AddGlobalFeedBool":false}`

	sendString := fmt.Sprintf(jsonString, last, pub58)
	jsonString = DoPost("api/v0/get-posts-stateless",
		[]byte(sendString))
	return jsonString
}

func GetPostsStateless(pub58 string, follow bool) string {
	jsonString := `{"GetPostsForGlobalWhitelist":%s,"GetPostsForFollowFeed":%s, "OrderBy":"newest", "ReaderPublicKeyBase58Check": "%s"}`

	withFollow := fmt.Sprintf(jsonString, "true", "false", pub58)
	if follow {
		withFollow = fmt.Sprintf(jsonString, "false", "true", pub58)
	}
	jsonString = DoPost("api/v0/get-posts-stateless",
		[]byte(withFollow))
	return jsonString
}
func GetFollowsStateless(pub58, username, last string) string {
	jsonString := `{"Username":"%s","PublicKeyBase58Check":"%s","GetEntriesFollowingUsername":%s,"LastPublicKeyBase58Check":"%s","NumToFetch":50}`

	withDirection := fmt.Sprintf(jsonString, username, pub58, "false", last)
	if username != "" {
		withDirection = fmt.Sprintf(jsonString, username, pub58, "true", last)
	}

	jsonString = DoPost("api/v0/get-follows-stateless",
		[]byte(withDirection))
	return jsonString
}
func GetPostsForPublicKey(key string) string {
	jsonString := `{"PublicKeyBase58Check":"","Username":"%s","ReaderPublicKeyBase58Check":"BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v","LastPostHashHex":"","NumToFetch":10}`
	jsonString = DoPost("api/v0/get-posts-for-public-key",
		[]byte(fmt.Sprintf(jsonString, key)))
	return jsonString
}
func GetSinglePostWithOffset(offset int64, pub58, key string) string {
	jsonString := `{"PostHashHex":"%s","ReaderPublicKeyBase58Check":"%s","FetchParents":true,"CommentOffset":%d,"CommentLimit":20,"AddGlobalFeedBool":false}`
	sendString := fmt.Sprintf(jsonString, key, pub58, offset)
	jsonString = DoPost("api/v0/get-single-post",
		[]byte(sendString))
	return jsonString
}
func GetSinglePost(pub58, key string) string {
	jsonString := `{"PostHashHex":"%s","ReaderPublicKeyBase58Check":"%s","FetchParents":true,"CommentOffset":0,"CommentLimit":20,"AddGlobalFeedBool":false}`
	sendString := fmt.Sprintf(jsonString, key, pub58)
	jsonString = DoPost("api/v0/get-single-post",
		[]byte(sendString))
	return jsonString
}
func GetSingleProfile(key string) string {
	jsonString := `{"PublicKeyBase58Check":"","Username":"%s"}`
	jsonString = DoPost("api/v0/get-single-profile",
		[]byte(fmt.Sprintf(jsonString, key)))
	return jsonString
}
func GetExchangeRate() string {
	jsonString := DoGet("api/v0/get-exchange-rate")
	return jsonString
}
func GetNotificationsWithOffset(offset int, pub58 string) string {
	jsonString := `{"PublicKeyBase58Check":"%s","FetchStartIndex":%d,"NumToFetch":50}`
	sendString := fmt.Sprintf(jsonString, pub58, offset)
	jsonString = DoPost("api/v0/get-notifications", []byte(sendString))
	return jsonString
}
func GetNotifications(pub58 string) string {
	jsonString := `{"PublicKeyBase58Check":"%s","FetchStartIndex":-1,"NumToFetch":50}`
	sendString := fmt.Sprintf(jsonString, pub58)
	jsonString = DoPost("api/v0/get-notifications", []byte(sendString))
	return jsonString
}
func GetMessagesStateless(pub58 string) string {
	jsonString := `{"PublicKeyBase58Check":"%s","FetchAfterPublicKeyBase58Check":"","NumToFetch":25,"HoldersOnly":false,"HoldingsOnly":false,"FollowersOnly":false,"FollowingOnly":false,"SortAlgorithm":"time"}`
	sendString := fmt.Sprintf(jsonString, pub58)
	jsonString = DoPost("api/v0/get-messages-stateless",
		[]byte(sendString))
	return jsonString
}
