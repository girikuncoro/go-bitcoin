package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/girikuncoro/go-bitcoin/pkg/client/base_client"

	"github.com/girikuncoro/go-bitcoin/pkg/client/fixer_client"
	"github.com/girikuncoro/go-bitcoin/pkg/factory"
)

func main() {
	coin := "eth"
	fmt.Printf("%s Arbitrage\n", strings.ToUpper(coin))

	f := fixer_client.NewFixerClient(base_client.NewBaseClient())
	rate, _ := f.GetRate("USD", "IDR")
	fmt.Printf("Rate USD-IDR: %v\n", rate.Rates.IDR)

	fmt.Println("==============")

	// secretKey := "TODO"
	// apiKey := "TODO"
	vip := factory.CoinExchangeFactory(factory.Vip)
	vipRes, _ := vip.GetPrice(coin)
	fmt.Printf("VIP: %s IDR\n", vipRes)

	cb := factory.CoinExchangeFactory(factory.Coinbase)
	cbRes, _ := cb.GetPrice(coin)
	cbPrice, _ := strconv.ParseFloat(cbRes, 64)
	fmt.Printf("Coinbase: %9.0f IDR\n", float64(cbPrice)*rate.Rates.IDR)

	g := factory.CoinExchangeFactory(factory.Gemini)
	gRes, _ := g.GetPrice(coin)
	gPrice, _ := strconv.ParseFloat(gRes, 64)
	fmt.Printf("Gemini: %9.0f IDR\n", float64(gPrice)*rate.Rates.IDR)
}
