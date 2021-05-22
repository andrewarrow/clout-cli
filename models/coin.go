package models

type CoinEntry struct {
	CreatorBasisPoints      int64
	BitCloutLockedNanos     int64
	NumberOfHolders         int64
	CoinsInCirculationNanos int64
	CoinWatermarkNanos      int64
}
