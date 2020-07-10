package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	case "kline":
		var msg CandlesResponse

		err := json.Unmarshal(response, &msg)
		if err != nil {
			util.UnmarshalError(err, Exchange)
		}

		f.Pub(msg, exchange.CandlesFeed+strings.ToLower(msg.Symbol))
	}
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

func (r *CandlesResponse) UnmarshalJSON(data []byte) error {
	var err error

	var v struct {
		Symbol string `json:"s,required"`
		Candle struct {
			StartAt     int64       `json:"t,required"`
			EndAt       int64       `json:"T,required"`
			Max         interface{} `json:"h,required"`
			Min         interface{} `json:"l,required"`
			Open        interface{} `json:"o,required"`
			Close       interface{} `json:"c,required"`
			Closed      bool        `json:"x,required"`
			Symbol      string      `json:"s,required"`
			Period      string      `json:"i,required"`
			Volume      interface{} `json:"v,required"`
			QuoteVolume interface{} `json:"q,required"`
		} `json:"k,required"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.Symbol = v.Symbol
	r.Candle.Closed = v.Candle.Closed
	r.Candle.Symbol = v.Candle.Symbol
	r.Candle.Period = v.Candle.Period

	startTime := time.Unix(v.Candle.StartAt, 0)
	r.Candle.StartAt = &startTime

	endTime := time.Unix(v.Candle.StartAt, 0)
	r.Candle.EndAt = &endTime

	r.Candle.Min, err = strconv.ParseFloat(v.Candle.Min.(string), 64)
	if err != nil {
		return err
	}

	r.Candle.Max, err = strconv.ParseFloat(v.Candle.Max.(string), 64)
	if err != nil {
		return err
	}

	r.Candle.Open, err = strconv.ParseFloat(v.Candle.Open.(string), 64)
	if err != nil {
		return err
	}

	r.Candle.Close, err = strconv.ParseFloat(v.Candle.Close.(string), 64)
	if err != nil {
		return err
	}

	r.Candle.Volume, err = strconv.ParseFloat(v.Candle.Volume.(string), 64)
	if err != nil {
		return err
	}

	r.Candle.QuoteVolume, err = strconv.ParseFloat(v.Candle.QuoteVolume.(string), 64)
	if err != nil {
		return err
	}

	return nil
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
	method := []string{
		fmt.Sprintf("%s@kline", strings.ToLower(params.Symbol)),
	}
	err = b.Stream(exchange.StreamParams{
		Method: unsubscribe,
		Params: method,
	}, nil)

	return
}
