package main

import (
	"fmt"

	"github.com/girikuncoro/go-bitcoin/pkg/factory"
)

func main() {
	// secretKey := "TODO"
	// apiKey := "TODO"

	vip := factory.CoinExchangeFactory(factory.Vip)
	vipRes, _ := vip.GetPrice("eth")
	fmt.Printf("VIP: %s IDR\n", vipRes)

	cb := factory.CoinExchangeFactory(factory.Coinbase)
	cbRes, _ := cb.GetPrice("eth")
	fmt.Printf("Coinbase: %s IDR\n", cbRes)

	g := factory.CoinExchangeFactory(factory.Gemini)
	gRes, _ := g.GetPrice("eth")
	fmt.Printf("Gemini: %s USD\n", gRes)
}
