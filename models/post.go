package models

type PostStateless struct {
	PostFound Post
}
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
	Comments                   []Post
	RecloutedPostEntryResponse RecloutedPostEntryResponse
}

type RecloutedPostEntryResponse struct {
	PostHashHex                string
	PosterPublicKeyBase58Check string
	Body                       string
	ProfileEntryResponse       ProfileEntryResponse
}
