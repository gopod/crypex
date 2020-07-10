package exchange

const (
	// Exchange feeds
	ReportsFeed = "REPORTS"
	CandlesFeed = "CANDLES::"
)

// RequestParams api request struct
type RequestParams struct {
	Auth     bool
	Method   string
	Endpoint string
	Params   interface{}
}

// StreamParams stream request struct
type StreamParams struct {
	Auth     bool   `json:"-"`
	Endpoint string `json:"-"`

	Method string      `json:"method"`
	Params interface{} `json:"params"`
}
