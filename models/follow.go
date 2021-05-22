package models

type PublicKeyToProfileEntry struct {
	PublicKeyToProfileEntry map[string]Follow
}

type Follow struct {
	Username    string
	Description string
}
