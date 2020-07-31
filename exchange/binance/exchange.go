package binance

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/ratelimit"

	"github.com/ramezanius/crypex/exchange"
	"github.com/ramezanius/crypex/exchange/util"
)

const (
	recvWindow = "5000"
	keepAlive  = 30 * time.Minute

	apiURL    = "https://api.binance.com/api/v3"
	streamURL = "wss://stream.binance.com:9443/ws"
)

// New returns a new binance.
func New() *Binance {
	return &Binance{
		wsLimit:      ratelimit.New(1),
		publicLimit:  ratelimit.New(20),
		tradingLimit: ratelimit.New(10),
		connections:  make(map[string]*websocket.Conn),
	}
}

// Shutdown closes the underlying network connections.
func (b *Binance) Shutdown() error {
	for _, conn := range b.connections {
		err := exchange.CloseConn(conn)
		if err != nil {
			return err
		}
	}

	b.connections = make(map[string]*websocket.Conn)

	return nil
}

// Authenticate creates an listen key and keeps it alive.
func (b *Binance) Authenticate() (err error) {
	if b.PublicKey == "" || b.SecretKey == "" {
		return ErrKeysNotSet
	}

	var response map[string]string

	// Send a POST request to create a listen key.
	err = b.Request(exchange.RequestParams{
		Method: "POST", Endpoint: "/userDataStream",
	}, &response)
	if err != nil {
		return
	}

	b.ListenKey = response["listenKey"]

	// Send a PUT request every 30 minutes for keep-alive listen key.
	go func() {
		ticker := time.NewTicker(keepAlive)

		for range ticker.C {
			b.Lock()

			_ = b.Request(exchange.RequestParams{
				Method: "PUT", Endpoint: "/userDataStream",
				Params: map[string]string{
					"listenKey": b.ListenKey,
				},
			}, nil)

			b.Unlock()
		}
	}()

	return nil
}

// Request sends an HTTP request and returns an HTTP response.
func (b *Binance) Request(request exchange.RequestParams, response interface{}) error {
	if request.Auth {
		b.tradingLimit.Take()
	} else {
		b.publicLimit.Take()
	}

	parsedURL, _ := url.ParseRequestURI(apiURL)
	parsedURL.Path = parsedURL.Path + request.Endpoint

	// Parse params to query string
	bin, err := json.Marshal(request.Params)
	if err != nil {
		return err
	}
	m := map[string]interface{}{}

	err = json.Unmarshal(bin, &m)
	if err != nil {
		return err
	}

	q := parsedURL.Query()
	for k, v := range m {
		q.Set(k, fmt.Sprintf("%v", v))
	}

	if request.Auth {
		if b.PublicKey == "" || b.SecretKey == "" {
			return ErrKeysNotSet
		}

		q.Add("recvWindow", recvWindow)
		// Timestamp is mandatory in signed request
		q.Add("timestamp", fmt.Sprintf("%v", time.Now().Unix()*1000))
		// Signature needs to be at the last param
		parsedURL.RawQuery =
			q.Encode() + "&signature=" + util.GenerateSignature(b.SecretKey, q)
	} else {
		parsedURL.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(request.Method, parsedURL.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-MBX-APIKEY", b.PublicKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			return
		}
	}()

	if res.StatusCode != 200 {
		e := &APIError{}
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return err
		}

		return e
	}

	if response != nil {
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return err
		}
	}

	return nil
}

// Stream returns a new connection with a specific endpoint.
func (b *Binance) Stream(request exchange.StreamParams, handler exchange.HandlerFunc) error {
	b.Lock()
	defer b.Unlock()
	b.wsLimit.Take()

	var (
		err  error
		conn *websocket.Conn
	)

	if request.Auth {
		err := b.Authenticate()
		if err != nil {
			return err
		}

		request.Endpoint = b.ListenKey
	}

	if _, ok := b.connections[request.Location]; !ok {
		conn, err = exchange.NewConn(streamURL, request.Endpoint, b.read, handler)
		if err != nil {
			return err
		}

		b.connections[request.Location] = conn

		return err
	}

	conn = b.connections[request.Location]

	if request.Method != "" {
		if request.ID != 0 {
			request.ID = rand.Int()
		}

		err = conn.WriteJSON(request)
		if err != nil {
			return err
		}
	}

	return nil
}
