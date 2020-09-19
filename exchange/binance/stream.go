package binance

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gopod/crypex/exchange"
)

// read redirects response to handler.
func (b *Binance) read(event *exchange.Event) {
	switch event.Method {
	case "executionReport":
		object := &ReportsStream{}

		err := json.Unmarshal(event.Params.([]byte), &object)
		if err != nil {
			b.OnErr(err)
		}

		b.Feeds.Reports <- object

	case "klines":
		b.Feeds.Candles <- event.Params.(*CandlesStream)

	case "kline":
		object := &CandlesStream{}

		err := json.Unmarshal(event.Params.([]byte), &object)
		if err != nil {
			b.OnErr(err)
		}

		b.Feeds.Candles <- object

	case "error":
		object := &APIError{}

		err := json.Unmarshal(event.Params.([]byte), &object)
		if err != nil {
			b.OnErr(err)
		}

		b.OnErr(object)
	}
}

// SubscribeReports subscribes to the reports.
func (b *Binance) SubscribeReports() (err error) {
	err = b.Stream(exchange.StreamParams{
		Auth:     true,
		Endpoint: b.ListenKey,
		Location: exchange.TradingLoc,
	})

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
func (b *Binance) SubscribeCandles(params CandlesParams) (err error) {
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
				Candles: Candles(*snapshot),
			},
		})
	}

	endpoint := fmt.Sprintf("%s@kline_%s", strings.ToLower(params.Symbol), params.Period)

	err = b.Stream(exchange.StreamParams{
		Endpoint: endpoint,
		Method:   "SUBSCRIBE",
		Params:   []string{endpoint},
		Location: exchange.MarketLoc,
	})

	return
}

// UnsubscribeCandles unsubscribes from candles.
func (b *Binance) UnsubscribeCandles(params CandlesParams) (err error) {
	endpoint := fmt.Sprintf("%s@kline_%s", strings.ToLower(params.Symbol), params.Period)

	err = b.Stream(exchange.StreamParams{
		Method:   "UNSUBSCRIBE",
		Params:   []string{endpoint},
		Location: exchange.MarketLoc,
	})

	return
}
