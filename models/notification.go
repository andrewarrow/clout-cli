package models

type NotificationList struct {
	Notifications []Notification
}
type Notification struct {
	Metadata Metadata
}

type Metadata struct {
	TxnType string
}
