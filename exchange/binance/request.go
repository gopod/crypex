package binance

import (
	"strings"

	"github.com/ramezanius/crypex/exchange"
	"github.com/ramezanius/crypex/exchange/util"
)

// SymbolsResponse struct
type SymbolsResponse struct {
	Symbols Symbols `json:"symbols,required"`
}

// GetSymbols returns exchange symbols.
func (b *Binance) GetSymbols() (response *Symbols, err error) {
	var rawResponse SymbolsResponse

	err = b.Request(exchange.RequestParams{
		Method: "GET", Endpoint: "/exchangeInfo",
	}, &rawResponse)
	response = &rawResponse.Symbols

	return
}

// AssetsResponse struct
type AssetsResponse struct {
	Assets Assets `json:"balances,required"`
}

// GetBalances returns user assets on exchange.
func (b *Binance) GetBalances() (response *Assets, err error) {
	var rawResponse AssetsResponse

	err = b.Request(exchange.RequestParams{
		Method: "GET", Endpoint: "/account", Auth: true,
	}, &rawResponse)
	response = &rawResponse.Assets

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

// NewOrder creates a new order.
func (b *Binance) NewOrder(params NewOrder) (response *OrderResponse, err error) {
	params.Symbol = strings.ToUpper(params.Symbol)

	if params.OrderID == "" {
		params.OrderID = util.GenerateUUID()
	}

	response = &OrderResponse{}

	err = b.Request(exchange.RequestParams{
		Auth: true, Params: params,
		Method: "POST", Endpoint: "/order",
	}, &response)

	return
}

// CancelOrder cancels an order.
func (b *Binance) CancelOrder(orderID, symbol string) (response *OrderResponse, err error) {
	params := struct {
		Symbol  string `json:"symbol,required"`
		OrderID string `json:"origClientOrderId,required"`
	}{
		OrderID: orderID,
		Symbol:  strings.ToUpper(symbol),
	}

	response = &OrderResponse{}

	err = b.Request(exchange.RequestParams{
		Auth: true, Params: params,
		Method: "DELETE", Endpoint: "/order",
	}, &response)

	return
}
