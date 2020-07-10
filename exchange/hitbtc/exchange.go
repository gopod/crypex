package hitbtc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cskr/pubsub"
	"github.com/gorilla/websocket"
	"github.com/ramezanius/jsonrpc2"
	ws "github.com/ramezanius/jsonrpc2/websocket"

	"github.com/ramezanius/crypex/exchange"
)

const (
	apiURL    = "https://api.hitbtc.com/api/2"
	streamURL = "wss://api.hitbtc.com/api/2/ws"
)

// New create a new hitbtc instance.
func New(public, secret string) (*HitBTC, error) {
	feeds := &Feeds{pubsub.New(1)}
	hitbtc := &HitBTC{
		Feeds: feeds,
	}

	conn, _, err := websocket.DefaultDialer.Dial(streamURL, nil)
	if err != nil {
		return nil, err
	}

	hitbtc.Connection = jsonrpc2.NewConn(
		context.Background(),
		ws.NewObjectStream(conn),
		jsonrpc2.AsyncHandler(feeds),
	)

	hitbtc.PublicKey, hitbtc.SecretKey = public, secret
	err = hitbtc.Authenticate()
	if err != nil {
		return nil, err
	}

	return hitbtc, nil
}

// authenticate sends a request for authorize instance.
func (h *HitBTC) Authenticate() (err error) {
	params := struct {
		PublicKey string `json:"pKey,required"`
		SecretKey string `json:"sKey,required"`
		Algorithm string `json:"algo,required"`
	}{
		PublicKey: h.PublicKey,
		SecretKey: h.SecretKey,
		Algorithm: "BASIC",
	}

	err = h.Stream(exchange.StreamParams{
		Method: "login",
		Params: params,
	}, nil)

	return
}

// Shutdown removes and closes subscribed channels.
func (h *HitBTC) Shutdown() error {
	h.Feeds.Shutdown()

	return h.Connection.Close()
}

// request sends an http request.
func (h *HitBTC) Request(request exchange.RequestParams, response interface{}) error {
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

	parsedURL.RawQuery = q.Encode()
	req, _ := http.NewRequest(request.Method, parsedURL.String(), nil)

	if request.Auth && h.PublicKey != "" && h.SecretKey != "" {
		req.SetBasicAuth(h.PublicKey, h.SecretKey)
	}

	req.Header.Add("Content-type", "application/json")
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
func (h *HitBTC) Stream(request exchange.StreamParams, response interface{}) error {
	return h.Connection.Call(
		context.Background(), request.Method, request.Params, &response,
	)
}
