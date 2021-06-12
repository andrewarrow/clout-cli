package models

type HodlersWrap struct {
	Hodlers []Hodler
}

type Hodler struct {
	HODLerPublicKeyBase58Check string
	BalanceNanos               int64
	ProfileEntryResponse       ProfileEntryResponse
}
