package models

type NotificationList struct {
	Notifications       []Notification
	PostsByHash         map[string]Post
	ProfilesByPublicKey map[string]ProfileEntryResponse
}
type Notification struct {
	Metadata Metadata
}

type Metadata struct {
	TxnType                            string
	TransactorPublicKeyBase58Check     string
	CreatorCoinTransferTxindexMetadata CreatorCoinTransferTxindexMetadata
	SubmitPostTxindexMetadata          SubmitPostTxindexMetadata
	CreatorCoinTxindexMetadata         CreatorCoinTxindexMetadata
	LikeTxindexMetadata                LikeTxindexMetadata
}

type LikeTxindexMetadata struct {
	PostHashHex string
	IsUnlike    bool
}

type CreatorCoinTxindexMetadata struct {
	OperationType          string
	BitCloutToSellNanos    int64
	CreatorCoinToSellNanos int64
	BitCloutToAddNanos     int64
}

type SubmitPostTxindexMetadata struct {
	PostHashBeingModifiedHex string
	ParentPostHashHex        string
}

type CreatorCoinTransferTxindexMetadata struct {
	CreatorUsername            string
	CreatorCoinToTransferNanos int64
	DiamondLevel               int64
	PostHashHex                string
}
