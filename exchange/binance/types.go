package binance

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"go.uber.org/ratelimit"
)

const (
	// Exchange name
	Exchange = "Binance"

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

	// Exchange order time in forces
	FillOrKill        TimeInForce = "FOK"
	GoodTillCancel    TimeInForce = "GTC"
	ImmediateOrCancel TimeInForce = "IOC"

	// Exchange base currencies
	USD = "usdt"
	BTC = "btc"
	ETH = "eth"
	XRP = "xrp"
	BNB = "bnb"
)

// Binance exchange struct
type Binance struct {
	sync.RWMutex

	// Feeds channels
	Feeds *Feeds
	// OnErr error handler
	OnErr func(error)
	// connections websocket pool
	connections map[string]*websocket.Conn
	// publicLimit, tradingLimit, and wsLimit rate limits
	publicLimit, tradingLimit, wsLimit ratelimit.Limiter

	// ListenKey websocket listen key
	ListenKey string
	// PublicKey, SecretKey API keys
	PublicKey, SecretKey string
}

// Feeds stream feeds struct
type Feeds struct {
	mu sync.Mutex

	// Reports feed
	Reports chan *ReportsStream
	// Candles feed
	Candles chan *CandlesStream
}

// Clock custom websocket rate limit type
type Clock struct{}

// Type order type
type TimeInForce string

// Period candles period type
type Period string

// Side order side type
type Side string

// Type order type
type Type string

// Symbol struct
type Symbol struct {
	ID    string `json:"symbol,required"`
	Base  string `json:"baseAsset,required"`
	Quote string `json:"quoteAsset,required"`
}

// Symbols struct
type Symbols []Symbol

// SymbolsResponse struct
type SymbolsResponse struct {
	Symbols Symbols `json:"symbols,required"`
}

// Asset struct
type Asset struct {
	Currency string `json:"asset,required"`

	Lock float64 `json:"locked,string"`
	Free float64 `json:"free,string"`
}

// Assets struct
type Assets []Asset

// AssetsResponse struct
type AssetsResponse struct {
	Assets Assets `json:"balances,required"`
}

// Report struct
type Report struct {
	ID              int64       `json:"orderId,required"`
	Side            Side        `json:"side,required"`
	Type            Type        `json:"type,required"`
	Price           float64     `json:"price,string"`
	Symbol          string      `json:"symbol,required"`
	Status          string      `json:"status,required"`
	Quantity        float64     `json:"quantity,string"`
	StopPrice       float64     `json:"stopPrice,string"`
	UpdatedAt       time.Time   `json:"createdAt,required"`
	CreatedAt       time.Time   `json:"updatedAt,required"`
	TimeInForce     TimeInForce `json:"timeInForce,required"`
	OrderID         string      `json:"clientOrderId,required"`
	OriginalOrderID string      `json:"origClientOrderId,omitempty"`
}

// ReportsStream response
type ReportsStream Report

