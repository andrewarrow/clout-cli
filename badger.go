package main

import (
	"bytes"
	"clout/draw"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	badger "github.com/dgraph-io/badger/v3"
)

func HandleBadger() {
	db, _ := badger.Open(badger.DefaultOptions(argMap["dir"]))
	defer db.Close()

	//PrefixPostHashToPostEntry := byte(17)
	PrefixPKIDToProfileEntry := byte(23)
	EnumerateKeysForPrefix(db, []byte{PrefixPKIDToProfileEntry})
}

func EnumerateKeysForPrefix(db *badger.DB, dbPrefix []byte) {

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
			if profile.CoinEntry.NumberOfHolders > 1 &&
				len(profile.Description) > 0 &&
				len(profile.ProfilePic) > 0 {
				fmt.Println(string(profile.Username))

				//data:image/jpeg;base64,
				//data:image/png;base64,
				tokens := strings.Split(string(profile.ProfilePic), ",")
				tokens = strings.Split(tokens[0], ";")
				flavor := tokens[0][11:]
				if flavor == "jpeg" {
					flavor = "jpg"
				}

				decodedBytes, _ := base64.StdEncoding.DecodeString(tokens[1])

				draw.SavePicWithPath(flavor, argMap["pic"], string(profile.Username), decodedBytes)
			}
		}
		return nil
	})

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
