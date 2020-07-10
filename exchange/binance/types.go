package binance

import (
	"sync"
	"time"

	"github.com/cskr/pubsub"
	"github.com/ramezanius/jsonrpc2"
)

const (
	// Exchange name
	Exchange = "Binance"

	// Exchange feeds
	ReportsFeed = "REPORTS"
	CandlesFeed = "CANDLES::"

	// Exchange periods
	Period1Day     Period = "1d"
	Period3Day     Period = "3d"
	Period7Day     Period = "1w"
	Period1Hour    Period = "1h"
	Period2Hour    Period = "2h"
	Period4Hour    Period = "4h"
	Period6Hour    Period = "6h"
	Period8Hour    Period = "8h"
	Period12Hour   Period = "12h"
	Period1Month   Period = "1M"
	Period1Minute  Period = "1m"
	Period3Minute  Period = "3m"
	Period5Minute  Period = "5m"
	Period15Minute Period = "15m"
	Period30Minute Period = "30m"

	// Exchange order sides
	Buy  Side = "BUY"
	Sell Side = "SELL"

	// Exchange order types
	Limit         Type = "LIMIT"
	Market        Type = "MARKET"
	StopLoss      Type = "STOP_LOSS"
	StopLossLimit Type = "STOP_LOSS_LIMIT"

	// Exchange base currencies
	USD  = "usdt"
	BTC  = "btc"
	ETH  = "eth"
	BNB  = "bnb"
	Demo = BTC + USD
)

// Binance exchange
type Binance struct {
	sync.RWMutex

	// Feeds collection
	Feeds *Feeds
	// Websocket connection
	Connection *jsonrpc2.Conn

	// Public API key, Secret API key
	PublicKey, SecretKey string
	// SecretKey signature, Websocket listen key
	ListenKey, Signature string
}

// Feeds struct
type Feeds struct {
	*pubsub.PubSub
}

// Period candles period type
type Period string

// Side order side type
type Side string

// Type order type
type Type string

// Report struct
type Report struct {
	ID              int64   `json:"i"`
	Side            string  `json:"S,required"`
	Type            string  `json:"o,required"`
	Price           float64 `json:"p,string"`
	Symbol          string  `json:"s,required"`
	Status          string  `json:"x,required"`
	Quantity        float64 `json:"q,string"`
	StopPrice       float64 `json:"P,string"`
	CreatedAt       int     `json:"O,required"`
	UpdatedAt       int     `json:"T,required"`
	TimeInForce     string  `json:"f,required"`
	OrderID         string  `json:"c,required"`
	OriginalOrderID string  `json:"C,required"`
}

// Candle struct
type Candle struct {
	StartAt     *time.Time `json:"t,required"`
	EndAt       *time.Time `json:"T,required"`
	Max         float64    `json:"h,string"`
	Min         float64    `json:"l,string"`
	Open        float64    `json:"o,string"`
	Close       float64    `json:"c,string"`
	Closed      bool       `json:"x,required"`
	Symbol      string     `json:"s,required"`
	Period      string     `json:"i,required"`
	Volume      float64    `json:"v,string"`
	QuoteVolume float64    `json:"q,string"`
}

// Candles struct
type Candles []Candle
