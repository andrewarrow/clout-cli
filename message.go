package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"encoding/json"
	"fmt"
	"time"

	"github.com/justincampbell/timeago"
)

/*
TODO
import {createCipheriv, createDecipheriv, randomBytes, createHmac, createHash} from "crypto";

const EC = require("elliptic").ec;
const ec = new EC("secp256k1");

const aesCtrDecrypt = function(counter, key, data) {
  const cipher = createDecipheriv('aes-128-ctr', key, counter);
  return cipher.update(data).toString();
}

function hmacSha256Sign(key, msg) {
  return createHmac('sha256', key).update(msg).digest();
}

export const kdf = function(secret, outputLength) {
  let ctr = 1;
  let written = 0;
  let result = Buffer.from('');
  while (written < outputLength) {
    const ctrs = Buffer.from([ctr >> 24, ctr >> 16, ctr >> 8, ctr]);
    const hashResult = createHash("sha256").update(Buffer.concat([ctrs, secret])).digest();
    result = Buffer.concat([result, hashResult])
    written += 32;
    ctr +=1;
  }
  return result;
}

export const derive = function(privateKeyA, publicKeyB) {
  assert(Buffer.isBuffer(privateKeyA), "Bad input");
  assert(Buffer.isBuffer(publicKeyB), "Bad input");
  assert(privateKeyA.length === 32, "Bad private key");
  assert(publicKeyB.length === 65, "Bad public key");
  assert(publicKeyB[0] === 4, "Bad public key");
  const keyA = ec.keyFromPrivate(privateKeyA);
  const keyB = ec.keyFromPublic(publicKeyB);
  const Px = keyA.derive(keyB.getPublic());  // BN instance
  return new Buffer(Px.toArray());
};


  const metaLength = 1 + 64 + 16 + 32;
  assert(encrypted.length > metaLength, "Invalid Ciphertext. Data is too small")
  assert(encrypted[0] >= 2 && encrypted[0] <= 4, "Not valid ciphertext.")

  // deserialize
  const ephemPublicKey = encrypted.slice(0, 65);
  const cipherTextLength = encrypted.length - metaLength;
  const iv = encrypted.slice(65, 65 + 16);
  const cipherAndIv = encrypted.slice(65, 65 + 16 + cipherTextLength);
  const ciphertext = cipherAndIv.slice(16);
  const msgMac = encrypted.slice(65 + 16 + cipherTextLength);

  // check HMAC
  const px = derive(privateKey, ephemPublicKey);
  const hash = kdf(px,32);
  const encryptionKey = hash.slice(0, 16);
  const macKey = createHash("sha256").update(hash.slice(16)).digest()
  const dataToMac = Buffer.from(cipherAndIv);
  const hmacGood = hmacSha256Sign(macKey, dataToMac);
  assert(hmacGood.equals(msgMac), "Incorrect MAC");

  // decrypt message
  return aesCtrDecrypt(iv, encryptionKey, ciphertext);

from https://github.com/bitclout/identity/blob/680c584e197eb086e63f0ba12e8142882c2849da/src/lib/ecies/index.js#L121


*/

func ListMessages() {
	m := ReadAccounts()
	for username, s := range m {
		fmt.Println(username)
		pub58, _ := keys.ComputeKeysFromSeed(SeedBytes(s))
		ListMessagesForPub(pub58)
	}
}
func ListMessagesForPub(pub58 string) {
	js := network.GetMessagesStateless(pub58)
	var list models.MessageList
	json.Unmarshal([]byte(js), &list)
	for _, oc := range list.OrderedContactsWithMessages {
		username := list.PublicKeyToProfileEntry[oc.PublicKeyBase58Check].Username
		fmt.Println("  ", username)
		for _, m := range oc.Messages {
			ts := time.Unix(m.TstampNanos/1000000000, 0)
			ago := timeago.FromDuration(time.Since(ts))
			fmt.Println("    ", ago)
		}
	}
}
