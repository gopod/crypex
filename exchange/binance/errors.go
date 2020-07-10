package binance

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e APIError) String() string {
	return e.Message
}
