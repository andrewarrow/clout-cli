package keys

import (
	"crypto/sha256"
	"fmt"
	"math"

	"github.com/btcsuite/btcutil/hdkeychain"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/base58"
)

func constructLength(orig []byte, l byte) []byte {
	arr := []byte{}
	arr = append(arr, orig...)
	if l < 0x80 {
		arr = append(arr, l)
		return arr
	}
	log1 := byte(math.Log(float64(l) / math.Ln2))

	octets := 1 + log1 //>> 3;
	arr = append(arr, (octets>>3)|0x80)
	for {
		arr = append(arr, (l>>(octets<<3))&0xff)
		octets--
		if octets == 0 {
			break
		}
	}
	arr = append(arr, l)
	return arr
}

func rmPadding(buf []byte) []byte {
	i := 0
	l := len(buf) - 1
	for buf[i] != 0 && (buf[i+1]&0x80) != 0 && i < l {
		i++
	}
	if i == 0 {
		return buf
	}
	return buf[i:]
}

func SerializeToDer(sig *btcec.Signature) []byte {
	r := sig.R.Bytes()
	s := sig.S.Bytes()

	fmt.Printf("ok starting with %x for r\n", r)
	fmt.Printf("ok starting with %x for s\n", s)

	if (r[0] & 0x80) != 0 {
		r = append([]byte{0}, r...)
	}
	if (s[0] & 0x80) != 0 {
		s = append([]byte{0}, s...)
	}

	r = rmPadding(r)
	s = rmPadding(s)

	fmt.Printf("s before %x\n", s)
	for s[0] != 0 && (s[1]&0x80) != 0 {
		s = s[1:]
	}
	fmt.Printf("s after %x\n", s)
	arr := []byte{0x02}

	arr = constructLength(arr, byte(len(r)))
	fmt.Printf("arr1 after %x\n", arr)
	arr = append(arr, r...)
	arr = append(arr, 0x02)
	fmt.Printf("arr2 after %x\n", arr)
	arr = constructLength(arr, byte(len(s)))
	fmt.Printf("arr3 after %x\n", arr)
	backHalf := append(arr, s...)
	res := []byte{0x30}
	res = constructLength(res, byte(len(backHalf)))
	res = append(res, backHalf...)
	fmt.Printf("res %x\n", res)
	return res
}

func ComputeKeysFromSeed(seedBytes []byte) (string, *btcec.PrivateKey) {
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
	//addressObj, _ := addressKey.Address(netParams)
	//btcDepositAddress := addressObj.EncodeAddress()

	prefix := [3]byte{0xcd, 0x14, 0x0}
	input := pubKey.SerializeCompressed()

	b := []byte{}
	b = append(b, prefix[:]...)
	b = append(b, input[:]...)
	cksum := _checksum(b)
	b = append(b, cksum[:]...)
	return base58.Encode(b), privKey
}

func _checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}