func (r *ReportsStream) UnmarshalJSON(data []byte) error {
	var v struct {
		ID              int64       `json:"i,required"`
		Side            Side        `json:"S,required"`
		Type            Type        `json:"o,required"`
		Price           float64     `json:"p,string"`
		Symbol          string      `json:"s,required"`
		Status          string      `json:"x,required"`
		OrderID         string      `json:"c,required"`
		Quantity        float64     `json:"q,string"`
		StopPrice       float64     `json:"P,string"`
		CreatedAt       int         `json:"O,required"`
		UpdatedAt       int         `json:"T,required"`
		TimeInForce     TimeInForce `json:"f,required"`
		OriginalOrderID string      `json:"C,omitempty"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v.UpdatedAt == -1 {
		v.UpdatedAt = v.CreatedAt
	}

	r.ID = v.ID
	r.Side = v.Side
	r.Type = v.Type
	r.Price = v.Price
	r.Symbol = v.Symbol
	r.Status = v.Status
	r.OrderID = v.OrderID
	r.Quantity = v.Quantity
	r.StopPrice = v.StopPrice
	r.TimeInForce = v.TimeInForce
	r.OriginalOrderID = v.OriginalOrderID
	r.CreatedAt = time.Unix(cast.ToInt64(strconv.Itoa(v.CreatedAt)[:10]), 0)
	r.UpdatedAt = time.Unix(cast.ToInt64(strconv.Itoa(v.UpdatedAt)[:10]), 0)

	return nil
}

// Order struct
type Order struct {
	ID              int64       `json:"orderId,required"`
	Side            Side        `json:"side,required"`
	Type            Type        `json:"type,required"`
	Price           float64     `json:"price,string"`
	Symbol          string      `json:"symbol,required"`
	Status          string      `json:"status,required"`
	Quantity        float64     `json:"origQty,string"`
	StopPrice       float64     `json:"stopPrice,omitempty"`
	TransactAt      time.Time   `json:"transactTime,omitempty"`
	TimeInForce     TimeInForce `json:"timeInForce,required"`
	OrderID         string      `json:"clientOrderId,required"`
	OriginalOrderID string      `json:"origClientOrderId,omitempty"`
}

// OrderResponse struct
type OrderResponse Order

func (r *OrderResponse) UnmarshalJSON(data []byte) error {
	var v struct {
		*Order
		TransactAt int `json:"transactTime,omitempty"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.ID = v.ID
	r.Side = v.Side
	r.Type = v.Type
	r.Price = v.Price
	r.Symbol = v.Symbol
	r.Status = v.Status
	r.OrderID = v.OrderID
	r.Quantity = v.Quantity
	r.StopPrice = v.StopPrice
	r.TimeInForce = v.TimeInForce
	r.OriginalOrderID = v.OriginalOrderID

	if v.TransactAt != 0 {
		r.TransactAt = time.Unix(cast.ToInt64(strconv.Itoa(v.TransactAt)[:10]), 0)
	}

	return nil
}

// NewOrder struct
type NewOrder struct {
	Side  Side    `json:"side,required"`
	Type  Type    `json:"type,required"`
	Price float64 `json:"price,string"`

	Symbol   string  `json:"symbol,required"`
	Quantity float64 `json:"quantity,string"`
	OrderID  string  `json:"newClientOrderId,omitempty"`

	StopPrice   float64     `json:"stopPrice,omitempty"`
	TimeInForce TimeInForce `json:"timeInForce,omitempty"`
}

// Candle struct
type Candle struct {
	Min         float64   `json:"min,string"`
	Max         float64   `json:"max,string"`
	Open        float64   `json:"open,string"`
	Close       float64   `json:"close,string"`
	Volume      float64   `json:"volume,string"`
	Timestamp   time.Time `json:"timestamp,required"`
	VolumeQuote float64   `json:"volumeQuote,string"`
}

// Candles struct
type Candles []Candle

// CandlesStream struct
type CandlesStream struct {
	Period  Period  `json:"-"`
	Symbol  string  `json:"s,required"`
	Candles Candles `json:"k,required"`
}

func (r *CandlesStream) UnmarshalJSON(data []byte) error {
	var v struct {
		Symbol string `json:"s,required"`
		Candle struct {
			Timestamp   int         `json:"T,required"`
			Max         interface{} `json:"h,required"`
			Min         interface{} `json:"l,required"`
			Open        interface{} `json:"o,required"`
			Close       interface{} `json:"c,required"`
			Closed      bool        `json:"x,required"`
			Symbol      string      `json:"s,required"`
			Period      string      `json:"i,required"`
			Volume      interface{} `json:"v,required"`
			VolumeQuote interface{} `json:"q,required"`
		} `json:"k,required"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.Symbol = v.Symbol
	r.Period = Period(v.Candle.Period)

	r.Candles = append(r.Candles, Candle{
		Max:         cast.ToFloat64(v.Candle.Max.(string)),
		Min:         cast.ToFloat64(v.Candle.Min.(string)),
		Open:        cast.ToFloat64(v.Candle.Open.(string)),
		Close:       cast.ToFloat64(v.Candle.Close.(string)),
		Volume:      cast.ToFloat64(v.Candle.Volume.(string)),
		VolumeQuote: cast.ToFloat64(v.Candle.VolumeQuote.(string)),
		Timestamp:   time.Unix(cast.ToInt64(strconv.Itoa(v.Candle.Timestamp)[:10]), 0),
	})

	return nil
}

type CandlesResponse Candles

func (r *CandlesResponse) UnmarshalJSON(data []byte) error {
	var v [][]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	for _, candle := range v {
		*r = append(*r, Candle{
			Max:         cast.ToFloat64(candle[2]),
			Min:         cast.ToFloat64(candle[3]),
			Open:        cast.ToFloat64(candle[1]),
			Close:       cast.ToFloat64(candle[4]),
			Volume:      cast.ToFloat64(candle[5]),
			VolumeQuote: cast.ToFloat64(candle[7]),
			Timestamp:   time.Unix(cast.ToInt64(strconv.Itoa(int(candle[6].(float64)))[:10]), 0),
		})
	}

	return nil
}

// CandlesParams struct
type CandlesParams struct {
	Snapshot bool `json:"-"`

	Limit  int    `json:"limit"`
	Symbol string `json:"symbol"`
	Period Period `json:"interval"`
}
