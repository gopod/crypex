package hitbtc

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
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
	USD = "USD"
	BTC = "BTC"
	ETH = "ETH"
	XRP = "XRP"
)

// HitBTC exchange struct
type HitBTC struct {
	// Websocket connections
	connections map[string]*websocket.Conn

	// Public API key, Secret API key
	PublicKey, SecretKey string
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
	ID    int64   `json:"id,int,string"`
	Side  Side    `json:"side,required"`
	Type  Type    `json:"type,required"`
	Price float64 `json:"price,string"`

	Symbol    string  `json:"symbol,required"`
	Status    string  `json:"status,required"`
	Quantity  float64 `json:"quantity,string"`
	StopPrice bool    `json:"stopPrice,omitempty"`

	CreatedAt time.Time `json:"createdAt,required"`
	UpdatedAt time.Time `json:"updatedAt,required"`

	TimeInForce     string `json:"timeInForce,omitempty"`
	OrderID         string `json:"clientOrderId,required"`
	OriginalOrderID string `json:"originalRequestClientOrderId,omitempty"`
}

func (r *ReportsResponse) UnmarshalJSON(data []byte) error {
	var v struct {
		ID    interface{} `json:"id,required"`
		Side  Side        `json:"side,required"`
		Type  Type        `json:"type,required"`
		Price float64     `json:"price,string"`

		Symbol    string  `json:"symbol,required"`
		Status    string  `json:"status,required"`
		Quantity  float64 `json:"quantity,string"`
		StopPrice bool    `json:"stopPrice,omitempty"`

		CreatedAt time.Time `json:"createdAt,required"`
		UpdatedAt time.Time `json:"updatedAt,required"`

		TimeInForce     string `json:"timeInForce,omitempty"`
		OrderID         string `json:"clientOrderId,required"`
		OriginalOrderID string `json:"originalRequestClientOrderId,omitempty"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.ID = cast.ToInt64(v.ID)
	r.Side = v.Side
	r.Type = v.Type
	r.Price = v.Price
	r.Symbol = v.Symbol
	r.Status = v.Status
	r.OrderID = v.OrderID
	r.Quantity = v.Quantity
	r.StopPrice = v.StopPrice
	r.CreatedAt = v.CreatedAt
	r.UpdatedAt = v.UpdatedAt
	r.TimeInForce = v.TimeInForce
	r.OriginalOrderID = v.OriginalOrderID

	return nil
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

func (r *CandlesResponse) UnmarshalJSON(data []byte) error {
	var v struct {
		Candles Candles `json:"data,required"`
		Symbol  string  `json:"symbol,required"`
		Period  string  `json:"period,required"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.Symbol = v.Symbol
	r.Period = v.Period

	if len(v.Candles) != 0 {
		r.Candle = v.Candles[0]
	}

	return nil
}
