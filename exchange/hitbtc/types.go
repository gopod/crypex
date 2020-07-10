package hitbtc

import (
	"sync"
	"time"

	"github.com/cskr/pubsub"
	"github.com/sourcegraph/jsonrpc2"
)

const (
	// Exchange name
	Exchange = "HitBTC"

	// Exchange periods
	Period7Day     Period = "D7"
	Period1Day     Period = "D1"
	Period4Hour    Period = "H4"
	Period1Hour    Period = "H1"
	Period1Month   Period = "1M"
	Period1Minute  Period = "M1"
	Period15Minute Period = "M15"
	Period30Minute Period = "M30"

	// Exchange order sides
	Buy  Side = "buy"
	Sell Side = "sell"

	// Exchange order types
	Limit      Type = "limit"
	Market     Type = "market"
	StopLimit  Type = "stopLimit"
	StopMarket Type = "stopMarket"

	// Exchange base currencies
	USD  = "USD"
	BTC  = "BTC"
	ETH  = "ETH"
	Demo = BTC + USD
)

// HitBTC exchange struct
type HitBTC struct {
	sync.RWMutex

	// Feeds collection
	Feeds *Feeds
	// Websocket connection
	Connection *jsonrpc2.Conn

	// Public API key, Secret API key
	PublicKey, SecretKey string
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

// Symbol struct
type Symbol struct {
	ID    string `json:"id,required"`
	Base  string `json:"baseCurrency,required"`
	Quote string `json:"quoteCurrency,required"`

	TickSize             string `json:"tickSize,required"`
	FeeCurrency          string `json:"feeCurrency,required"`
	QuantityIncrement    string `json:"quantityIncrement,required"`
	TakeLiquidityRate    string `json:"takeLiquidityRate,required"`
	ProvideLiquidityRate string `json:"provideLiquidityRate,required"`
}

// Symbols struct
type Symbols []Symbol

// Asset struct
type Asset struct {
	Currency string `json:"currency,required"`

	Lock float64 `json:"reserved,string"`
	Free float64 `json:"available,string"`
}

// Assets struct
type Assets []Asset

// Report struct
type Report struct {
	ID    int64   `json:"id,string"`
	Side  Side    `json:"side,required"`
	Type  Type    `json:"type,required"`
	Price float64 `json:"price,string"`

	Symbol    string  `json:"symbol,required"`
	Status    string  `json:"status,required"`
	Quantity  float64 `json:"quantity,string"`
	StopPrice bool    `json:"stopPrice,required"`

	CreatedAt time.Time `json:"createdAt,required"`
	UpdatedAt time.Time `json:"updatedAt,required"`

	TimeInForce     string `json:"timeInForce,required"`
	OrderID         string `json:"clientOrderId,required"`
	OriginalOrderID string `json:"originalRequestClientOrderId,required"`
}

// Candle struct
type Candle struct {
	Min   float64 `json:"min,string"`
	Max   float64 `json:"max,string"`
	Open  float64 `json:"open,string"`
	Close float64 `json:"close,string"`

	Volume      float64   `json:"volume,string"`      // Total trading amount within 24 hours in base currency
	VolumeQuote float64   `json:"volumeQuote,string"` // Total trading amount within 24 hours in quote currency
	Timestamp   time.Time `json:"timestamp,required"`
}

// Candles struct
type Candles []Candle
