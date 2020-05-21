package hitbtc

type Symbol struct {
	ID    string `json:"id,required"`
	Base  string `json:"baseCurrency,required"`
	Quote string `json:"quoteCurrency,required"`

	TickSize             string `json:"tickSize,required"`
	FeeCurrency          string `json:"feeCurrency,required"`
	QuantityIncrement    string `json:"quantityIncrement,required"`
	TakeLiquidityRate    string `json:"takeLiquidityRate,required"`
	ProvideLiquidityRate string `json:"provideLiquidityRate,required"`
}

type Symbols []Symbol

func (h *HitBTC) GetSymbol(symbol string) (response *Symbol, err error) {
	request := struct {
		Symbol string `json:"symbol,required"`
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
