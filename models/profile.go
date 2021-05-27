package models

type SingleProfile struct {
	Profile ProfileEntryResponse
}
type ProfileEntryResponse struct {
	PublicKeyBase58Check string
	Username             string
	Description          string
	CoinEntry            CoinEntry
	UsersYouHODL         []HODLerThing
}

type HODLerThing struct {
	HODLerPublicKeyBase58Check  string
	CreatorPublicKeyBase58Check string
	BalanceNanos                int64
	ProfileEntryResponse        ProfileEntryResponse
}
