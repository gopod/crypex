package hitbtc

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

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
		redirect(&ReportsResponse{})

	case "updateCandles":
		redirect(&CandlesResponse{})
	}
}

// ReportsResponse struct
type ReportsResponse Report

// SubscribeReports subscribes to the reports.
func (h *HitBTC) SubscribeReports(handler exchange.HandlerFunc) (err error) {
	err = h.Stream(exchange.StreamParams{
		Auth:     true,
		Location: h.SecretKey,
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

// CandlesResponse struct
type CandlesResponse struct {
	Candle Candle `json:"data,required"`
	Symbol string `json:"symbol,required"`
	Period string `json:"period,required"`
}

// CandlesResponse struct
type CandlesParams struct {
	Period Period `json:"period,omitempty"`
	Symbol string `json:"symbol,required"`
}

// SubscribeCandles subscribes to the candles.
func (h *HitBTC) SubscribeCandles(params CandlesParams, handler exchange.HandlerFunc) (err error) {
	err = h.Stream(exchange.StreamParams{
		Location: fmt.Sprintf("%s@candle_%s", strings.ToLower(params.Symbol), params.Period),
		Method:   "subscribeCandles",
		Params:   params,
	}, handler)

	return
}

// UnsubscribeCandles unsubscribes from candles.
func (h *HitBTC) UnsubscribeCandles(params CandlesParams) (err error) {
	location := fmt.Sprintf("%s@candle_%s", strings.ToLower(params.Symbol), params.Period)

	h.Lock()
	defer h.Unlock()

	err = exchange.CloseConn(h.connections[location])
	h.connections[location] = nil

	return err
}
