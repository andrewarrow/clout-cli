package models

type PostsStateless struct {
	PostsFound []Post
}

type PostsPublicKey struct {
	Posts []Post
}

type Post struct {
	PostHashHex                string
	PosterPublicKeyBase58Check string
	ParentStakeID              string
	Body                       string
	TimestampNanos             int64
	ProfileEntryResponse       ProfileEntryResponse
	RecloutedPostEntryResponse RecloutedPostEntryResponse
}

type RecloutedPostEntryResponse struct {
	PostHashHex                string
	PosterPublicKeyBase58Check string
	Body                       string
	ProfileEntryResponse       ProfileEntryResponse
}

type ProfileEntryResponse struct {
	PublicKeyBase58Check string
	Username             string
	Description          string
	CoinEntry            CoinEntry
}

type CoinEntry struct {
	CreatorBasisPoints      int64
	BitCloutLockedNanos     int64
	NumberOfHolders         int64
	CoinsInCirculationNanos int64
	CoinWatermarkNanos      int64
}
