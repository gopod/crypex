package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ramezanius/jsonrpc2"

	"github.com/ramezanius/crypex/exchange"
	"github.com/ramezanius/crypex/exchange/util"
)

const (
	subscribe   = "SUBSCRIBE"
	unsubscribe = "UN" + subscribe
)

// Handle handles stream feeds.
func (f *Feeds) Handle(_ context.Context, _ *jsonrpc2.Conn, request *jsonrpc2.Request) {
	response := *request.Params

	switch request.Method {
	case "executionReport":
		var msg ReportsResponse

		err := json.Unmarshal(response, &msg)
		if err != nil {
			util.UnmarshalError(err, Exchange)
		}

		f.Pub(msg, exchange.ReportsFeed)

	case "kline":
		var msg CandlesResponse

		err := json.Unmarshal(response, &msg)
		if err != nil {
			util.UnmarshalError(err, Exchange)
		}

		f.Pub(msg, exchange.CandlesFeed+strings.ToLower(msg.Symbol))
	}
}

// ReportsResponse candles stream response
type ReportsResponse Report

// SubscribeReports Subscribe to all reports. @authenticate
func (b *Binance) SubscribeReports() (reports <-chan interface{}, err error) {
	reports = b.Feeds.Sub(exchange.ReportsFeed)

	return
}

// CandlesParams candles stream params
type CandlesParams struct {
	Symbol string
	Period Period
}

// CandlesResponse candles stream response
type CandlesResponse struct {
	Symbol string `json:"s,required"`
	Candle Candle `json:"k,required"`
}

// SubscribeCandles subscribe to symbol candles.
func (b *Binance) SubscribeCandles(params CandlesParams) (candles <-chan interface{}, err error) {
	method := []string{
		fmt.Sprintf("%s@kline_%s", strings.ToLower(params.Symbol), params.Period),
	}

	err = b.Stream(exchange.StreamParams{
		Method: subscribe,
		Params: method,
	}, nil)

	candles = b.Feeds.Sub(exchange.CandlesFeed + params.Symbol)

	return
}

// UnsubscribeCandles unsubscribe from symbol candles.
func (b *Binance) UnsubscribeCandles(params CandlesParams) (err error) {
	b.Feeds.Close(exchange.CandlesFeed + params.Symbol)

	method := []string{
		fmt.Sprintf("%s@kline", strings.ToLower(params.Symbol)),
	}
	err = b.Stream(exchange.StreamParams{
		Method: unsubscribe,
		Params: method,
	}, nil)

	return
}
