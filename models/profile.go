package models

type SingleProfile struct {
	Profile ProfileEntryResponse
}
type ProfileEntryResponse struct {
	PublicKeyBase58Check   string
	Username               string
	Description            string
	ProfilePic             string
	CoinEntry              CoinEntry
	CoinPriceBitCloutNanos int64
}

func (p ProfileEntryResponse) MarketCap() float64 {
	coins := float64(p.CoinEntry.CoinsInCirculationNanos) / 1000000000.0
	marketCap := coins * float64(p.CoinPriceBitCloutNanos)
	return marketCap / 1000000000.0
}
