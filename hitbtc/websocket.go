package hitbtc

import (
	"context"
	"encoding/json"

	"github.com/cskr/pubsub"
	"github.com/gorilla/websocket"
	JsonRPC2 "github.com/sourcegraph/jsonrpc2"
	JsonRPC2WS "github.com/sourcegraph/jsonrpc2/websocket"
)

// Feeds channels struct
type Feeds struct {
	*pubsub.PubSub
}

// Handle websocket response handler
func (r *Feeds) Handle(_ context.Context, _ *JsonRPC2.Conn, request *JsonRPC2.Request) {
	message := *request.Params

	if request.Params == nil {
		return
	}

	switch request.Method {
	case "report":
		var msg ReportResponse

		err := json.Unmarshal(message, &msg)
		if err != nil {
			r.Pub(err, ErrorFeed)
		}

		r.Pub(msg, ReportsFeed)

	case "snapshotCandles", "updateCandles":
		var msg CandlesResponse

		err := json.Unmarshal(message, &msg)
		if err != nil {
			r.Pub(err, ErrorFeed)
		}

		r.Pub(msg, CandlesFeed+msg.Symbol)

	case "snapshotOrderbook", "updateOrderbook":
		var msg OrderbookResponse

		err := json.Unmarshal(message, &msg)
		if err != nil {
			r.Pub(err, ErrorFeed)
		}

		r.Pub(msg, OrderbookFeed+msg.Symbol)
	}
}

// Publish parse and publish data to channel
func (r *Feeds) Publish(origin []byte, msg interface{}, feed string) {
	err := json.Unmarshal(origin, &msg)
	if err != nil {
		r.Pub(err, ErrorFeed)
	}

	r.Pub(msg, feed)
}

// HitBTC struct
type HitBTC struct {
	Feeds *Feeds

	Conn *JsonRPC2.Conn

	PublicKey, SecretKey string
}

// New create a new hitbtc socket connection
func New() (*HitBTC, error) {
	connection, _, err := websocket.DefaultDialer.Dial(Endpoint, nil)
	if err != nil {
		return nil, err
	}

	feeds := &Feeds{pubsub.New(1)}

	return &HitBTC{
		Conn: JsonRPC2.NewConn(
			context.Background(),
			JsonRPC2WS.NewObjectStream(connection), JsonRPC2.AsyncHandler(feeds),
		),
		Feeds: feeds,
	}, nil
}

// Authenticate basic authenticate with public, and secret key
func (h *HitBTC) Authenticate() (err error) {
	request := struct {
		PublicKey string `json:"pKey,required"`
		SecretKey string `json:"sKey,required"`
		Algorithm string `json:"algo,required"`
	}{
		PublicKey: h.PublicKey,
		SecretKey: h.SecretKey,
		Algorithm: "BASIC",
	}

	err = h.Request("login", &request, nil)

	return
}

// Shutdown close and delete all subscribed channels
func (h *HitBTC) Shutdown() (err error) {
	h.Feeds.Shutdown()
	err = h.Conn.Close()

	return
}

// Response type
type Response bool

// Subscribe subscribe to specific method
func (h *HitBTC) Subscribe(method string, request interface{}) (err error) {
	var response Response

	err = h.Conn.Call(context.Background(), method, &request, &response)

	return
}

// Request call a method in websocket connection
func (h *HitBTC) Request(method string, request, response interface{}) (err error) {
	err = h.Conn.Call(context.Background(), method, &request, &response)

	return
}
