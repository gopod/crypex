package hitbtc

import (
	"encoding/json"
	"log"

	"github.com/ramezanius/crypex/exchange"
)

// read redirects response to handler.
func (h *HitBTC) read(event *exchange.Event, handler exchange.HandlerFunc) {
	var redirect = func(response interface{}) {
		err := json.Unmarshal(event.Params, &response)
		if err != nil {
			log.Fatalf("unmarshal response: [hitbtc]: %v", err)
		}

		go handler(response)
	}

	switch event.Method {
	case "report":
		redirect(&ReportsStream{})

	case "updateCandles":
		redirect(&CandlesStream{})
	}
}

// SubscribeReports subscribes to the reports.
func (h *HitBTC) SubscribeReports(handler exchange.HandlerFunc) (err error) {
	err = h.Stream(exchange.StreamParams{
		Auth:     true,
		Location: "/trading",
		Method:   "subscribeReports",
	}, handler)

	return
}

// UnsubscribeCandles unsubscribes from reports.
func (h *HitBTC) UnsubscribeReports() (err error) {
	h.Lock()
	defer h.Unlock()

	err = exchange.CloseConn(h.connections[h.SecretKey])
	h.connections[h.SecretKey] = nil

	return
}

// SubscribeCandles subscribes to the candles.
func (h *HitBTC) SubscribeCandles(params CandlesParams, handler exchange.HandlerFunc) (err error) {
	err = h.Stream(exchange.StreamParams{
		Method:   "subscribeCandles",
		Location: "/public",
		Params:   params,
	}, handler)

	return
}

// UnsubscribeCandles unsubscribes from candles.
func (h *HitBTC) UnsubscribeCandles(params CandlesParams) (err error) {
	err = h.Stream(exchange.StreamParams{
		Method:   "unsubscribeCandles",
		Location: "/public",
		Params:   params,
	}, func(interface{}) {})

	return
}
