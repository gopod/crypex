package exchange

// RequestParams struct
type RequestParams struct {
	Auth     bool
	Endpoint string
	Method   string
	Params   interface{}
}

// StreamParams struct
type StreamParams struct {
	Auth     bool        `json:"-"`
	Location string      `json:"-"`
	Endpoint string      `json:"-"`
	Method   string      `json:"method"`
	Params   interface{} `json:"params"`
}
