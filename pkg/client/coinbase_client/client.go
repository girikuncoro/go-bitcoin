package coinbase_client

import (
	"github.com/girikuncoro/go-bitcoin/pkg/client/base_client"
)

type CoinbaseClienter interface {
	GetPrice()
}

type CoinbaseClient struct {
	baseClient *base_client.BaseClient
}

func NewCoinbaseClient(baseClient *base_client.BaseClient) *CoinbaseClient {
	return &CoinbaseClient{
		baseClient: baseClient,
	}
}
