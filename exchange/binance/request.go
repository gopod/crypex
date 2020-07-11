package binance

import (
	"strings"

	"github.com/ramezanius/crypex/exchange"
)

// SymbolsResponse struct
type SymbolsResponse struct {
	Data Symbols `json:"symbols,required"`
}

// GetSymbols gets a list of exchange symbols.
func (b *Binance) GetSymbols() (response *Symbols, err error) {
	var rawResponse SymbolsResponse

	err = b.Request(exchange.RequestParams{
		Method: "GET", Endpoint: "/exchangeInfo",
	}, &rawResponse)
	if err != nil {
		return
	}

	response = &rawResponse.Data

	return
}

// BalancesResponse struct
type BalancesResponse struct {
	Data Assets `json:"balances,required"`
}

// GetBalances gets user balances on exchange. @authenticate
func (b *Binance) GetBalances() (response *Assets, err error) {
	var rawResponse BalancesResponse

	err = b.Request(exchange.RequestParams{
		Method: "GET", Endpoint: "/account", Auth: true,
	}, &rawResponse)
	if err != nil {
		return
	}

	response = &rawResponse.Data

	return
}

// OrderResponse struct
type OrderResponse Order

// NewOrder struct
type NewOrder struct {
	Side  Side    `json:"side,required"`
	Type  Type    `json:"type,required"`
	Price float64 `json:"price,string"`

	Symbol   string  `json:"symbol,required"`
	Quantity float64 `json:"quantity,string"`
	OrderID  string  `json:"newClientOrderId,omitempty"`

	StopPrice   float64     `json:"stopPrice,omitempty"`
	TimeInForce TimeInForce `json:"timeInForce,omitempty"`
}

// NewOrder places a new order. @authenticate
func (b *Binance) NewOrder(params NewOrder) (response *Report, err error) {
	params.Symbol = strings.ToUpper(params.Symbol)

	var rawResponse OrderResponse

	err = b.Request(exchange.RequestParams{
		Auth: true, Params: params,
		Method: "POST", Endpoint: "/order",
	}, &rawResponse)
	if err != nil {
		return
	}

	report := Report(rawResponse)
	response = &report

	return
}

// CancelOrder cancels an order. @authenticate
func (b *Binance) CancelOrder(orderID, symbol string) (response *Report, err error) {
	params := struct {
		Symbol  string `json:"symbol,required"`
		OrderID string `json:"origClientOrderId,required"`
	}{
		OrderID: orderID,
		Symbol:  strings.ToUpper(symbol),
	}

	var rawResponse OrderResponse

	err = b.Request(exchange.RequestParams{
		Auth: true, Params: params,
		Method: "DELETE", Endpoint: "/order",
	}, &rawResponse)
	if err != nil {
		return
	}

	report := Report(rawResponse)
	response = &report

	return
}
