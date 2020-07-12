package hitbtc

import (
	"time"

	"github.com/ramezanius/crypex/exchange"
	"github.com/ramezanius/crypex/exchange/util"
)

// GetSymbol get a specific of exchange symbols.
func (h *HitBTC) GetSymbol(symbol string) (response *Symbol, err error) {
	params := struct {
		Symbol string `json:"symbol,required"`
	}{Symbol: symbol}

	err = h.Stream(exchange.StreamParams{
		Params: params,
		Method: "getSymbol",
	}, &response)
	if err != nil {
		return
	}

	return
}

// GetSymbols gets a list of exchange symbols.
func (h *HitBTC) GetSymbols() (response *Symbols, err error) {
	err = h.Stream(exchange.StreamParams{
		Method: "getSymbols",
	}, &response)
	if err != nil {
		return
	}

	return
}

// GetBalances gets user balances on exchange. @authenticate
func (h *HitBTC) GetBalances() (response *Assets, err error) {
	err = h.Stream(exchange.StreamParams{
		Auth:   true,
		Method: "getTradingBalance",
	}, &response)
	if err != nil {
		return
	}

	return
}

// NewOrder struct
type NewOrder struct {
	Side  Side    `json:"side,required"`
	Type  Type    `json:"type,required"`
	Price float64 `json:"price,string"`

	Symbol   string  `json:"symbol,required"`
	Quantity float64 `json:"quantity,string"`
	OrderID  string  `json:"clientOrderId,required"`

	StopPrice  float64   `json:"stopPrice,required"`
	ExpireTime time.Time `json:"expireTime,required"`

	PostOnly       bool   `json:"postOnly,required"`
	TimeInForce    string `json:"timeInForce,required"`
	StrictValidate bool   `json:"strictValidate,required"`
}

// NewOrder places a new order. @authenticate
func (h *HitBTC) NewOrder(params NewOrder) (response *Report, err error) {
	if params.OrderID == "" {
		params.OrderID = util.GenerateUUID()
	}

	err = h.Stream(exchange.StreamParams{
		Auth:   true,
		Params: params,
		Method: "newOrder",
	}, &response)
	if err != nil {
		return
	}

	return
}

// CancelOrder cancels an order. @authenticate
func (h *HitBTC) CancelOrder(orderID string) (response *Report, err error) {
	params := struct {
		OrderID string `json:"clientOrderId,required"`
	}{OrderID: orderID}

	err = h.Stream(exchange.StreamParams{
		Auth:   true,
		Params: params,
		Method: "cancelOrder",
	}, &response)
	if err != nil {
		return
	}

	return
}

// ReplaceOrder struct
type ReplaceOrder struct {
	Price    float64 `json:"price,string"`
	Quantity float64 `json:"quantity,string"`

	OrderID        string `json:"clientOrderId,required"`
	RequestOrderID string `json:"requestClientId,required"`
	StrictValidate bool   `json:"strictValidate,required"`
}

// ReplaceOrder replaces a new order. @authenticate
func (h *HitBTC) ReplaceOrder(params ReplaceOrder) (response *Report, err error) {
	if params.RequestOrderID == "" {
		params.RequestOrderID = util.GenerateUUID()
	}

	err = h.Stream(exchange.StreamParams{
		Auth:   true,
		Params: params,
		Method: "cancelReplaceOrder",
	}, &response)
	if err != nil {
		return
	}

	return
}
