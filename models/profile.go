package models

type SingleProfile struct {
	Profile      ProfileEntryResponse
	BalanceNanos int64
}
type ProfileEntryResponse struct {
	PublicKeyBase58Check string
	Username             string
	Description          string
	CoinEntry            CoinEntry
}
