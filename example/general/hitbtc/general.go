package main

import (
	"log"

	"github.com/ramezanius/crypex/exchange/hitbtc"
	"github.com/ramezanius/crypex/exchange/hitbtc/converter"
)

var HitBTC *hitbtc.HitBTC

func main() {
	var err error

	HitBTC, err = hitbtc.New("YOUR_HITBTC_PUBLIC_KEY", "YOUR_HITBTC_SECRET_KEY")
	if err != nil {
		log.Panic(err)
	}

	ToUSD()

	GetSymbol()
	GetSymbols()
	GetBalances()

	NewOrder()
	CancelOrder()
	ReplaceOrder()
}

const price, quantity = 10000.0, 10.0

type repository struct{}

// GetPrice returns fake price (BTC/USD)
func (r *repository) GetPrice(_, _ string) float64 {
	return price
}

// GetSymbol returns fake symbol detail (BTC/USD)[Demo]
func (r *repository) GetSymbol(_, _ string) interface{} {
	return &hitbtc.Symbol{
		Base:  hitbtc.BTC,
		Quote: hitbtc.USD,
		ID:    hitbtc.Demo,
	}
}

func ToUSD() {
	cache := &repository{}
	value, err := converter.ToUSD(cache, hitbtc.BTC, quantity, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(value)
}

func GetSymbol() {
	symbol, err := HitBTC.GetSymbol(hitbtc.Demo)
	if err != nil {
		log.Panic(err)
	}

	log.Println(symbol)
}

func GetSymbols() {
	symbols, err := HitBTC.GetSymbols()
	if err != nil {
		log.Panic(err)
	}

	log.Println(symbols)
}

func GetBalances() {
	balances, err := HitBTC.GetBalances()
	if err != nil {
		log.Panic(err)
	}

	log.Println(balances)
}

func NewOrder() {
	order, err := HitBTC.NewOrder(hitbtc.NewOrder{
		Price:    2000,
		Quantity: 0.00002,

		Side:   hitbtc.Buy,
		Type:   hitbtc.Limit,
		Symbol: hitbtc.Demo,
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}

func CancelOrder() {
	order, err := HitBTC.CancelOrder("FAKE_ORDER_ID")
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}

func ReplaceOrder() {
	order, err := HitBTC.ReplaceOrder(hitbtc.ReplaceOrder{
		Price:    1000,
		Quantity: 0.00001,
		OrderID:  "FAKE_ORDER_ID",
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}
