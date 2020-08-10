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
		err := json.Unmarshal(event.Params.([]byte), &response)
		if err != nil {
			log.Fatalf("unmarshal response: [binance]: %v", err)
		}

		go handler(response)
	}

	switch event.Method {
	case "executionReport":
		redirect(&ReportsStream{})

	case "klines":
		go handler(event.Params.(*CandlesStream))

	case "kline":
		redirect(&CandlesStream{})

	case "error":
		redirect(&APIError{})
	}
}

// SubscribeReports subscribes to the reports.
func (b *Binance) SubscribeReports(handler exchange.HandlerFunc) (err error) {
	err = b.Stream(exchange.StreamParams{
		Auth:     true,
		Endpoint: b.ListenKey,
		Location: exchange.TradingLoc,
	}, handler)

	return
}

// UnsubscribeCandles unsubscribes from reports.
func (b *Binance) UnsubscribeReports() (err error) {
	b.Lock()
	defer b.Unlock()

	err = exchange.CloseConn(b.connections[exchange.TradingLoc])
	b.connections[exchange.TradingLoc] = nil

	return
}

// SubscribeCandles subscribes to the candles.
func (b *Binance) SubscribeCandles(params CandlesParams, handler exchange.HandlerFunc) (err error) {
	if params.Limit <= 0 {
		params.Limit = 100
	}

	if params.Snapshot {
		snapshot, err := b.GetCandles(params)
		if err != nil {
			return err
		}

		b.read(&exchange.Event{
			Method: "klines",
			Params: &CandlesStream{
				Period:  params.Period,
				Symbol:  strings.ToUpper(params.Symbol),
				Candles: *snapshot,
			},
		}, handler)
	}

	endpoint := fmt.Sprintf("%s@kline_%s", strings.ToLower(params.Symbol), params.Period)

	err = b.Stream(exchange.StreamParams{
		Endpoint: endpoint,
		Method:   "SUBSCRIBE",
		Params:   []string{endpoint},
		Location: exchange.MarketLoc,
	}, handler)

	return
}

// UnsubscribeCandles unsubscribes from candles.
func (b *Binance) UnsubscribeCandles(params CandlesParams) (err error) {
	endpoint := fmt.Sprintf("%s@kline_%s", strings.ToLower(params.Symbol), params.Period)

	err = b.Stream(exchange.StreamParams{
		Method:   "UNSUBSCRIBE",
		Params:   []string{endpoint},
		Location: exchange.MarketLoc,
	}, func(interface{}) {})

	return
}
