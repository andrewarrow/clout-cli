package models

type PostsStateless struct {
	PostsFound []Post
}

type Post struct {
	PostHashHex                string
	PosterPublicKeyBase58Check string
	ParentStakeID              string
	Body                       string
	TimestampNanos             int64
	ProfileEntryResponse       ProfileEntryResponse
}

type ProfileEntryResponse struct {
	PublicKeyBase58Check string
	Username             string
	Description          string
}
