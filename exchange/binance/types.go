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

	// Websocket connections
	connections map[string]*websocket.Conn
	// Public, Trading rate limits
	publicLimit, tradingLimit ratelimit.Limiter

	// Websocket listen key
	ListenKey string
	// Public API key, Secret API key
	PublicKey, SecretKey string
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
		Quantity        float64     `json:"q,string"`
		StopPrice       float64     `json:"P,string"`
		CreatedAt       int         `json:"O,required"`
		UpdatedAt       int         `json:"T,required"`
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
	r.OrderID = v.OrderID
	r.Quantity = v.Quantity
	r.StopPrice = v.StopPrice
	r.TimeInForce = v.TimeInForce
	r.OriginalOrderID = v.OriginalOrderID

	if v.UpdatedAt == -1 {
		v.UpdatedAt = v.CreatedAt
	}

	createdAt := time.Unix(cast.ToInt64(strconv.Itoa(v.CreatedAt)[:10]), 0)
	updatedAt := time.Unix(cast.ToInt64(strconv.Itoa(v.UpdatedAt)[:10]), 0)

	r.CreatedAt = &createdAt
	r.UpdatedAt = &updatedAt

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
	TransactAt      *time.Time  `json:"transactTime,omitempty"`
	TimeInForce     TimeInForce `json:"timeInForce,required"`
	OrderID         string      `json:"clientOrderId,required"`
	OriginalOrderID string      `json:"origClientOrderId,omitempty"`
}

// OrderResponse struct
type OrderResponse Order

func (r *OrderResponse) UnmarshalJSON(data []byte) error {
	var v struct {
		ID              int64       `json:"orderId,required"`
		Side            Side        `json:"side,required"`
		Type            Type        `json:"type,required"`
		Price           float64     `json:"price,string"`
		Symbol          string      `json:"symbol,required"`
		Status          string      `json:"status,required"`
		Quantity        float64     `json:"origQty,string"`
		StopPrice       float64     `json:"stopPrice,omitempty"`
		TransactAt      int         `json:"transactTime,omitempty"`
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
	r.OrderID = v.OrderID
	r.Quantity = v.Quantity
	r.StopPrice = v.StopPrice
	r.TimeInForce = v.TimeInForce
	r.OriginalOrderID = v.OriginalOrderID

	if v.TransactAt != 0 {
		transactAt := time.Unix(cast.ToInt64(strconv.Itoa(v.TransactAt)[:10]), 0)
		r.TransactAt = &transactAt
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
	StartAt     *time.Time `json:"t,required"`
	EndAt       *time.Time `json:"T,required"`
	Max         float64    `json:"h,string"`
	Min         float64    `json:"l,string"`
	Open        float64    `json:"o,string"`
	Close       float64    `json:"c,string"`
	Closed      bool       `json:"x,required"`
	Volume      float64    `json:"v,string"`
	QuoteVolume float64    `json:"q,string"`
}

// Candles struct
type Candles []Candle

func (r *Candle) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.Min = cast.ToFloat64(v[3])
	r.Max = cast.ToFloat64(v[2])
	r.Open = cast.ToFloat64(v[1])
	r.Close = cast.ToFloat64(v[4])
	r.Volume = cast.ToFloat64(v[5])
	r.QuoteVolume = cast.ToFloat64(v[7])

	endAt := time.Unix(cast.ToInt64(strconv.Itoa(int(v[6].(float64)))[:10]), 0)
	startAt := time.Unix(cast.ToInt64(strconv.Itoa(int(v[0].(float64)))[:10]), 0)

	r.EndAt = &endAt
	r.StartAt = &startAt

	return nil
}

// CandlesStream struct
type CandlesStream struct {
	Period Period `json:"-"`
	Symbol string `json:"s,required"`
	Candle Candle `json:"k,required"`
}

func (r *CandlesStream) UnmarshalJSON(data []byte) error {
	var v struct {
		Symbol string `json:"s,required"`
		Candle struct {
			StartAt     int         `json:"t,required"`
			EndAt       int         `json:"T,required"`
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
	r.Period = Period(v.Candle.Period)

	r.Candle.Closed = v.Candle.Closed
	r.Candle.Min = cast.ToFloat64(v.Candle.Min.(string))
	r.Candle.Max = cast.ToFloat64(v.Candle.Max.(string))
	r.Candle.Open = cast.ToFloat64(v.Candle.Open.(string))
	r.Candle.Close = cast.ToFloat64(v.Candle.Close.(string))
	r.Candle.Volume = cast.ToFloat64(v.Candle.Volume.(string))
	r.Candle.QuoteVolume = cast.ToFloat64(v.Candle.QuoteVolume.(string))

	endAt := time.Unix(cast.ToInt64(strconv.Itoa(v.Candle.EndAt)[:10]), 0)
	startAt := time.Unix(cast.ToInt64(strconv.Itoa(v.Candle.StartAt)[:10]), 0)

	r.Candle.EndAt = &endAt
	r.Candle.StartAt = &startAt

	return nil
}

// CandlesParams struct
type CandlesParams struct {
	Limit  int    `json:"limit"`
	Symbol string `json:"symbol"`
	Period Period `json:"interval"`
}
