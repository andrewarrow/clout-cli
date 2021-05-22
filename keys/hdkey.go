package keys

import (
	"fmt"

	"github.com/btcsuite/btcutil/hdkeychain"

	"github.com/btcsuite/btcd/chaincfg"
)

func ComputeKeysFromSeed(seedBytes []byte) {
	netParams := &chaincfg.MainNetParams
	masterKey, _ := hdkeychain.NewMaster(seedBytes, netParams)
	index := uint32(0)

	purpose, _ := masterKey.Child(hdkeychain.HardenedKeyStart + 44)
	coinTypeKey, _ := purpose.Child(hdkeychain.HardenedKeyStart + 0)
	accountKey, _ := coinTypeKey.Child(hdkeychain.HardenedKeyStart + 0)
	changeKey, _ := accountKey.Child(0)
	addressKey, _ := changeKey.Child(index)
	pubKey, _ := addressKey.ECPubKey()
	privKey, _ := addressKey.ECPrivKey()
	addressObj, _ := addressKey.Address(netParams)
	btcDepositAddress := addressObj.EncodeAddress()

	fmt.Println("pubKey", pubKey)
	fmt.Println("privKey", privKey)
	fmt.Println("addressObj", addressObj)
	fmt.Println("btcDepositAddress", btcDepositAddress)
}
