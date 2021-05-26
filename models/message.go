package models

type MessageList struct {
	NumberOfUnreadThreads       int64
	OrderedContactsWithMessages []MessageThing
	PublicKeyToProfileEntry     map[string]ProfileEntryResponse
	UnreadStateByContact        map[string]bool
}

type MessageThing struct {
	PublicKeyBase58Check string
	Messages             []Message
}

type Message struct {
	SenderPublicKeyBase58Check    string
	RecipientPublicKeyBase58Check string
	EncryptedText                 string
	TstampNanos                   int64
	IsSender                      bool
}
