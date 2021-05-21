package models

type SubmitPost struct {
	UpdaterPublicKeyBase58Check string
	PostHashHexToModify         string
	ParentStakeID               string
	Title                       string
	BodyObj                     BodyObj
	RecloutedPostHashHex        string
	IsHidden                    bool
	MinFeeRateNanosPerKB        int64
}

type BodyObj struct {
	Body string
}

type SubmitPostResponse struct {
	TransactionHex string
}
