package models

type Rate struct {
	SatoshisPerBitCloutExchangeRate int64 `json:"SatoshisPerBitCloutExchangeRate"`
	NanosSold                       int64
	USDCentsPerBitcoinExchangeRate  int64
}
