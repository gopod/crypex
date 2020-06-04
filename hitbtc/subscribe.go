package hitbtc

import (
	"time"
)

type Report struct {
	ID          int64     `json:"id,string"`
	Side        string    `json:"side"`
	Type        string    `json:"type"`
	Price       float64   `json:"price,string"`
	Symbol      string    `json:"symbol"`
	Status      string    `json:"status"`
	Quantity    float64   `json:"quantity,string"`
	StopPrice   bool      `json:"stopPrice"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ExpireTime  time.Time `json:"expireTime"`
	TimeInForce string    `json:"timeInForce"`
	OrderID     string    `json:"clientOrderId"`
}

type Reports []Report

type ReportsUpdate Report
type ReportsSnapshot Reports

func (h *HitBTC) SubscribeReports() (update <-chan ReportsUpdate, snapshot <-chan ReportsSnapshot, err error) {
	err = h.Subscribe("subscribeReports", nil)
	if err != nil {
		return nil, nil, err
	}

	update = h.Feeds.Notifications.ReportsFeed
	snapshot = h.Feeds.ReportsFeed

	return
}

type CandlesResponse struct {
	Data   Candles `json:"data"`
	Symbol string  `json:"symbol"`
	Period string  `json:"period"`
}

type Candle struct {
	Min         float64   `json:"min,string"`
	Max         float64   `json:"max,string"`
	Open        float64   `json:"open,string"`
	Close       float64   `json:"close,string"`
	Volume      float64   `json:"volume,string"`      // Total trading amount within 24 hours in base currency
	VolumeQuote float64   `json:"volumeQuote,string"` // Total trading amount within 24 hours in quote currency
	Timestamp   time.Time `json:"timestamp"`
}

type Candles []Candle

type CandlesUpdate CandlesResponse
type CandlesSnapshot CandlesResponse

func (h *HitBTC) SubscribeCandles(symbol string, period string, limit int64) (
	update <-chan CandlesUpdate, snapshot <-chan CandlesSnapshot, err error) {
	request := struct {
		Limit  int64  `json:"limit"`
		Period string `json:"period"`
		Symbol string `json:"symbol"`
	}{
		Symbol: symbol,
		Period: period,
		Limit:  limit,
	}

	err = h.Subscribe("subscribeCandles", &request)
	if err != nil {
		return nil, nil, err
	}

	if h.Feeds.Notifications.CandlesFeed[symbol] == nil {
		h.Feeds.Notifications.CandlesFeed[symbol] = make(chan CandlesUpdate)
	}

	if h.Feeds.CandlesFeed[symbol] == nil {
		h.Feeds.CandlesFeed[symbol] = make(chan CandlesSnapshot)
	}

	update = h.Feeds.Notifications.CandlesFeed[symbol]
	snapshot = h.Feeds.CandlesFeed[symbol]

	return
}

func (h *HitBTC) UnsubscribeCandles(symbol string) (err error) {
	request := struct {
		Symbol string `json:"symbol"`
	}{
		Symbol: symbol,
	}

	err = h.Subscribe("unsubscribeCandles", &request)
	if err != nil {
		return err
	}

	close(h.Feeds.Notifications.CandlesFeed[symbol])
	delete(h.Feeds.Notifications.CandlesFeed, symbol)

	close(h.Feeds.CandlesFeed[symbol])
	delete(h.Feeds.CandlesFeed, symbol)

	return
}

type WSSubtypeTrade struct {
	Size  string `json:"size"`
	Price string `json:"price"`
}

type OrderbookUpdate struct {
	Ask      []WSSubtypeTrade `json:"ask"`
	Bid      []WSSubtypeTrade `json:"bid"`
	Symbol   string           `json:"symbol"`
	Sequence int64            `json:"sequence"` // used to see if the snapshot is the latest
}

type OrderbookSnapshot struct {
	Ask      []WSSubtypeTrade `json:"ask"`
	Bid      []WSSubtypeTrade `json:"bid"`
	Symbol   string           `json:"symbol"`
	Sequence int64            `json:"sequence"` // used to see if update is the latest received
}

func (h *HitBTC) SubscribeOrderbook(symbol string) (
	update <-chan OrderbookUpdate, snapshot <-chan OrderbookSnapshot, err error) {
	request := struct {
		Symbol string `json:"symbol"`
	}{
		Symbol: symbol,
	}

	err = h.Subscribe("subscribeOrderbook", &request)
	if err != nil {
		return nil, nil, err
	}

	if h.Feeds.Notifications.OrderbookFeed[symbol] == nil {
		h.Feeds.Notifications.OrderbookFeed[symbol] = make(chan OrderbookUpdate)
	}

	if h.Feeds.OrderbookFeed[symbol] == nil {
		h.Feeds.OrderbookFeed[symbol] = make(chan OrderbookSnapshot)
	}

	update = h.Feeds.Notifications.OrderbookFeed[symbol]
	snapshot = h.Feeds.OrderbookFeed[symbol]

	return
}

func (h *HitBTC) UnsubscribeOrderbook(symbol string) (err error) {
	request := struct {
		Symbol string `json:"symbol"`
	}{
		Symbol: symbol,
	}

	err = h.Subscribe("unsubscribeOrderbook", &request)
	if err != nil {
		return err
	}

	close(h.Feeds.Notifications.OrderbookFeed[symbol])
	delete(h.Feeds.Notifications.OrderbookFeed, symbol)

	close(h.Feeds.OrderbookFeed[symbol])
	delete(h.Feeds.OrderbookFeed, symbol)

	return
}
