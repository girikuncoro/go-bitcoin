package coinbase_client

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Ticker struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	}
}

func (c *CoinbaseClient) GetPrice(coinName string) (string, error) {
	req, err := c.baseClient.Request("GET", c.buildTickerURL(coinName, CoinbaseCurrency), "")
	if err != nil {
		return "", err
	}

	res, err := c.baseClient.Response(req)
	if err != nil {
		return "", err
	}

	t := &Ticker{}
	err = json.Unmarshal(res, &t)
	if err != nil {
		return "", err
	}

	return t.Data.Amount, nil
}

func (c *CoinbaseClient) buildTickerURL(coinName, currency string) string {
	url := fmt.Sprintf(TickerURL, strings.ToUpper(coinName), strings.ToUpper(currency))
	return url
}
