package main

import (
	"log"

	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/binance/converter"
)

var Binance *binance.Binance

func main() {
	var err error

	Binance, err = binance.New("YOUR_BINANCE_PUBLIC_KEY", "YOUR_BINANCE_SECRET_KEY")
	if err != nil {
		log.Panic(err)
	}

	ToUSD()

	GetSymbols()
	GetBalances()

	NewOrder()
	CancelOrder()
}

const price, quantity = 10000.0, 10.0

type repository struct{}

// GetPrice returns fake price (BTC/USD)
func (r *repository) GetPrice(_, _ string) float64 {
	return price
}

// GetSymbol returns fake symbol detail (BTC/USD)[Demo]
func (r *repository) GetSymbol(_, _ string) interface{} {
	return &binance.Symbol{
		Base:  binance.BTC,
		Quote: binance.USD,
		ID:    binance.Demo,
	}
}

func ToUSD() {
	cache := &repository{}
	value, err := converter.ToUSD(cache, binance.BTC, quantity, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(value)
}

func GetSymbols() {
	symbols, err := Binance.GetSymbols()
	if err != nil {
		log.Panic(err)
	}

	log.Println(symbols)
}

func GetBalances() {
	balances, err := Binance.GetBalances()
	if err != nil {
		log.Panic(err)
	}

	log.Println(balances)
}

func NewOrder() {
	order, err := Binance.NewOrder(binance.NewOrder{
		Price:    10,
		Quantity: 1,

		Side:        binance.Buy,
		Type:        binance.Limit,
		Symbol:      binance.BNB + binance.USD,
		TimeInForce: binance.GoodTillCancel,
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}

func CancelOrder() {
	order, err := Binance.CancelOrder("FAKE_ORDER_ID", binance.BNB+binance.USD)
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}
