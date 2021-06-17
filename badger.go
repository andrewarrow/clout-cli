package main

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

func HandleBadger() {
	db, _ := badger.Open(badger.DefaultOptions(argMap["dir"]))
	defer db.Close()

	for prefixByte := byte(0); prefixByte < byte(40); prefixByte++ {
		fmt.Println(prefixByte)
		EnumerateKeysForPrefix(db, []byte{prefixByte})
	}
}

func EnumerateKeysForPrefix(db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		i := 0
		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()
			val, _ := nodeIterator.Item().ValueCopy(nil)
			fmt.Println(len(key), len(val))
			i++
			if i > 9 {
				break
			}
		}
		return nil
	})

}
