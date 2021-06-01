package models

type PublicKeyToProfileEntry struct {
	PublicKeyToProfileEntry map[string]ProfileEntryResponse
	NumFollowers            int64
}
