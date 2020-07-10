package binance

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cskr/pubsub"
	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	ws "github.com/sourcegraph/jsonrpc2/websocket"

	"github.com/ramezanius/crypex/exchange"
)

const (
	apiURL    = "https://api.binance.com/api/v3"
	streamURL = "wss://stream.binance.com:9443/ws/"

	keepAlive = 30 * time.Minute
)

// New create a new binance instance.
func New(public, secret string) (*Binance, error) {
	feeds := &Feeds{pubsub.New(1)}
	binance := &Binance{
		Feeds: feeds,
	}

	binance.PublicKey, binance.SecretKey = public, secret
	err := binance.Authenticate()
	if err != nil {
		return nil, err
	}

	conn, _, err := websocket.DefaultDialer.Dial(streamURL+binance.ListenKey, nil)
	if err != nil {
		return nil, err
	}

	binance.Connection = jsonrpc2.NewConn(
		context.Background(),
		ws.NewObjectStream(conn),
		jsonrpc2.AsyncHandler(feeds),
	)

	// Serialize event to JSON-RPC message
	jsonrpc2.MessageSerializer = func(data []byte) ([]byte, error) {
		var payload map[string]json.RawMessage
		if err := json.Unmarshal(data, &payload); err != nil {
			return nil, err
		}

		var (
			ok     bool
			method json.RawMessage
		)

		if method, ok = payload["e"]; !ok {
			return data, nil
		}

		var msg struct {
			Result interface{} `json:"params"`

			Method json.RawMessage `json:"method"`
		}

		delete(payload, "e")
		delete(payload, "E")
		msg.Method, msg.Result = method, payload

		output, err := json.Marshal(msg)
		if err != nil {
			return nil, err
		}

		return output, nil
	}

	return binance, nil
}

// Shutdown removes and closes subscribed channels.
func (b *Binance) Shutdown() error {
	b.Feeds.Shutdown()

	return b.Connection.Close()
}

// Authenticate sends a request for authorize instance.
func (b *Binance) Authenticate() (err error) {
	var response map[string]string

	// Send a POST request to create a listen key.
	err = b.Request(exchange.RequestParams{
		Method: "POST", Endpoint: "/userDataStream",
	}, &response)
	if err != nil {
		return
	}

	b.ListenKey = response["listenKey"]
	b.Signature = hex.EncodeToString(
		hmac.New(sha256.New, []byte(b.SecretKey)).Sum(nil))

	// Send a PUT request every 30 minutes for keep-alive listen key.
	go func() {
		ticker := time.NewTicker(keepAlive)

		for range ticker.C {
			b.Lock()
			if b.ListenKey == "" {
				continue
			}

			err = b.Request(exchange.RequestParams{
				Method: "PUT", Endpoint: "/userDataStream",
				Params: map[string]string{
					"listenKey": b.ListenKey,
				},
			}, nil)
			if err != nil {
				return
			}

			b.Unlock()
		}
	}()

	return err
}

// request sends an http request.
func (b *Binance) Request(request exchange.RequestParams, response interface{}) error {
	parsedURL, _ := url.ParseRequestURI(apiURL)
	parsedURL.Path = parsedURL.Path + request.Endpoint

	// Parse params to query string
	bin, _ := json.Marshal(request.Params)
	m := map[string]interface{}{}

	err := json.Unmarshal(bin, &m)
	if err != nil {
		return err
	}

	q := parsedURL.Query()
	for k, v := range m {
		q.Set(k, fmt.Sprintf("%v", v))
	}

	if request.Auth && b.Signature != "" {
		// Signature needs to be at the last param
		parsedURL.RawQuery = q.Encode() + "&signature=" + b.Signature
	}

	parsedURL.RawQuery = q.Encode()
	req, _ := http.NewRequest(request.Method, parsedURL.String(), nil)

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-MBX-APIKEY", b.PublicKey)
	res, err := http.DefaultClient.Do(req)

	defer func() {
		err = res.Body.Close()
		if err != nil {
			return
		}
	}()

	if res.StatusCode != 200 {
		e := APIError{}

		err = json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return err
		}

		return fmt.Errorf(e.String())
	}

	if response != nil {
		err = json.NewDecoder(res.Body).Decode(&response)
	}

	return err
}

// Stream sends a request to websocket connection.
func (b *Binance) Stream(request exchange.StreamParams, response interface{}) error {
	if request.Auth {
		err := b.Authenticate()
		if err != nil {
			return err
		}
	}

	return b.Connection.Call(
		context.Background(), request.Method, request.Params, &response,
	)
}
