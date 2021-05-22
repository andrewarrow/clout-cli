package keys

import (
	"fmt"

	"github.com/btcsuite/btcutil/hdkeychain"

	"github.com/btcsuite/btcd/chaincfg"
)

func ComputeKeysFromSeed(seedBytes []byte) {
	netParams := &chaincfg.MainNetParams
	masterKey, err := hdkeychain.NewMaster(seedBytes, netParams)
	fmt.Println("hi", masterKey, err)
}
