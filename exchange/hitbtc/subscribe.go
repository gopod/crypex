package hitbtc

import (
	"context"
	"encoding/json"

	"github.com/sourcegraph/jsonrpc2"

	"github.com/ramezanius/crypex/exchange"
	"github.com/ramezanius/crypex/exchange/util"
)

// Handle handles stream feeds.
func (f *Feeds) Handle(_ context.Context, _ *jsonrpc2.Conn, request *jsonrpc2.Request) {
	response := *request.Params

	switch request.Method {
	case "report":
		var msg ReportResponse

		err := json.Unmarshal(response, &msg)
		if err != nil {
			util.UnmarshalError(err, Exchange)
		}

		f.Pub(msg, exchange.ReportsFeed)

	case "snapshotCandles", "updateCandles":
		var msg CandlesResponse

		err := json.Unmarshal(response, &msg)
		if err != nil {
			util.UnmarshalError(err, Exchange)
		}

		f.Pub(msg, exchange.CandlesFeed+msg.Symbol)
	}
}

// ReportResponse struct
type ReportResponse Report

// SubscribeReports Subscribe to all reports. @authenticate
func (h *HitBTC) SubscribeReports() (reports <-chan interface{}, err error) {
	err = h.Stream(exchange.StreamParams{
		Auth:   true,
		Method: "subscribeReports",
	}, nil)
	if err != nil {
		return
	}

	reports = h.Feeds.Sub(exchange.ReportsFeed)

	return
}

// CandlesResponse struct
type CandlesResponse struct {
	Data   Candles `json:"data,required"`
	Symbol string  `json:"symbol,required"`
	Period string  `json:"period,required"`
}

// CandlesResponse struct
type CandlesParams struct {
	Limit  int64  `json:"limit,omitempty"`
	Period Period `json:"period,omitempty"`
	Symbol string `json:"symbol,required"`
}

// SubscribeCandles subscribe to symbol candles.
func (h *HitBTC) SubscribeCandles(params CandlesParams) (candles <-chan interface{}, err error) {
	err = h.Stream(exchange.StreamParams{
		Method: "subscribeCandles",
		Params: params,
	}, nil)
	if err != nil {
		return
	}

	candles = h.Feeds.Sub(exchange.CandlesFeed + params.Symbol)

	return
}

// UnsubscribeCandles unsubscribe from symbol candles.
func (h *HitBTC) UnsubscribeCandles(params CandlesParams) (err error) {
	h.Feeds.Close(exchange.CandlesFeed + params.Symbol)

	err = h.Stream(exchange.StreamParams{
		Method: "unsubscribeCandles",
		Params: params,
	}, nil)

	return err
}
