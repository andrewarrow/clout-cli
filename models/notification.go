package models

type NotificationList struct {
	Notifications []Notification
}
type Notification struct {
	Metadata Metadata
}

type Metadata struct {
	TxnType                            string
	TransactorPublicKeyBase58Check     string
	CreatorCoinTransferTxindexMetadata CreatorCoinTransferTxindexMetadata
}

type CreatorCoinTransferTxindexMetadata struct {
	CreatorUsername            string
	CreatorCoinToTransferNanos int64
	DiamondLevel               int64
	PostHashHex                string
}
