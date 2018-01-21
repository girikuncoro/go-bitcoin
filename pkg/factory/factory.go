package factory

import (
	"github.com/girikuncoro/go-bitcoin/pkg/client/base_client"

	"github.com/girikuncoro/go-bitcoin/pkg/client/coinbase_client"

	"github.com/girikuncoro/go-bitcoin/pkg/client/vip_client"
)

type CoinExchange interface {
	GetPrice(string) (string, error)
}

type CoinExchangeType int

const (
	Vip CoinExchangeType = 1 << iota
	Coinbase
)

func CoinExchangeFactory(t CoinExchangeType) CoinExchange {
	switch t {
	case Vip:
		return vip_client.NewVipClient(base_client.NewBaseClient())
	case Coinbase:
		return coinbase_client.NewCoinbaseClient(base_client.NewBaseClient())
	default:
		return nil
	}
}
