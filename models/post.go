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
	PostExtraData              PostExtraData
	ImageURLs                  []string
	TimestampNanos             int64
	ProfileEntryResponse       ProfileEntryResponse
	LikeCount                  int64
	Comments                   []Post
	RecloutedPostEntryResponse *Post
	CommentCount               int64
	RecloutCount               int64
}

type PostExtraData struct {
	EmbedVideoURL string
}
