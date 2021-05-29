package models

type CoinEntry struct {
	CreatorBasisPoints      int64
	BitCloutLockedNanos     int64
	NumberOfHolders         int64
	CoinsInCirculationNanos int64
	CoinWatermarkNanos      int64
}

// It means that it's multiplied by 1 billion and divided by 1 million. 1e9 means 1 * 10 to the 9th power, which is 1 billion (1000000000).

// return this.profile.CoinEntry.CoinsInCirculationNanos / 1e9;

// usdMarketCap
// nanosToUSDNumber(this.coinsInCirculation() * this.profile.CoinPriceBitCloutNa0nos),

// nanosToUSDNumber
// nanos / this.nanosPerUSDExchangeRate;
