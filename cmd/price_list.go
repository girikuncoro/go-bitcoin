package main

import (
	"fmt"
	"strconv"

	"github.com/girikuncoro/go-bitcoin/pkg/client/base_client"
	"github.com/girikuncoro/go-bitcoin/pkg/client/fixer_client"
	"github.com/girikuncoro/go-bitcoin/pkg/factory"
)

func priceListRun(cli *Cli) error {
	if len(cli.Args) == 0 {
		return fmt.Errorf("Coin name must be provided")
	}

	coinName := cli.Args[0]

	cli.printer.Printf("Listing price for (coinname: %s)\n", coinName)

	f := fixer_client.NewFixerClient(base_client.NewBaseClient())
	rate, _ := f.GetRate("USD", "IDR")
	cli.printer.Printf("Rate USD-IDR: %v\n", rate.Rates.IDR)

	cli.printer.Printf("========================\n")

	vip := factory.CoinExchangeFactory(factory.Vip)
	vipRes, _ := vip.GetPrice(coinName)
	cli.printer.Printf("VIP: %s IDR\n", vipRes)

	cb := factory.CoinExchangeFactory(factory.Coinbase)
	cbRes, _ := cb.GetPrice(coinName)
	cbPrice, _ := strconv.ParseFloat(cbRes, 64)
	cli.printer.Printf("Coinbase: %9.0f IDR\n", float64(cbPrice)*rate.Rates.IDR)

	g := factory.CoinExchangeFactory(factory.Gemini)
	gRes, _ := g.GetPrice(coinName)
	gPrice, _ := strconv.ParseFloat(gRes, 64)
	cli.printer.Printf("Gemini: %9.0f IDR\n", float64(gPrice)*rate.Rates.IDR)

	return nil
}
