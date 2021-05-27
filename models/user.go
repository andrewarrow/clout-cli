package models

type UsersStateless struct {
	UserList []User
}

type User struct {
	PublicKeyBase58Check string
	ProfileEntryResponse ProfileEntryResponse
	BalanceNanos         int64
	UsersYouHODL         []HODLerThing
	UsersWhoHODLYou      []HODLerThing
}
type HODLerThing struct {
	HODLerPublicKeyBase58Check  string
	CreatorPublicKeyBase58Check string
	BalanceNanos                int64
	ProfileEntryResponse        ProfileEntryResponse
}
