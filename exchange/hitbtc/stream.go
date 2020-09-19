package hitbtc

import (
	"encoding/json"
	"strings"

	"github.com/gopod/crypex/exchange"
)

// read redirects response to handler.
func (h *HitBTC) read(event *exchange.Event) {
	switch event.Method {
	case "error":
		object := &APIError{}

		err := json.Unmarshal(event.Params.([]byte), &object)
		if err != nil {
			h.OnErr(err)
		}

		h.OnErr(object)

	case "report":
		object := &ReportsStream{}

		err := json.Unmarshal(event.Params.([]byte), &object)
		if err != nil {
			h.OnErr(err)
		}

		h.Feeds.Reports <- object

	case "candles":
		h.Feeds.Candles <- event.Params.(*CandlesStream)

	case "updateCandles":
		object := &CandlesStream{}

		err := json.Unmarshal(event.Params.([]byte), &object)
		if err != nil {
			h.OnErr(err)
		}

		h.Feeds.Candles <- object
	}
}

// SubscribeReports subscribes to the reports.
func (h *HitBTC) SubscribeReports() (err error) {
	err = h.Stream(exchange.StreamParams{
		Auth:     true,
		Method:   "subscribeReports",
		Location: exchange.TradingLoc,
	})

	return
}

// UnsubscribeCandles unsubscribes from reports.
func (h *HitBTC) UnsubscribeReports() (err error) {
	h.Lock()
	defer h.Unlock()

	err = exchange.CloseConn(h.connections[exchange.TradingLoc])
	h.connections[exchange.TradingLoc] = nil

	return
}

// SubscribeCandles subscribes to the candles.
func (h *HitBTC) SubscribeCandles(params CandlesParams) (err error) {
	if params.Limit <= 0 {
		params.Limit = 100
	}

	if params.Snapshot {
		snapshot, err := h.GetCandles(params)
		if err != nil {
			return err
		}

		h.read(&exchange.Event{
			Method: "candles",
			Params: &CandlesStream{
				Period:  params.Period,
				Symbol:  strings.ToUpper(params.Symbol),
				Candles: Candles(*snapshot),
			},
		})
	}

	err = h.Stream(exchange.StreamParams{
		Location: exchange.MarketLoc,
		Method:   "subscribeCandles",
		Params:   params,
	})

	return
}

// UnsubscribeCandles unsubscribes from candles.
func (h *HitBTC) UnsubscribeCandles(params CandlesParams) (err error) {
	err = h.Stream(exchange.StreamParams{
		Location: exchange.MarketLoc,
		Method:   "unsubscribeCandles",
		Params:   params,
	})

	return
}
