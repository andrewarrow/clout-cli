package keys

import (
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/dgrijalva/jwt-go"
)

func MakeJWT(priv *btcec.PrivateKey) string {

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 60).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	signed, _ := token.SignedString(priv.ToECDSA())
	return signed
}

/*
 identityServiceParamsForKey(publicKey: string) {
    const { encryptedSeedHex, accessLevel, accessLevelHmac } = this.identityServiceUsers[publicKey];
    return { encryptedSeedHex, accessLevel, accessLevelHmac };
  }


	 const { id, payload: { encryptedSeedHex } } = data;
    const seedHex = this.cryptoService.decryptSeedHex(encryptedSeedHex, this.globalVars.hostname);
    const jwt = this.signingService.signJWT(seedHex);


		 decryptSeedHex(encryptedSeedHex: string, hostname: string): string {
    const encryptionKey = this.seedHexEncryptionKey(hostname);
    const decipher = createDecipher('aes-256-gcm', encryptionKey);
    return decipher.update(Buffer.from(encryptedSeedHex, 'hex')).toString();
  }

	 signJWT(seedHex: string): string {
    const keyEncoder = new KeyEncoder('secp256k1');
    const encodedPrivateKey = keyEncoder.encodePrivate(seedHex, 'raw', 'pem');
  return jsonwebtoken.sign({ }, encodedPrivateKey, { algorithm: 'ES256', expiresIn: 60 });
  }


		https://github.com/bitclout/identity/blob/3b4fadc9c376cc20da48df4a13a8adc918382e04/src/app/identity.service.ts#L114

https://github.com/bitclout/frontend/blob/adbc4e0aa83cf7ebef6a980f5033fb283d982d10/src/app/identity.service.ts#L118*/
