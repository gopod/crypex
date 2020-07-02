package hitbtc

import "time"

// Report struct
type Report struct {
	ID              int64     `json:"id,string"`
	Side            string    `json:"side,required"`
	Type            string    `json:"type,required"`
	Price           float64   `json:"price,string"`
	Symbol          string    `json:"symbol,required"`
	Status          string    `json:"status,required"`
	Quantity        float64   `json:"quantity,string"`
	PostOnly        bool      `json:"postOnly,required"`
	StopPrice       bool      `json:"stopPrice,required"`
	CreatedAt       time.Time `json:"createdAt,required"`
	UpdatedAt       time.Time `json:"updatedAt,required"`
	ExpireTime      time.Time `json:"expireTime,required"`
	ReportType      string    `json:"reportType,required"`
	TimeInForce     string    `json:"timeInForce,required"`
	CumQuantity     string    `json:"cumQuantity,required"`
	OrderID         string    `json:"clientOrderId,required"`
	OriginalOrderID string    `json:"originalRequestClientOrderId,required"`
}

// ReportResponse struct
type ReportResponse Report

// SubscribeReports Subscribe to all reports [!Authenticate]
func (h *HitBTC) SubscribeReports() (reports <-chan interface{}, err error) {
	err = h.Subscribe("subscribeReports", nil)
	if err != nil {
		return
	}

	reports = h.Feeds.Sub(ReportsFeed)

	return
}

// Candle struct
type Candle struct {
	Min         float64   `json:"min,string"`
	Max         float64   `json:"max,string"`
	Open        float64   `json:"open,string"`
	Close       float64   `json:"close,string"`
	Volume      float64   `json:"volume,string"`      // Total trading amount within 24 hours in base currency
	VolumeQuote float64   `json:"volumeQuote,string"` // Total trading amount within 24 hours in quote currency
	Timestamp   time.Time `json:"timestamp,required"`
}

// Candles struct
type Candles []Candle

// CandlesResponse struct
type CandlesResponse struct {
	Data   Candles `json:"data,required"`
	Symbol string  `json:"symbol,required"`
	Period string  `json:"period,required"`
}

// SubscribeCandles subscribe to symbol candles
func (h *HitBTC) SubscribeCandles(symbol string, period string, limit int64) (candles <-chan interface{}, err error) {
	request := struct {
		Limit  int64  `json:"limit,required"`
		Period string `json:"period,required"`
		Symbol string `json:"symbol,required"`
	}{
		Symbol: symbol,
		Period: period,
		Limit:  limit,
	}

	err = h.Subscribe("subscribeCandles", &request)
	if err != nil {
		return
	}

	candles = h.Feeds.Sub(CandlesFeed + symbol)

	return
}

// UnsubscribeCandles unsubscribe from symbol candles
func (h *HitBTC) UnsubscribeCandles(symbol string) (err error) {
	request := struct {
		Symbol string `json:"symbol,required"`
	}{
		Symbol: symbol,
	}

	h.Feeds.Close(CandlesFeed + symbol)

	err = h.Subscribe("unsubscribeCandles", &request)

	return err
}

// OrderbookResponse struct
type OrderbookResponse struct {
	Ask      []SubtypeTrade `json:"ask,required"`
	Bid      []SubtypeTrade `json:"bid,required"`
	Symbol   string         `json:"symbol,required"`
	Sequence int64          `json:"sequence,required"` // used to see if the snapshot is the latest
}

// SubtypeTrade struct
type SubtypeTrade struct {
	Size  string `json:"size,required"`
	Price string `json:"price,required"`
}

// SubscribeOrderbook subscribe to symbol orderbook
func (h *HitBTC) SubscribeOrderbook(symbol string) (orderbook <-chan interface{}, err error) {
	request := struct {
		Symbol string `json:"symbol,required"`
	}{
		Symbol: symbol,
	}

	orderbook = h.Feeds.Sub(OrderbookFeed + symbol)

	err = h.Subscribe("subscribeOrderbook", &request)

	return
}

// UnsubscribeOrderbook unsubscribe from symbol orderbook
func (h *HitBTC) UnsubscribeOrderbook(symbol string) (err error) {
	request := struct {
		Symbol string `json:"symbol,required"`
	}{
		Symbol: symbol,
	}

	h.Feeds.Close(CandlesFeed + symbol)

	err = h.Subscribe("unsubscribeOrderbook", &request)

	return
}
