package factory

import (
	"github.com/girikuncoro/go-bitcoin/pkg/client/base_client"

	"github.com/girikuncoro/go-bitcoin/pkg/client/coinbase_client"

	"github.com/girikuncoro/go-bitcoin/pkg/client/vip_client"

	"github.com/girikuncoro/go-bitcoin/pkg/client/gemini_client"
)

type CoinExchanger interface {
	GetPrice(string) (string, error)
}

type CoinExchangeType int

const (
	Vip CoinExchangeType = 1 << iota
	Coinbase
	Gemini
)

func CoinExchangeFactory(t CoinExchangeType) CoinExchanger {
	switch t {
	case Vip:
		return vip_client.NewVipClient(base_client.NewBaseClient())
	case Coinbase:
		return coinbase_client.NewCoinbaseClient(base_client.NewBaseClient())
	case Gemini:
		return gemini_client.NewGeminiClient(base_client.NewBaseClient())
	default:
		return nil
	}
}
