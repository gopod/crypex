package exchange

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Event struct
type Event struct {
	Method string
	Params []byte
}

type (
	// HandlerFunc func
	HandlerFunc func(interface{})
	// ReaderFunc func
	ReaderFunc func(*Event, HandlerFunc)
)

// NewConn creates a new websocket client connection.
func NewConn(stream, endpoint string, reader ReaderFunc, handler HandlerFunc) (conn *websocket.Conn, err error) {
	var url string

	if endpoint == "" {
		url = stream
	} else {
		url = fmt.Sprintf("%s/%s", stream, endpoint)
	}

	conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return
	}

	go readMessages(conn, reader, handler)

	return
}

// CloseConn closes websocket connection.
func CloseConn(conn *websocket.Conn) (err error) {
	if conn == nil {
		return
	}

	err = conn.Close()
	if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		return err
	}

	return nil
}

// readMessages reads the next JSON-encoded message from the connection
// and send it to the reader func with handler.
func readMessages(conn *websocket.Conn, reader ReaderFunc, handler HandlerFunc) {
	for {
		var (
			event   = &Event{}
			payload = map[string]json.RawMessage{}
		)

		_, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Fatalf("read message: %v", err)
			}

			break
		}

		err = json.Unmarshal(msg, &payload)
		if err != nil {
			log.Fatalf("unmarshal payload: %v", err)
		}

		// Serialize jsonrpc2 response message
		if method, ok := payload["method"]; ok {
			event.Method = string(method)
			event.Method = event.Method[1 : len(event.Method)-1]

			event.Params, err = json.Marshal(payload["params"])
			if err != nil {
				log.Fatalf("marshal jsonrpc2 params: %v", err)
			}
		}

		// Serialize non-jsonrpc2 response message
		if method, ok := payload["e"]; ok {
			event.Method = string(method)
			event.Method = event.Method[1 : len(event.Method)-1]

			delete(payload, "e")
			delete(payload, "E")

			event.Params, err = json.Marshal(&payload)
			if err != nil {
				log.Fatalf("marshal non-jsonrpc2 params: %v", err)
			}
		}

		go reader(event, handler)
	}

	err := CloseConn(conn)
	if err != nil {
		log.Fatalf("close websocket connection: %v", err)
	}
}
