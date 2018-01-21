package fixer_client

import (
	"encoding/json"
	"fmt"

	"github.com/girikuncoro/go-bitcoin/pkg/client/base_client"
)

const (
	FixerUrl = "https://api.fixer.io/latest?base=%s&symbols=%s"

	ParamBase   = "base"
	ParamSymbol = "symbol"
)

type FixerClient struct {
	baseClient *base_client.BaseClient
}

type FixerResp struct {
	Base  string `json:"base"`
	Rates struct {
		IDR float64 `json:"IDR"`
	}
}

func NewFixerClient(baseClient *base_client.BaseClient) *FixerClient {
	return &FixerClient{
		baseClient: baseClient,
	}
}

func (c *FixerClient) GetRate(baseCurency, convertedCurrency string) (*FixerResp, error) {
	req, err := c.baseClient.Request("GET", c.buildURL(baseCurency, convertedCurrency), "")
	if err != nil {
		return nil, err
	}

	res, err := c.baseClient.Response(req)
	if err != nil {
		return nil, err
	}

	f := &FixerResp{}
	err = json.Unmarshal(res, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (c *FixerClient) buildURL(baseCurency, convertedCurrency string) string {
	url := fmt.Sprintf(FixerUrl, baseCurency, convertedCurrency)
	return url
}
