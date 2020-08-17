package hitbtc

import (
	"fmt"

	"github.com/ramezanius/crypex/exchange"
	"github.com/ramezanius/crypex/exchange/util"
)

// GetSymbols returns exchange symbols.
func (h *HitBTC) GetSymbols() (response *Symbols, err error) {
	response = &Symbols{}

	err = h.Request(exchange.RequestParams{
		Method: "GET", Endpoint: "/public/symbol",
	}, &response)

	return
}

// GetCandles returns symbol candles.
func (h *HitBTC) GetCandles(params CandlesParams) (response *CandlesResponse, err error) {
	if params.Sort == "" {
		params.Sort = "DESC"
	}

	response = &CandlesResponse{}

	err = h.Request(exchange.RequestParams{
		Method:   "GET",
		Params:   params,
		Endpoint: fmt.Sprintf("/public/candles/%s", params.Symbol),
	}, &response)

	return
}

// GetBalances returns user assets on exchange.
func (h *HitBTC) GetBalances() (response *Assets, err error) {
	response = &Assets{}

	err = h.Request(exchange.RequestParams{
		Method: "GET", Endpoint: "/trading/balance", Auth: true,
	}, &response)

	return
}

// NewOrder creates a new order.
func (h *HitBTC) NewOrder(params NewOrder) (response *OrderResponse, err error) {
	if params.OrderID == "" {
		params.OrderID = util.GenerateUUID()
	}

	response = &OrderResponse{}

	err = h.Request(exchange.RequestParams{
		Auth:     true,
		Params:   params,
		Method:   "PUT",
		Endpoint: fmt.Sprintf("/order/%s", params.OrderID),
	}, &response)

	return
}

// CancelOrder cancels an order.
func (h *HitBTC) CancelOrder(orderID string) (response *OrderResponse, err error) {
	response = &OrderResponse{}

	err = h.Request(exchange.RequestParams{
		Auth:     true,
		Method:   "DELETE",
		Endpoint: fmt.Sprintf("/order/%s", orderID),
	}, &response)

	return
}
