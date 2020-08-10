package hitbtc

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/ramezanius/crypex/exchange"
)

// read redirects response to handler.
func (h *HitBTC) read(event *exchange.Event, handler exchange.HandlerFunc) {
	var redirect = func(response interface{}) {
		err := json.Unmarshal(event.Params.([]byte), &response)
		if err != nil {
			log.Fatalf("unmarshal response: [hitbtc]: %v", err)
		}

		go handler(response)
	}

	switch event.Method {
	case "error":
		redirect(&APIError{})

	case "report":
		redirect(&ReportsStream{})

	case "candles":
		go handler(event.Params.(*CandlesStream))

	case "updateCandles":
		redirect(&CandlesStream{})
	}
}

// SubscribeReports subscribes to the reports.
func (h *HitBTC) SubscribeReports(handler exchange.HandlerFunc) (err error) {
	err = h.Stream(exchange.StreamParams{
		Auth:     true,
		Method:   "subscribeReports",
		Location: exchange.TradingLoc,
	}, handler)

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
func (h *HitBTC) SubscribeCandles(params CandlesParams, handler exchange.HandlerFunc) (err error) {
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
				Candles: *snapshot,
			},
		}, handler)
	}

	err = h.Stream(exchange.StreamParams{
		Location: exchange.MarketLoc,
		Method:   "subscribeCandles",
		Params:   params,
	}, handler)

	return
}

// UnsubscribeCandles unsubscribes from candles.
func (h *HitBTC) UnsubscribeCandles(params CandlesParams) (err error) {
	err = h.Stream(exchange.StreamParams{
		Location: exchange.MarketLoc,
		Method:   "unsubscribeCandles",
		Params:   params,
	}, func(interface{}) {})

	return
}
