package binance

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ramezanius/crypex/exchange"
)

// read redirects response to handler.
func (b *Binance) read(event *exchange.Event, handler exchange.HandlerFunc) {
	var redirect = func(response interface{}) {
		err := json.Unmarshal(event.Params, &response)
		if err != nil {
			log.Fatalf("unmarshal response: [binance]: %v", err)
		}

		go handler(response)
	}

	switch event.Method {
	case "executionReport":
		redirect(&ReportsResponse{})

	case "kline":
		redirect(&CandlesResponse{})
	}
}

// ReportsResponse response
type ReportsResponse Report

// SubscribeReports subscribes to the reports.
func (b *Binance) SubscribeReports(handler exchange.HandlerFunc) (err error) {
	err = b.Stream(exchange.StreamParams{Auth: true, Endpoint: b.ListenKey}, handler)

	return
}

// UnsubscribeCandles unsubscribes from reports.
func (b *Binance) UnsubscribeReports() (err error) {
	b.Lock()
	defer b.Unlock()

	err = exchange.CloseConn(b.connections[b.ListenKey])
	b.connections[b.ListenKey] = nil

	return
}

// CandlesParams struct
type CandlesParams struct {
	Symbol string
	Period Period
}

// CandlesResponse struct
type CandlesResponse struct {
	Period Period `json:"-"`
	Symbol string `json:"s,required"`
	Candle Candle `json:"k,required"`
}

// SubscribeCandles subscribes to the candles.
func (b *Binance) SubscribeCandles(params CandlesParams, handler exchange.HandlerFunc) (err error) {
	err = b.Stream(exchange.StreamParams{
		Endpoint: fmt.Sprintf("%s@kline_%s", strings.ToLower(params.Symbol), params.Period),
	}, handler)

	return
}

// UnsubscribeCandles unsubscribes from candles.
func (b *Binance) UnsubscribeCandles(params CandlesParams) (err error) {
	location := fmt.Sprintf("%s@kline_%s", strings.ToLower(params.Symbol), params.Period)

	b.Lock()
	defer b.Unlock()

	err = exchange.CloseConn(b.connections[location])
	b.connections[location] = nil

	return
}
