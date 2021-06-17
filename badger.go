package main

import (
	"clout/sync"

	badger "github.com/dgraph-io/badger/v3"
)

func HandleBadger() {
	db, _ := badger.Open(badger.DefaultOptions(argMap["dir"]))
	defer db.Close()

	//PrefixPostHashToPostEntry := byte(17)
	PrefixPKIDToProfileEntry := byte(23)
	sync.EnumerateKeysForPrefix(db, []byte{PrefixPKIDToProfileEntry})
}
