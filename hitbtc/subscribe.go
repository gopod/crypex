package hitbtc

import "time"

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
	Data   Candles `json:"data,required"`
	Symbol string  `json:"symbol,required"`
	Period string  `json:"period,required"`
}

type Candle struct {
	Min         float64   `json:"min,string"`
	Max         float64   `json:"max,string"`
	Open        float64   `json:"open,string"`
	Close       float64   `json:"close,string"`
	Volume      float64   `json:"volume,string"`      // Total trading amount within 24 hours in base currency
	VolumeQuote float64   `json:"volumeQuote,string"` // Total trading amount within 24 hours in quote currency
	Timestamp   time.Time `json:"timestamp,required"`
}

type Candles []Candle

type CandlesUpdate CandlesResponse
type CandlesSnapshot CandlesResponse

func (h *HitBTC) SubscribeCandles(symbol string, period string, limit int64) (
	update <-chan CandlesUpdate, snapshot <-chan CandlesSnapshot, err error) {
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
		return nil, nil, err
	}

	if _, ok := h.Feeds.CandlesFeed.Load(symbol); !ok {
		h.Feeds.CandlesFeed.Store(symbol, make(chan CandlesSnapshot))
	}

	if _, ok := h.Feeds.Notifications.CandlesFeed.Load(symbol); !ok {
		h.Feeds.Notifications.CandlesFeed.Store(symbol, make(chan CandlesUpdate))
	}

	snapshotChan, _ := h.Feeds.CandlesFeed.Load(symbol)
	updateChan, _ := h.Feeds.Notifications.CandlesFeed.Load(symbol)

	snapshot = snapshotChan.(chan CandlesSnapshot)
	update = updateChan.(chan CandlesUpdate)

	return
}

func (h *HitBTC) UnsubscribeCandles(symbol string) (err error) {
	request := struct {
		Symbol string `json:"symbol,required"`
	}{
		Symbol: symbol,
	}

	err = h.Subscribe("unsubscribeCandles", &request)
	if err != nil {
		return err
	}

	snapshot, _ := h.Feeds.CandlesFeed.Load(symbol)
	update, _ := h.Feeds.Notifications.CandlesFeed.Load(symbol)

	close(snapshot.(chan CandlesSnapshot))
	close(update.(chan CandlesUpdate))

	h.Feeds.Notifications.CandlesFeed.Delete(symbol)
	h.Feeds.CandlesFeed.Delete(symbol)

	return
}

type WSSubtypeTrade struct {
	Size  string `json:"size,required"`
	Price string `json:"price,required"`
}

type OrderbookUpdate struct {
	Ask      []WSSubtypeTrade `json:"ask,required"`
	Bid      []WSSubtypeTrade `json:"bid,required"`
	Symbol   string           `json:"symbol,required"`
	Sequence int64            `json:"sequence,required"` // used to see if the snapshot is the latest
}

type OrderbookSnapshot struct {
	Ask      []WSSubtypeTrade `json:"ask,required"`
	Bid      []WSSubtypeTrade `json:"bid,required"`
	Symbol   string           `json:"symbol,required"`
	Sequence int64            `json:"sequence,required"` // used to see if update is the latest received
}

func (h *HitBTC) SubscribeOrderbook(symbol string) (
	update <-chan OrderbookUpdate, snapshot <-chan OrderbookSnapshot, err error) {
	request := struct {
		Symbol string `json:"symbol,required"`
	}{
		Symbol: symbol,
	}

	err = h.Subscribe("subscribeOrderbook", &request)
	if err != nil {
		return nil, nil, err
	}

	if _, ok := h.Feeds.OrderbookFeed.Load(symbol); !ok {
		h.Feeds.OrderbookFeed.Store(symbol, make(chan OrderbookSnapshot))
	}

	if _, ok := h.Feeds.Notifications.OrderbookFeed.Load(symbol); !ok {
		h.Feeds.Notifications.OrderbookFeed.Store(symbol, make(chan OrderbookUpdate))
	}

	snapshotChan, _ := h.Feeds.OrderbookFeed.Load(symbol)
	updateChan, _ := h.Feeds.Notifications.OrderbookFeed.Load(symbol)

	snapshot = snapshotChan.(chan OrderbookSnapshot)
	update = updateChan.(chan OrderbookUpdate)

	return
}

func (h *HitBTC) UnsubscribeOrderbook(symbol string) (err error) {
	request := struct {
		Symbol string `json:"symbol,required"`
	}{
		Symbol: symbol,
	}

	err = h.Subscribe("unsubscribeOrderbook", &request)
	if err != nil {
		return err
	}

	snapshot, _ := h.Feeds.OrderbookFeed.Load(symbol)
	update, _ := h.Feeds.Notifications.OrderbookFeed.Load(symbol)

	close(snapshot.(chan OrderbookSnapshot))
	close(update.(chan OrderbookUpdate))

	h.Feeds.Notifications.OrderbookFeed.Delete(symbol)
	h.Feeds.OrderbookFeed.Delete(symbol)

	return
}
