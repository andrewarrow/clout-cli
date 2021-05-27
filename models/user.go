package models

type UsersStateless struct {
	UserList []User
}

type User struct {
	PublicKeyBase58Check string
	ProfileEntryResponse ProfileEntryResponse
	BalanceNanos         int64
	UsersYouHODL         string
}
