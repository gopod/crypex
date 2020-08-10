package hitbtc

import "fmt"

type APIError struct {
	Err struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		Description string `json:"description"`
	} `json:"error"`
}

func (e *APIError) Error() string {
	return e.Err.Message
}

var (
	// ErrCurrencyNotFound error
	ErrCurrencyNotFound = fmt.Errorf("currency not found")
	// ErrSymbolNotFound error
	ErrSymbolNotFound = fmt.Errorf("symbol not found")
	// ErrHandlerNotSet error
	ErrHandlerNotSet = fmt.Errorf("handlers not set")
	// ErrKeysNotSet error
	ErrKeysNotSet = fmt.Errorf("hitbtc api keys not set")
)
