package hitbtc

import "time"

const (
	// Feed types
	ErrorFeed     = "ERROR"
	ReportsFeed   = "REPORTS"
	CandlesFeed   = "CANDLES:"
	OrderbookFeed = "ORDERBOOK:"

	// Exchange name
	Exchange = "HitBTC"

	// Test timeout
	Timeout = time.Second * 3

	// Exchange periods
	Period7Day     = "D7"
	Period1Day     = "D1"
	Period4Hour    = "H4"
	Period1Hour    = "H1"
	Period1Month   = "1M"
	Period1Minute  = "M1"
	Period15Minute = "M15"
	Period30Minute = "M30"

	// Exchange primary currencies
	USD = "USD"
	BTC = "BTC"
	ETH = "ETH"

	// Exchange websocket endpoint
	Endpoint = "wss://api.hitbtc.com/api/2/ws"
)
