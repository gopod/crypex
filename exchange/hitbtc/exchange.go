package hitbtc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"

	"github.com/ramezanius/crypex/exchange"
	"github.com/ramezanius/crypex/exchange/binance"
)

const (
	apiURL    = "https://api.hitbtc.com/api/2"
	streamURL = "wss://api.hitbtc.com/api/2/ws"
)

// New returns a new hitbtc.
func New() *HitBTC {
	return &HitBTC{
		connections: make(map[string]*websocket.Conn),
	}
}

// Shutdown closes the underlying network connections.
func (h *HitBTC) Shutdown() error {
	for _, conn := range h.connections {
		err := exchange.CloseConn(conn)
		if err != nil {
			return err
		}
	}

	h.connections = make(map[string]*websocket.Conn)

	return nil
}

// Authenticate authenticates a pair of public and secret key
// with the websocket connection.
func (h *HitBTC) Authenticate(conn *websocket.Conn) (err error) {
	if h.PublicKey == "" || h.SecretKey == "" {
		return ErrKeysNotSet
	}

	params := struct {
		PublicKey string `json:"pKey,required"`
		SecretKey string `json:"sKey,required"`
		Algorithm string `json:"algo,required"`
	}{
		PublicKey: h.PublicKey,
		SecretKey: h.SecretKey,
		Algorithm: "BASIC",
	}

	err = conn.WriteJSON(exchange.StreamParams{
		Method: "login",
		Params: params,
	})

	return
}

// Request sends an HTTP request and returns an HTTP response.
func (h *HitBTC) Request(request exchange.RequestParams, response interface{}) error {
	if request.Auth && h.PublicKey == "" || h.SecretKey == "" {
		return binance.ErrKeysNotSet
	}

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

	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)

	defer func() {
		err = res.Body.Close()
		if err != nil {
			return
		}
	}()

	if res.StatusCode != 200 {
		e := &APIError{}
		_ = json.NewDecoder(res.Body).Decode(&e)

		return e
	}

	if response != nil {
		err = json.NewDecoder(res.Body).Decode(&response)
	}

	return err
}

// Stream returns a new connection and writes a request
// to the connection if there's a method.
func (h *HitBTC) Stream(request exchange.StreamParams, handler exchange.HandlerFunc) error {
	conn, err := exchange.NewConn(streamURL, request.Endpoint, h.read, handler)
	if err != nil {
		return err
	}

	if request.Auth {
		err := h.Authenticate(conn)
		if err != nil {
			return err
		}
	}

	if request.Method != "" {
		err = conn.WriteJSON(request)
		if err != nil {
			return err
		}
	}

	h.connections[request.Location] = conn

	return nil
}
