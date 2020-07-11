package binance

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/cskr/pubsub"
	"github.com/ramezanius/jsonrpc2"
	"github.com/spf13/cast"
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

	// Exchange order time in forces
	FillOrKill        TimeInForce = "FOK"
	GoodTillCancel    TimeInForce = "GTC"
	ImmediateOrCancel TimeInForce = "IOC"

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
	// Websocket listen key
	ListenKey string
}

// Feeds struct
type Feeds struct {
	*pubsub.PubSub
}

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

// Asset struct
type Asset struct {
	Currency string `json:"asset,required"`

	Lock float64 `json:"locked,string"` // "0.00000000" may occurs error!
	Free float64 `json:"free,string"`
}

// Assets struct
type Assets []Asset

// Report struct
type Report struct {
	ID              int64       `json:"i,required"`
	Side            Side        `json:"S,required"`
	Type            Type        `json:"o,required"`
	Price           float64     `json:"p,string"`
	Symbol          string      `json:"s,required"`
	Status          string      `json:"x,required"`
	Quantity        float64     `json:"q,string"`
	StopPrice       float64     `json:"P,string"`
	CreatedAt       *time.Time  `json:"O,required"`
	UpdatedAt       *time.Time  `json:"T,required"`
	TimeInForce     TimeInForce `json:"f,required"`
	OrderID         string      `json:"c,required"`
	OriginalOrderID string      `json:"C,omitempty"`
}

func (r *ReportsResponse) UnmarshalJSON(data []byte) error {
	var v struct {
		ID              int64       `json:"i,required"`
		Side            Side        `json:"S,required"`
		Type            Type        `json:"o,required"`
		Price           float64     `json:"p,string"`
		Symbol          string      `json:"s,required"`
		Status          string      `json:"x,required"`
		Quantity        float64     `json:"q,string"`
		StopPrice       float64     `json:"P,string"`
		CreatedAt       int64       `json:"O,required"`
		UpdatedAt       int64       `json:"T,required"`
		TimeInForce     TimeInForce `json:"f,required"`
		OrderID         string      `json:"c,required"`
		OriginalOrderID string      `json:"C,omitempty"`
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
	r.Quantity = v.Quantity
	r.StopPrice = v.StopPrice
	r.TimeInForce = v.TimeInForce
	r.OrderID = v.OrderID
	r.OriginalOrderID = v.OriginalOrderID

	createdAt := time.Unix(v.CreatedAt, 0)
	r.CreatedAt = &createdAt

	updatedAt := time.Unix(v.UpdatedAt, 0)
	r.UpdatedAt = &updatedAt

	return nil
}

// Report struct
type Order struct {
	ID              int64       `json:"orderId,required"`
	Side            Side        `json:"side,required"`
	Type            Type        `json:"type,required"`
	Price           float64     `json:"price,string"`
	Symbol          string      `json:"symbol,required"`
	Status          string      `json:"status,required"`
	Quantity        float64     `json:"origQty,string"`
	StopPrice       float64     `json:"stopPrice,omitempty"`
	CreatedAt       *time.Time  `json:"time,omitempty"`
	UpdatedAt       *time.Time  `json:"updateTime,omitempty"`
	TimeInForce     TimeInForce `json:"timeInForce,required"`
	OrderID         string      `json:"clientOrderId,required"`
	OriginalOrderID string      `json:"origClientOrderId,omitempty"`
}

func (r *Order) UnmarshalJSON(data []byte) error {
	var v struct {
		ID              int64       `json:"orderId,required"`
		Side            Side        `json:"side,required"`
		Type            Type        `json:"type,required"`
		Price           float64     `json:"price,string"`
		Symbol          string      `json:"symbol,required"`
		Status          string      `json:"status,required"`
		Quantity        float64     `json:"origQty,string"`
		StopPrice       float64     `json:"stopPrice,omitempty"`
		CreatedAt       int64       `json:"time,omitempty"`
		UpdatedAt       int64       `json:"updateTime,omitempty"`
		TimeInForce     TimeInForce `json:"timeInForce,required"`
		OrderID         string      `json:"clientOrderId,required"`
		OriginalOrderID string      `json:"origClientOrderId,omitempty"`
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
	r.Quantity = v.Quantity
	r.StopPrice = v.StopPrice
	r.TimeInForce = v.TimeInForce
	r.OrderID = v.OrderID
	r.OriginalOrderID = v.OriginalOrderID

	createdAt := time.Unix(v.CreatedAt, 0)
	r.CreatedAt = &createdAt

	updatedAt := time.Unix(v.UpdatedAt, 0)
	r.UpdatedAt = &updatedAt

	return nil
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

func (r *CandlesResponse) UnmarshalJSON(data []byte) error {
	var v struct {
		Symbol string `json:"s,required"`
		Candle struct {
			StartAt     int64       `json:"t,required"`
			EndAt       int64       `json:"T,required"`
			Max         interface{} `json:"h,required"`
			Min         interface{} `json:"l,required"`
			Open        interface{} `json:"o,required"`
			Close       interface{} `json:"c,required"`
			Closed      bool        `json:"x,required"`
			Symbol      string      `json:"s,required"`
			Period      string      `json:"i,required"`
			Volume      interface{} `json:"v,required"`
			QuoteVolume interface{} `json:"q,required"`
		} `json:"k,required"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.Symbol = v.Symbol
	r.Candle.Closed = v.Candle.Closed
	r.Candle.Symbol = v.Candle.Symbol
	r.Candle.Period = v.Candle.Period

	r.Candle.Min = cast.ToFloat64(v.Candle.Min.(string))
	r.Candle.Max = cast.ToFloat64(v.Candle.Max.(string))
	r.Candle.Open = cast.ToFloat64(v.Candle.Open.(string))
	r.Candle.Close = cast.ToFloat64(v.Candle.Close.(string))
	r.Candle.Volume = cast.ToFloat64(v.Candle.Volume.(string))
	r.Candle.QuoteVolume = cast.ToFloat64(v.Candle.QuoteVolume.(string))

	endAt := time.Unix(v.Candle.EndAt, 0)
	startAt := time.Unix(v.Candle.StartAt, 0)

	r.Candle.EndAt = &endAt
	r.Candle.StartAt = &startAt

	return nil
}
