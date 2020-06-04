package hitbtc

import (
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
	JsonRPC2 "github.com/sourcegraph/jsonrpc2"
	JsonRPC2WS "github.com/sourcegraph/jsonrpc2/websocket"
)

const Endpoint = "wss://api.hitbtc.com/api/2/ws"

type Feeds struct {
	ErrorFeed chan error

	Notifications Notifications

	ReportsFeed   chan ReportsSnapshot
	CandlesFeed   map[string]chan CandlesSnapshot
	OrderbookFeed map[string]chan OrderbookSnapshot
}

type Notifications struct {
	ReportsFeed   chan ReportsUpdate
	CandlesFeed   map[string]chan CandlesUpdate
	OrderbookFeed map[string]chan OrderbookUpdate
}

func (r *Feeds) Handle(_ context.Context, _ *JsonRPC2.Conn, request *JsonRPC2.Request) {
	if request.Params != nil {
		return
	}

	message := *request.Params

	switch request.Method {
	case "activeOrders":
		var msg ReportsSnapshot

		err := json.Unmarshal(message, &msg)
		if err != nil {
			r.ErrorFeed <- err
		} else {
			r.ReportsFeed <- msg
		}
	case "report":
		var msg ReportsUpdate

		err := json.Unmarshal(message, &msg)

		if err != nil {
			r.ErrorFeed <- err
		} else {
			r.Notifications.ReportsFeed <- msg
		}

	case "snapshotCandles":
		var msg CandlesSnapshot

		err := json.Unmarshal(message, &msg)

		if err != nil {
			r.ErrorFeed <- err
		} else {
			r.CandlesFeed[msg.Symbol] <- msg
		}
	case "updateCandles":
		var msg CandlesUpdate

		err := json.Unmarshal(message, &msg)

		if err != nil {
			r.ErrorFeed <- err
		} else {
			r.Notifications.CandlesFeed[msg.Symbol] <- msg
		}

	case "snapshotOrderbook":
		var msg OrderbookSnapshot

		err := json.Unmarshal(message, &msg)

		if err != nil {
			r.ErrorFeed <- err
		} else {
			r.OrderbookFeed[msg.Symbol] <- msg
		}
	case "updateOrderbook":
		var msg OrderbookUpdate

		err := json.Unmarshal(message, &msg)

		if err != nil {
			r.ErrorFeed <- err
		} else {
			r.Notifications.OrderbookFeed[msg.Symbol] <- msg
		}
	}
}

type HitBTC struct {
	Feeds *Feeds

	Conn *JsonRPC2.Conn
}

func New() (instance *HitBTC, err error) {
	connection, _, err := websocket.DefaultDialer.Dial(Endpoint, nil)
	if err != nil {
		return nil, err
	}

	feeds := &Feeds{
		ErrorFeed: make(chan error),

		Notifications: Notifications{
			ReportsFeed:   make(chan ReportsUpdate),
			CandlesFeed:   make(map[string]chan CandlesUpdate),
			OrderbookFeed: make(map[string]chan OrderbookUpdate),
		},

		ReportsFeed:   make(chan ReportsSnapshot),
		CandlesFeed:   make(map[string]chan CandlesSnapshot),
		OrderbookFeed: make(map[string]chan OrderbookSnapshot),
	}

	instance = &HitBTC{
		Conn: JsonRPC2.NewConn(
			context.Background(),
			JsonRPC2WS.NewObjectStream(connection), JsonRPC2.AsyncHandler(feeds),
		),
		Feeds: feeds,
	}

	return
}

func (h *HitBTC) Authenticate(public, secret string) error {
	request := struct {
		PublicKey string `json:"pKey,required"`
		SecretKey string `json:"sKey,required"`
		Algorithm string `json:"algo,required"`
	}{
		PublicKey: public,
		SecretKey: secret,
		Algorithm: "BASIC",
	}

	err := h.Request("login", &request, nil)
	if err != nil {
		return err
	}

	return nil
}

func (h *HitBTC) Close() {
	_ = h.Conn.Close()

	close(h.Feeds.ErrorFeed)
	close(h.Feeds.ReportsFeed)
	close(h.Feeds.Notifications.ReportsFeed)

	for _, channel := range h.Feeds.CandlesFeed {
		close(channel)
	}

	for _, channel := range h.Feeds.Notifications.CandlesFeed {
		close(channel)
	}

	for _, channel := range h.Feeds.OrderbookFeed {
		close(channel)
	}

	for _, channel := range h.Feeds.Notifications.OrderbookFeed {
		close(channel)
	}

	h.Feeds.ErrorFeed = make(chan error)

	h.Feeds.CandlesFeed = make(map[string]chan CandlesSnapshot)
	h.Feeds.Notifications.CandlesFeed = make(map[string]chan CandlesUpdate)

	h.Feeds.OrderbookFeed = make(map[string]chan OrderbookSnapshot)
	h.Feeds.Notifications.OrderbookFeed = make(map[string]chan OrderbookUpdate)
}

type Response bool

func (h *HitBTC) Subscribe(method string, request interface{}) error {
	var response Response

	err := h.Conn.Call(context.Background(), method, &request, &response)
	if err != nil {
		return err
	}

	return nil
}

func (h *HitBTC) Request(method string, request interface{}, response interface{}) error {
	err := h.Conn.Call(context.Background(), method, &request, &response)
	if err != nil {
		return err
	}

	return nil
}
