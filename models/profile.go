package models

type SingleProfile struct {
	Profile ProfileEntryResponse
}
type ProfileEntryResponse struct {
	PublicKeyBase58Check   string
	Username               string
	Description            string
	CoinEntry              CoinEntry
	CoinPriceBitCloutNanos int64
}
