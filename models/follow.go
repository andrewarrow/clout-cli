package models

type PublicKeyToProfileEntry struct {
	PublicKeyToProfileEntry map[string]Follow
	NumFollowers            int64
}

type Follow struct {
	Username    string
	Description string
}
