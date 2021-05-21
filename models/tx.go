package models

type Tx struct {
	//TxIns     []TxInput
	//TxOuts    []TxOutput
	TxnMeta   TxnMeta
	PublicKey string
}
type TxnMeta struct {
	Body                     string
	CreatorBasisPoints       int64
	StakeMultipleBasisPoints int64
	TimestampNanos           int64
}
type SubmitTx struct {
	TransactionHex string
}

/*

https://stackoverflow.com/questions/51279520/bip32-keys-encryption-in-golang-nacl-secretbox

 encrypted := secretbox.Seal(nonce[:], []byte("hello world"), &nonce, &a)

  signTransaction(seedHex: string, transactionHex: string): string {
    const privateKey = this.cryptoService.seedHexToPrivateKey(seedHex);

    const transactionBytes = new Buffer(transactionHex, 'hex');
    const transactionHash = new Buffer(sha256.x2(transactionBytes), 'hex');
    const signature = privateKey.sign(transactionHash);
    const signatureBytes = new Buffer(signature.toDER());
    const signatureLength = uvarint64ToBuf(signatureBytes.length);

    const signedTransactionBytes = Buffer.concat([
      transactionBytes.slice(0, -1),
      signatureLength,
      signatureBytes,
    ]);

    return signedTransactionBytes.toString('hex');
  }

*/
