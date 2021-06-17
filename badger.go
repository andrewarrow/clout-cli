package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

func HandleBadger() {
	db, _ := badger.Open(badger.DefaultOptions(argMap["dir"]))
	defer db.Close()

	PrefixPostHashToPostEntry := byte(17)
	EnumerateKeysForPrefix(db, []byte{PrefixPostHashToPostEntry})
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

			postEntryObj := &PostEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(postEntryObj)
			fmt.Println(string(postEntryObj.Body))
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
