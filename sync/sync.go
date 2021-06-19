package sync

import (
	"bytes"
	"clout/models"
	"clout/network"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/dgraph-io/badger/v3"
	"github.com/justincampbell/timeago"
)

func HandleSync(argMap map[string]string) {
	CreateSchema()
	db, _ := badger.Open(badger.DefaultOptions(argMap["dir"]))
	defer db.Close()
	PrefixPKIDToProfileEntry := byte(23)
	EnumerateKeysForPrefix(db, []byte{PrefixPKIDToProfileEntry})
}

func EnumerateKeysForPrefix(db *badger.DB, dbPrefix []byte) {
	sdb := OpenTheDB()
	defer sdb.Close()
	tx, _ := sdb.Begin()

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			//key := nodeIterator.Item().Key()
			val, _ := nodeIterator.Item().ValueCopy(nil)

			//postEntryObj := &PostEntry{}
			profile := &ProfileEntry{}
			//gob.NewDecoder(bytes.NewReader(val)).Decode(postEntryObj)
			gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
			//fmt.Println(string(postEntryObj.Body))

			pub58 := base58.Encode(profile.PublicKey)

			if InsertUser(tx, string(profile.Username), pub58) {
				LoadUserAndLook(string(profile.Username))
			}
		}
		return nil
	})
	tx.Commit()
}

func LoadUserAndLook(username string) {
	js := network.GetSingleProfile(username)
	var sp models.SingleProfile
	json.Unmarshal([]byte(js), &sp)

	js = network.GetPostsForPublicKey(username)
	var ppk models.PostsPublicKey
	json.Unmarshal([]byte(js), &ppk)
	for _, p := range ppk.Posts {
		fmt.Println(username)
		ts := time.Unix(p.TimestampNanos/1000000000, 0)
		ago := timeago.FromDuration(time.Since(ts))
		tokens := strings.Split(p.Body, "\n")
		fmt.Println(tokens[0])
		fmt.Println(ago)
		fmt.Println(sp.Profile.CoinEntry.CreatorBasisPoints, len(sp.Profile.Description))
		fmt.Println("")
		break
	}
}

type StakeEntryStats struct {
}

const HashSizeBytes = 32

type BlockHash [HashSizeBytes]byte
type StakeEntry struct {
}

type PostEntry struct {
	PostHash                 *BlockHash
	PosterPublicKey          []byte
	ParentStakeID            []byte
	Body                     []byte
	RecloutedPostHash        *BlockHash
	IsQuotedReclout          bool
	CreatorBasisPoints       uint64
	StakeMultipleBasisPoints uint64
	ConfirmationBlockHeight  uint32
	TimestampNanos           uint64
	IsHidden                 bool
	StakeEntry               *StakeEntry
	LikeCount                uint64
	RecloutCount             uint64
	QuoteRecloutCount        uint64
	DiamondCount             uint64
	stakeStats               *StakeEntryStats
	isDeleted                bool
	CommentCount             uint64
	IsPinned                 bool
	PostExtraData            map[string][]byte
}

type CoinEntry struct {
	CreatorBasisPoints      uint64
	BitCloutLockedNanos     uint64
	NumberOfHolders         uint64
	CoinsInCirculationNanos uint64
	CoinWatermarkNanos      uint64
}

type ProfileEntry struct {
	PublicKey   []byte
	Username    []byte
	Description []byte
	ProfilePic  []byte
	IsHidden    bool
	CoinEntry
	isDeleted                bool
	StakeMultipleBasisPoints uint64
	StakeEntry               *StakeEntry
	stakeStats               *StakeEntryStats
}

/*
	if profile.CoinEntry.NumberOfHolders > 1 &&
		len(profile.Description) > 0 &&
		len(profile.ProfilePic) > 0 {
		fmt.Println(string(profile.Username))

		//data:image/jpeg;base64,
		//data:image/png;base64,
		tokens := strings.Split(string(profile.ProfilePic), ",")
		base64data := tokens[1]
		tokens = strings.Split(tokens[0], ";")
		flavor := tokens[0][11:]
		if flavor == "jpeg" {
			flavor = "jpg"
		}

		decodedBytes, _ := base64.StdEncoding.DecodeString(base64data)

		draw.SavePicWithPath(flavor, argMap["pic"], string(profile.Username), decodedBytes)
	}*/
