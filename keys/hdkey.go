package keys

import (
	"crypto/sha256"

	"github.com/btcsuite/btcutil/hdkeychain"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/base58"
)

func ComputeKeysFromSeed(seedBytes []byte) string {
	netParams := &chaincfg.MainNetParams
	masterKey, _ := hdkeychain.NewMaster(seedBytes, netParams)
	index := uint32(0)

	purpose, _ := masterKey.Child(hdkeychain.HardenedKeyStart + 44)
	coinTypeKey, _ := purpose.Child(hdkeychain.HardenedKeyStart + 0)
	accountKey, _ := coinTypeKey.Child(hdkeychain.HardenedKeyStart + 0)
	changeKey, _ := accountKey.Child(0)
	addressKey, _ := changeKey.Child(index)
	pubKey, _ := addressKey.ECPubKey()
	//privKey, _ := addressKey.ECPrivKey()
	//addressObj, _ := addressKey.Address(netParams)
	//btcDepositAddress := addressObj.EncodeAddress()

	prefix := [3]byte{0xcd, 0x14, 0x0}
	input := pubKey.SerializeCompressed()

	b := []byte{}
	b = append(b, prefix[:]...)
	b = append(b, input[:]...)
	cksum := _checksum(b)
	b = append(b, cksum[:]...)
	return base58.Encode(b)
}

func _checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}
