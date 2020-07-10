package hitbtc

type APIError struct {
	Error struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		Description string `json:"description"`
	} `json:"error"`
}

func (e APIError) String() string {
	return e.Error.Message
}
