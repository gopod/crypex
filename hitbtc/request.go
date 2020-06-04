package hitbtc

type Symbol struct {
	ID    string `json:"id"`
	Base  string `json:"baseCurrency"`
	Quote string `json:"quoteCurrency"`

	TickSize             string `json:"tickSize"`
	FeeCurrency          string `json:"feeCurrency"`
	QuantityIncrement    string `json:"quantityIncrement"`
	TakeLiquidityRate    string `json:"takeLiquidityRate"`
	ProvideLiquidityRate string `json:"provideLiquidityRate"`
}

type Symbols []Symbol

func (h *HitBTC) GetSymbol(symbol string) (response *Symbol, err error) {
	request := struct {
		Symbol string `json:"symbol"`
	}{Symbol: symbol}

	err = h.Request("getSymbol", &request, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (h *HitBTC) GetSymbols() (response *Symbols, err error) {
	err = h.Request("getSymbols", nil, &response)
	if err != nil {
		return nil, err
	}

	return
}

type Balance struct {
	Currency string `json:"currency"`

	Reserved  float64 `json:"reserved,string"`
	Available float64 `json:"available,string"`
}

type Balances []Balance

func (h *HitBTC) GetBalances() (response *Balances, err error) {
	err = h.Request("getTradingBalance", nil, &response)
	if err != nil {
		return nil, err
	}

	return
}
