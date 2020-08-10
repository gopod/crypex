package binance

import "fmt"

// APIError struct
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e *APIError) Error() string {
	return e.Message
}

var (
	// ErrCurrencyNotFound error
	ErrCurrencyNotFound = fmt.Errorf("currency not found")
	// ErrSymbolNotFound error
	ErrSymbolNotFound = fmt.Errorf("symbol not found")
	// ErrHandlerNotSet error
	ErrHandlerNotSet = fmt.Errorf("handlers not set")
	// ErrKeysNotSet error
	ErrKeysNotSet = fmt.Errorf("binance api keys are not set")
)
