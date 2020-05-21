package hitbtc

import (
	"context"
	"github.com/gorilla/websocket"
	JsonRPC2 "github.com/sourcegraph/jsonrpc2"
	JsonRPC2WS "github.com/sourcegraph/jsonrpc2/websocket"
)

const Endpoint = "wss://api.hitbtc.com/api/2/ws"

type Feeds struct {
	ErrorFeed chan error
}

func (r *Feeds) Handle(_ context.Context, _ *JsonRPC2.Conn, _ *JsonRPC2.Request) {
	panic("implement me...")
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

type AuthenticateRequest struct {
	PublicKey string `json:"pKey,required"`
	SecretKey string `json:"sKey,required"`
	Algorithm string `json:"algo,required"`
}

func (h *HitBTC) Authenticate(public, secret string) error {
	err := h.Request("login",
		&AuthenticateRequest{
			PublicKey: public,
			SecretKey: secret,
			Algorithm: "BASIC",
		}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (h *HitBTC) Close() {
	_ = h.Conn.Close()

	close(h.Feeds.ErrorFeed)

	h.Feeds.ErrorFeed = make(chan error)
}

type Response bool

func (h *HitBTC) Subscription(method string, request interface{}) error {
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
