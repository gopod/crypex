package hitbtc

import (
	"fmt"
	"time"

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

// AssetsResponse struct
type AssetsResponse Assets

// GetBalances returns user assets on exchange.
func (h *HitBTC) GetBalances() (response *Assets, err error) {
	response = &Assets{}

	err = h.Request(exchange.RequestParams{
		Method: "GET", Endpoint: "/trading/balance", Auth: true,
	}, &response)

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

	StopPrice  float64    `json:"stopPrice,omitempty"`
	ExpireTime *time.Time `json:"expireTime,omitempty"`

	PostOnly       bool   `json:"postOnly,omitempty"`
	TimeInForce    string `json:"timeInForce,omitempty"`
	StrictValidate bool   `json:"strictValidate,omitempty"`
}

// NewOrder creates a new order.
func (h *HitBTC) NewOrder(params NewOrder) (response *ReportsResponse, err error) {
	if params.OrderID == "" {
		params.OrderID = util.GenerateUUID()
	}

	response = &ReportsResponse{}

	err = h.Request(exchange.RequestParams{
		Auth:     true,
		Params:   params,
		Method:   "PUT",
		Endpoint: fmt.Sprintf("/order/%s", params.OrderID),
	}, &response)

	return
}

// CancelOrder cancels an order.
func (h *HitBTC) CancelOrder(orderID string) (response *ReportsResponse, err error) {
	response = &ReportsResponse{}

	err = h.Request(exchange.RequestParams{
		Auth:     true,
		Method:   "DELETE",
		Endpoint: fmt.Sprintf("/order/%s", orderID),
	}, &response)

	return
}
